package ts

import (
	"fmt"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"testing"
)

func Test_TypeScriptClassNode(t *testing.T) {
	g := NewGomegaWithT(t)

	app := new(TypeScriptApiApp)
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

	results := codefile.ClassNodes

	g.Expect(results[0].Class).To(Equal("IPerson"))
	g.Expect(results[1].Class).To(Equal("Person"))
	g.Expect(results[1].Methods[0].Name).To(Equal("constructor"))
	g.Expect(results[1].Implements[0]).To(Equal("IPerson"))
}

func Test_TypeScriptMultipleClass(t *testing.T) {
	g := NewGomegaWithT(t)

	app := new(TypeScriptApiApp)
	code, _ := ioutil.ReadFile("../../../../_fixtures/ts/grammar/Class.ts")

	codeFile := app.Analysis(string(code), "")

	results := codeFile.ClassNodes

	g.Expect(len(results)).To(Equal(4))
	g.Expect(len(codeFile.DataStructures)).To(Equal(3))
	g.Expect(results[1].Implements[0]).To(Equal("IPerson"))
}

func Test_TypeScriptAbstractClass(t *testing.T) {

	g := NewGomegaWithT(t)

	app := new(TypeScriptApiApp)
	code, _ := ioutil.ReadFile("../../../../_fixtures/ts/grammar/AbstractClass.ts")

	results := app.Analysis(string(code), "").ClassNodes

	g.Expect(len(results)).To(Equal(3))
	g.Expect(results[0].Type).To(Equal("Class"))
	g.Expect(results[1].Class).To(Equal("Employee"))
	g.Expect(results[1].Extend).To(Equal("Person"))
}

func Test_ShouldGetClassFromModule(t *testing.T) {
	g := NewGomegaWithT(t)

	app := new(TypeScriptApiApp)
	code, _ := ioutil.ReadFile("../../../../_fixtures/ts/grammar/Module.ts")

	results := app.Analysis(string(code), "")

	for _, node := range results.ClassNodes {
		fmt.Println(node)
	}
	g.Expect(len(results.ClassNodes)).To(Equal(1))
	g.Expect(results.ClassNodes[0].Class).To(Equal("Employee"))
}

func Test_ShouldEnableGetClassMethod(t *testing.T) {

	g := NewGomegaWithT(t)

	app := new(TypeScriptApiApp)

	codefile := app.Analysis(`
class Employee  {
    displayName():void {
        console.log("hello, world");
    }
}
`, "")

	g.Expect(len(codefile.DataStructures[0].Functions)).To(Equal(1))
	g.Expect(len(codefile.ClassNodes[0].Methods)).To(Equal(1))
}

func Test_ShouldGetInterfaceImplements(t *testing.T) {

	g := NewGomegaWithT(t)

	app := new(TypeScriptApiApp)

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
`, "").ClassNodes

	g.Expect(results[1].Extend).To(Equal("IPerson"))
}

func Test_ShouldGetInterfaceProperty(t *testing.T) {

	g := NewGomegaWithT(t)

	app := new(TypeScriptApiApp)
	results := app.Analysis(`
export interface IPerson {
    name: string;
    gender: string;
    getSalary: (number) => number;
    getManagerName(number): string;
}
`, "").ClassNodes

	firstMethod := results[0].Methods[0]
	secondMethod := results[0].Methods[1]

	g.Expect(len(results[0].Fields)).To(Equal(2))
	g.Expect(len(results[0].Methods)).To(Equal(2))
	g.Expect(firstMethod.Name).To(Equal("getSalary"))
	g.Expect(secondMethod.Name).To(Equal("getManagerName"))
	g.Expect(secondMethod.Parameters[0].Type).To(Equal("number"))
}

func Test_ShouldGetDefaultFunctionName(t *testing.T) {

	g := NewGomegaWithT(t)

	app := new(TypeScriptApiApp)

	results := app.Analysis(`
function Sum(x: number, y: number) : void {
    console.log('processNumKeyPairs: key = ' + key + ', value = ' + value)
    return x + y;
}
`, "").ClassNodes

	firstMethod := results[0].Methods[0]
	parameters := firstMethod.Parameters
	g.Expect(firstMethod.Type).To(Equal("void"))
	g.Expect(len(parameters)).To(Equal(2))
	g.Expect(parameters[0].Name).To(Equal("x"))
	g.Expect(parameters[0].Type).To(Equal("number"))
}

func Test_ShouldHandleRestParameters(t *testing.T) {

	g := NewGomegaWithT(t)

	app := new(TypeScriptApiApp)

	results := app.Analysis(`
function buildName(firstName: string, ...restOfName: string[]) {
  return firstName + " " + restOfName.join(" ");
}
`, "").ClassNodes

	firstMethod := results[0].Methods[0]
	parameters := firstMethod.Parameters
	g.Expect(len(parameters)).To(Equal(2))
	g.Expect(parameters[0].Name).To(Equal("firstName"))
	g.Expect(parameters[1].Name).To(Equal("restOfName"))
}

func Test_ShouldGetClassFields(t *testing.T) {
	g := NewGomegaWithT(t)

	app := new(TypeScriptApiApp)
	code, _ := ioutil.ReadFile("../../../../_fixtures/ts/grammar/Class.ts")

	results := app.Analysis(string(code), "").ClassNodes

	fields := results[1].Fields
	g.Expect(len(fields)).To(Equal(5))
	g.Expect(fields[0].Modifier).To(Equal("public"))
}

func Test_ShouldReturnBlockImports(t *testing.T) {

	g := NewGomegaWithT(t)

	app := new(TypeScriptApiApp)

	results := app.Analysis(`
import { ZipCodeValidator } from "./ZipCodeValidator";

`, "")

	g.Expect(len(results.Imports)).To(Equal(1))
	g.Expect(results.Imports[0].Source).To(Equal("./ZipCodeValidator"))
}

func Test_ShouldReturnAsImports(t *testing.T) {

	g := NewGomegaWithT(t)

	app := new(TypeScriptApiApp)

	results := app.Analysis(`
import zip = require("./ZipCodeValidator");
`, "")

	g.Expect(len(results.Imports)).To(Equal(1))

	g.Expect(results.Imports[0].Source).To(Equal("./ZipCodeValidator"))
}

// Todo: fix for $ and *
func Test_ShouldReturnAllImports(t *testing.T) {

	g := NewGomegaWithT(t)

	app := new(TypeScriptApiApp)
	code, _ := ioutil.ReadFile("../../../../_fixtures/ts/grammar/Import.ts")
	results := app.Analysis(string(code), "")

	g.Expect(len(results.Imports)).To(Equal(5))
}
