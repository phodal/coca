package js_ident

import (
	"fmt"
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

	g.Expect(results[0].Class).To(Equal("Person"))
	g.Expect(results[0].Implements[0]).To(Equal("IPerson"))
}

func Test_TypeScriptMultipleClass(t *testing.T) {
	g := NewGomegaWithT(t)

	app := new(TypeScriptApiApp)
	code, _ := ioutil.ReadFile("../../../../_fixtures/ts/Class.ts")

	results := app.Analysis(string(code))

	g.Expect(len(results)).To(Equal(2))
	g.Expect(results[0].Implements[0]).To(Equal("IPerson"))
}