package models

type JFullMethod struct {
	Name              string
	StartLine         int
	StartLinePosition int
	StopLine          int
	StopLinePosition  int
}

type JField struct {
	Name   string
	Source string
	StartLine         int
	StopLine          int
	//StartLinePosition int
	//StopLinePosition  int
}

type JImport struct {
	Name   string
	StartLine         int
	StopLine          int
}

type JPkgInfo struct {
	Name   string
	StartLine         int
	StopLine          int
}

var methods []JFullMethod
var fields = make(map[string]JField)
var imports []JImport
var pkgInfo JPkgInfo

type JFullIdentifier struct {
	Pkg  string
	Name string
	Type string
}

func NewJFullIdentifier() *JFullIdentifier {
	identifier := &JFullIdentifier{"", "", ""}
	methods = nil
	fields = make(map[string]JField)
	imports = nil
	return identifier
}

func (identifier *JFullIdentifier) AddMethod(method JFullMethod) {
	methods = append(methods, method)
}

func (identifier *JFullIdentifier) GetMethods() []JFullMethod {
	return methods
}

func (identifier *JFullIdentifier) AddField(field JField) {
	fields[field.Name] = field
}

func (identifier *JFullIdentifier) GetFields() map[string]JField {
	return fields
}

func (identifier *JFullIdentifier) AddImport(jImport JImport) {
	imports = append(imports, jImport)
}

func (identifier *JFullIdentifier) GetImports() []JImport {
	return imports
}

func (identifier *JFullIdentifier) SetPkgInfo(info JPkgInfo) {
	pkgInfo = info
}

func (identifier *JFullIdentifier) GetPkgInfo() JPkgInfo {
	return pkgInfo
}