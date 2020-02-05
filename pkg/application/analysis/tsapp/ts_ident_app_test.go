package tsapp

import (
	. "github.com/onsi/gomega"
	"io/ioutil"
	"testing"
)

func Test_TypeScriptClassNode(t *testing.T) {
	g := NewGomegaWithT(t)

	app := new(TypeScriptIdentApp)
	codefile := app.Analysis(`
interface IPerson {
    name: string;
}

class Person implements IPerson {
    public publicString: string;
    private privateString: string;
    protected protectedString: string;
    readonly readonlyString: string;
    name: string;

    constructor(name: string) {
        this.name = name;
    }
}
`, "")

	g.Expect(codefile.DataStructures[0].NodeName).To(Equal("IPerson"))
	g.Expect(codefile.DataStructures[1].NodeName).To(Equal("Person"))
	g.Expect(codefile.DataStructures[1].Functions[0].Name).To(Equal("constructor"))
	g.Expect(codefile.DataStructures[1].Implements[0]).To(Equal("IPerson"))
}

func Test_TypeScriptMultipleClass(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	app := new(TypeScriptIdentApp)
	code, _ := ioutil.ReadFile("../../../../_fixtures/ts/grammar/Class.ts")

	codeFile := app.Analysis(string(code), "")

	g.Expect(len(codeFile.DataStructures)).To(Equal(4))
	g.Expect(len(codeFile.DataStructures[1].Implements)).To(Equal(1))
	g.Expect(codeFile.DataStructures[1].Implements[0]).To(Equal("IPerson"))
}

func Test_TypeScriptAbstractClass(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	app := new(TypeScriptIdentApp)
	code, _ := ioutil.ReadFile("../../../../_fixtures/ts/grammar/AbstractClass.ts")

	codeFile := app.Analysis(string(code), "")

	g.Expect(len(codeFile.DataStructures)).To(Equal(3))
	g.Expect(codeFile.DataStructures[0].Type).To(Equal("Class"))
	g.Expect(codeFile.DataStructures[1].NodeName).To(Equal("Employee"))
	g.Expect(codeFile.DataStructures[1].Extend).To(Equal("Person"))
}

func Test_ShouldGetClassFromModule(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	app := new(TypeScriptIdentApp)
	code, _ := ioutil.ReadFile("../../../../_fixtures/ts/grammar/Module.ts")

	results := app.Analysis(string(code), "")

	g.Expect(len(results.DataStructures)).To(Equal(1))
	g.Expect(results.DataStructures[0].NodeName).To(Equal("Employee"))
}

func Test_ShouldEnableGetClassMethod(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	app := new(TypeScriptIdentApp)

	codefile := app.Analysis(`
class Employee  {
    displayName():void {
        console.log("hello, world");
    }
}
`, "")

	g.Expect(len(codefile.DataStructures[0].Functions)).To(Equal(1))
}

func Test_ShouldGetInterfaceImplements(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	app := new(TypeScriptIdentApp)

	results := app.Analysis(`
export interface IPerson {
    name: string;
    gender: string;
}

interface IEmployee extends IPerson{
    empCode: number;
    readonly empName: string;
    empDept?:string;
    getSalary: (number) => number; // arrow function
    getManagerName(number): string;
}
`, "")

	g.Expect(results.DataStructures[1].Extend).To(Equal("IPerson"))
}

func Test_ShouldGetInterfaceProperty(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	app := new(TypeScriptIdentApp)
	codeFile := app.Analysis(`
export interface IPerson {
    name: string;
    gender: string;
    getSalary: (number) => number;
    getManagerName(number): string;
}
`, "")

	firstDataStruct := codeFile.DataStructures[0]

	g.Expect(len(firstDataStruct.Fields)).To(Equal(2))
	g.Expect(len(firstDataStruct.Functions)).To(Equal(2))

	firstMethod := firstDataStruct.Functions[0]
	secondMethod := firstDataStruct.Functions[1]

	g.Expect(firstMethod.Name).To(Equal("getSalary"))
	g.Expect(secondMethod.Name).To(Equal("getManagerName"))
	g.Expect(secondMethod.Parameters[0].TypeType).To(Equal("number"))
}

func Test_ShouldGetDefaultFunctionName(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	app := new(TypeScriptIdentApp)

	codeFile := app.Analysis(`
function Sum(x: number, y: number) : void {
    console.log('processNumKeyPairs: key = ' + key + ', value = ' + value)
    return x + y;
}
`, "")

	ds := codeFile.DataStructures

	firstFunction := ds[0].Functions[0]
	params := firstFunction.Parameters
	g.Expect(firstFunction.MultipleReturns[0].TypeType).To(Equal("void"))
	g.Expect(len(params)).To(Equal(2))
	g.Expect(params[0].TypeValue).To(Equal("x"))
	g.Expect(params[0].TypeType).To(Equal("number"))
}

func Test_ShouldHandleRestParameters(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	app := new(TypeScriptIdentApp)

	codeFile := app.Analysis(`
function buildName(firstName: string, ...restOfName: string[]) {
  return firstName + " " + restOfName.join(" ");
}
`, "")

	firstFunction:= codeFile.DataStructures[0].Functions[0]
	params := firstFunction.Parameters
	g.Expect(len(params)).To(Equal(2))
	g.Expect(params[0].TypeValue).To(Equal("firstName"))
	g.Expect(params[1].TypeValue).To(Equal("restOfName"))
}

func Test_ShouldGetClassFields(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	app := new(TypeScriptIdentApp)
	code, _ := ioutil.ReadFile("../../../../_fixtures/ts/grammar/Class.ts")

	codeFile := app.Analysis(string(code), "")

	codeFields := codeFile.DataStructures[1].Fields
	g.Expect(len(codeFields)).To(Equal(5))
	g.Expect(codeFields[0].Modifiers[0]).To(Equal("public"))
}

func Test_ShouldReturnBlockImports(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	app := new(TypeScriptIdentApp)

	results := app.Analysis(`
import { ZipCodeValidator } from "./ZipCodeValidator";

`, "")

	g.Expect(len(results.Imports)).To(Equal(1))
	g.Expect(results.Imports[0].Source).To(Equal("./ZipCodeValidator"))
}

func Test_ShouldReturnAsImports(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	app := new(TypeScriptIdentApp)

	results := app.Analysis(`
import zip = require("./ZipCodeValidator");
`, "")

	g.Expect(len(results.Imports)).To(Equal(1))

	g.Expect(results.Imports[0].Source).To(Equal("./ZipCodeValidator"))
}

// Todo: fix for $ and *
func Test_ShouldReturnAllImports(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	app := new(TypeScriptIdentApp)
	code, _ := ioutil.ReadFile("../../../../_fixtures/ts/grammar/Import.ts")
	results := app.Analysis(string(code), "")

	g.Expect(len(results.Imports)).To(Equal(5))
}
