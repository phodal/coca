package ast_go

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"github.com/yourbasic/radix"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

var currentPackage *core_domain.CodePackage
var identCodeMembers []core_domain.CodeMember

type CocagoParser struct {
	CodeMembers    []core_domain.CodeMember
	PackageManager core_domain.CodePackageInfo
}

var output io.Writer

func NewCocagoParser() *CocagoParser {
	currentPackage = &core_domain.CodePackage{}
	output = os.Stdout
	if flag.Lookup("test") == nil {
		output = new(bytes.Buffer)
	}
	return &CocagoParser{}
}

func (n *CocagoParser) SetOutput(out io.Writer) io.Writer {
	output = out
	return output
}

func (n *CocagoParser) ProcessFile(fileName string) core_domain.CodeContainer {
	absPath, _ := filepath.Abs(fileName)
	content, _ := ioutil.ReadFile(absPath)

	fmt.Fprintf(output, "process file %s\n", fileName)

	code := string(content)

	codeFile := n.ProcessString(code, fileName, nil)
	return *codeFile
}

func (n *CocagoParser) ProcessString(code string, fileName string, codeMembers []core_domain.CodeMember) *core_domain.CodeContainer {
	identCodeMembers = codeMembers
	n.CodeMembers = codeMembers
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, fileName, code, 0)
	if err != nil {
		panic(err)
	}

	codeFile := n.Visitor(f, fset, fileName)
	currentPackage.CodeFiles = append(currentPackage.CodeFiles, *codeFile)
	return codeFile
}

func (n *CocagoParser) IdentAnalysis(code string, fileName string) *core_domain.CodeContainer {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, fileName, code, 0)
	if err != nil {
		panic(err)
	}

	codeFile := n.Visitor(f, fset, fileName)
	return codeFile
}

func (n *CocagoParser) Visitor(f *ast.File, fset *token.FileSet, fileName string) *core_domain.CodeContainer {
	var currentStruct core_domain.CodeDataStruct
	var currentFile core_domain.CodeContainer
	var currentFunc *core_domain.CodeFunction
	var dsMap = make(map[string]*core_domain.CodeDataStruct)

	packageName := BuildImportName(fileName)
	currentFile.FullName = packageName
	currentPackage.Name = packageName

	var funcType = ""
	var lastIdent = ""

	ast.Inspect(f, func(node ast.Node) bool {
		switch x := node.(type) {
		case *ast.Ident:
			lastIdent = x.Name
		case *ast.File:
			currentFile.PackageName = x.Name.String()
		case *ast.ImportSpec:
			imp := BuildImport(x, fileName, n.PackageManager)
			currentFile.Imports = append(currentFile.Imports, *imp)
		case *ast.ValueSpec:
			names := x.Names
			for _, name := range names {
				_, selSource, selName := BuildValSpec(x.Type)
				field := core_domain.CodeField{
					TypeType:  selSource + "." + selName,
					TypeValue: name.Name,
				}

				currentFile.Fields = append(currentFile.Fields, field)
			}
		case *ast.TypeSpec:
			currentStruct = core_domain.CodeDataStruct{}
			currentStruct.NodeName = x.Name.Name
			currentStruct.Package = currentFile.PackageName
			//currentStruct.FilePath = BuildImportName(fileName)
			dsMap[currentStruct.NodeName] = &currentStruct
		case *ast.StructType:
			AddStructType(currentStruct.NodeName, x, &currentFile, dsMap)
		case *ast.FuncDecl:
			funcType = "FuncDecl"
			currentFunc, recv := AddFunctionDecl(x, &currentFile)
			if recv != "" {
				dsMap[recv].Functions = append(dsMap[recv].Functions, *currentFunc)
			}
		case *ast.FuncType:
			if funcType != "FuncDecl" {
				AddNestedFunction(currentFunc, x)
			}

			funcType = ""
		case *ast.InterfaceType:
			// todo: dirty fix
			if len(x.Methods.List) < 1 {
				break
			}
			currentStruct := AddInterface(x, lastIdent, &currentFile)
			dsMap[currentStruct.NodeName] = &currentStruct
		default:
			if reflect.TypeOf(x) != nil && reflect.TypeOf(output).String() != "*bytes.Buffer" {
				//fmt.Fprintf(output, "Visitor case %s\n", reflect.TypeOf(x))
			}
		}
		return true
	})

	currentFile.DataStructures = nil
	for _, ds := range dsMap {
		currentFile.DataStructures = append(currentFile.DataStructures, *ds)
	}
	SortInterface(currentFile.DataStructures)

	return &currentFile
}

func (n *CocagoParser) SetPackageManager(manager core_domain.CodePackageInfo) {
	n.PackageManager = manager
}

func BuildValSpec(expr ast.Expr) (string, string, string) {
	switch x := expr.(type) {
	case *ast.StarExpr:
		return BuildExpr(x.X)
	default:
		fmt.Fprintf(output, "Visitor case %s\n", reflect.TypeOf(x))
	}
	return "", "", ""
}

func SortInterface(slice []core_domain.CodeDataStruct) {
	radix.SortSlice(slice, func(i int) string { return slice[i].NodeName })
}

