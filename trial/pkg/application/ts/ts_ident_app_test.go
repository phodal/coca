package js_ident

import (
	. "github.com/onsi/gomega"
	"io/ioutil"
	"testing"
)

func Test_TypeScriptConsoleLog(t *testing.T) {
	g := NewGomegaWithT(t)

	app := new(TypeScriptApiApp)
	results := app.Analysis("console.log('hello, world')")

	g.Expect(len(results[0].MethodCalls)).To(Equal(1))
	g.Expect(results[0].MethodCalls[0].Class).To(Equal("console"))
}

func Test_TypeScriptClassNode(t *testing.T) {
	g := NewGomegaWithT(t)

	app := new(TypeScriptApiApp)
	results := app.Analysis(`
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
`)

	g.Expect(results[0].Class).To(Equal("IPerson"))
	g.Expect(results[1].Class).To(Equal("Person"))
	g.Expect(results[1].Methods[0].Name).To(Equal("constructor"))
	g.Expect(results[1].Implements[0]).To(Equal("IPerson"))
}

func Test_TypeScriptMultipleClass(t *testing.T) {
	g := NewGomegaWithT(t)

	app := new(TypeScriptApiApp)
	code, _ := ioutil.ReadFile("../../../../_fixtures/ts/Class.ts")

	results := app.Analysis(string(code))

	g.Expect(len(results)).To(Equal(4))
	g.Expect(results[1].Implements[0]).To(Equal("IPerson"))
	g.Expect(results[3].Class).To(Equal("default"))
}

func Test_ShouldEnableGetClassMethod(t *testing.T) {
	g := NewGomegaWithT(t)

	app := new(TypeScriptApiApp)

	results := app.Analysis(`
class Employee  {
    displayName():void {
        console.log("hello, world");
    }
}
`)

	g.Expect(len(results[0].Methods)).To(Equal(1))
}

func Test_ShouldGetDefaultFunctionName(t *testing.T) {
	g := NewGomegaWithT(t)

	app := new(TypeScriptApiApp)

	results := app.Analysis(`
function Sum(x: number, y: number) : void {
    console.log('processNumKeyPairs: key = ' + key + ', value = ' + value)
    return x + y;
}
`)

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
`)

	firstMethod := results[0].Methods[0]
	parameters := firstMethod.Parameters
	g.Expect(len(parameters)).To(Equal(2))
	g.Expect(parameters[0].Name).To(Equal("firstName"))
	g.Expect(parameters[1].Name).To(Equal("restOfName"))
}