func BuildImport(x *ast.ImportSpec, fileName string, manager core_domain.CodePackageInfo) *core_domain.CodeImport {
	path := x.Path.Value
	cleanPath := path[1 : len(path)-1]
	asName := ""
	if x.Name != nil {
		asName = x.Name.String()
	}
	moduleName := manager.ProjectName
	if moduleName == "" {
		moduleName = "github.com/phodal/coca"
	}
	withOutModuleName := strings.ReplaceAll(cleanPath, moduleName, "")
	all := strings.ReplaceAll(withOutModuleName, "/", ".")
	if strings.HasPrefix(all, ".") {
		all = all[1:]
	}
	imp := &core_domain.CodeImport{
		Source:     all,
		AsName:     asName,
		ImportName: "",
		UsageName:  nil,
		Scope:      "",
	}

	return imp
}

func BuildImportName(fileName string) string {
	fileName = filepath.FromSlash(fileName)
	splitFileName := strings.Split(fileName, string(filepath.Separator))
	importName := ""
	if len(splitFileName) > 2 {
		importName = strings.Join(splitFileName[:len(splitFileName)-1], ".")
	}
	if strings.HasPrefix(importName, ".") {
		importName = importName[1:]
	}
	return importName
}

func AddInterface(x *ast.InterfaceType, ident string, codeFile *core_domain.CodeContainer) core_domain.CodeDataStruct {
	properties := BuildFieldToProperty(x.Methods.List)

	dataStruct := core_domain.CodeDataStruct{
		NodeName:        ident,
		InOutProperties: properties,
	}

	member := core_domain.NewCodeMember()
	member.DataStructID = ident
	member.Type = "interface"
	setMemberPackageInfo(member, codeFile)

	codeFile.Members = append(codeFile.Members, *member)

	return dataStruct
}

func setMemberPackageInfo(member *core_domain.CodeMember, codeFile *core_domain.CodeContainer) {
	member.AliasPackage = codeFile.PackageName
	member.FileID = codeFile.FullName
	member.BuildMemberId()
}

func AddNestedFunction(currentFunc *core_domain.CodeFunction, x *ast.FuncType) {

}

func AddFunctionDecl(x *ast.FuncDecl, currentFile *core_domain.CodeContainer) (*core_domain.CodeFunction, string) {
	recv := ""
	if x.Recv != nil {
		recv = BuildReceiver(x, recv)
	}
	codeFunc := BuildFunction(x, currentFile)

	if recv == "" {
		member := &core_domain.CodeMember{
			DataStructID: "default",
			Type:         "method",
		}

		member.FunctionNodes = append(member.FunctionNodes, *codeFunc)
		setMemberPackageInfo(member, currentFile)
		currentFile.Members = append(currentFile.Members, *member)
	}

	return codeFunc, recv
}

func BuildReceiver(x *ast.FuncDecl, recv string) string {
	for _, item := range x.Recv.List {
		switch x := item.Type.(type) {
		case *ast.StarExpr:
			recv = getStarExprName(*x)
		case *ast.Ident:
			recv = x.Name
		default:
			fmt.Fprintf(output, "AddFunctionDecl %s\n", reflect.TypeOf(x))
		}
	}
	return recv
}

func BuildExpr(expr ast.Expr) (string, string, string) {
	switch x := expr.(type) {
	case *ast.SelectorExpr:
		selector := ""
		switch sele := x.X.(type) {
		case *ast.Ident:
			selector = sele.String()
		case *ast.IndexExpr:
			_, indexVar, _ := BuildExpr(sele.X)
			selector = indexVar
		default:
			fmt.Println("BuildExpr selector", reflect.TypeOf(sele))
		}

		selName := x.Sel.Name
		return "selector", selector, selName
	case *ast.BasicLit:
		return "basiclit", x.Value, x.Kind.String()
	case *ast.Ident:
		name := ""
		if x.Obj != nil {
			name = x.Obj.Kind.String()
		}
		return "ident", x.Name, name
	case *ast.CallExpr:
		_, value, _ := BuildExpr(x.Fun)
		var callArgs []string
		for _, arg := range x.Args {
			argType, argValue, argKind := BuildExpr(arg)
			if argType == "selector" {
				callArgs = append(callArgs, argValue+"."+argKind)
			}
		}
		return "call", value, strings.Join(callArgs, ",")
	case *ast.FuncLit:
		// inner function
	case *ast.TypeAssertExpr:
	case *ast.BinaryExpr:
		_, val, s2 := BuildExpr(x.X)
		return "call", val, s2
	default:
		fmt.Fprintf(output, "BuildExpr %s\n", reflect.TypeOf(x))
	}
	return "", "", ""
}

func getFieldName(field *ast.Field) string {
	if len(field.Names) < 1 {
		return ""
	}
	return field.Names[0].Name
}

func AddStructType(currentNodeName string, x *ast.StructType, currentFile *core_domain.CodeContainer, dsMap map[string]*core_domain.CodeDataStruct) {
	member := core_domain.NewCodeMember()
	member.DataStructID = currentNodeName
	member.Type = "struct"
	setMemberPackageInfo(member, currentFile)

	var ioproperties []core_domain.CodeProperty
	var calls []core_domain.CodeCall
	for _, field := range x.Fields.List {
		property := BuildPropertyField(getFieldName(field), field)
		member.FileID = currentFile.FullName
		ioproperties = append(ioproperties, *property)

		call := core_domain.CodeCall{
			Package:  getPackageName(property.TypeValue, "", currentFile.Imports),
			NodeName: property.TypeValue,
		}
		calls = append(calls, call)
	}

	// todo : when dsMap key-value create it
	if dsMap[currentNodeName] != nil {
		dsMap[currentNodeName].InOutProperties = ioproperties
		dsMap[currentNodeName].FunctionCalls = append(dsMap[currentNodeName].FunctionCalls, calls...)
	}
	currentFile.Members = append(currentFile.Members, *member)
}
