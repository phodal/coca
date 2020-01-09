package js_ident

import (
	. "github.com/onsi/gomega"
	"testing"
)

func Test_TypeScriptConsoleLog(t *testing.T) {
	g := NewGomegaWithT(t)

	app := new(TypeScriptApiApp)
	results := app.Analysis("console.log('hello, world')")

	g.Expect(len(results.MethodCalls)).To(Equal(1))
	g.Expect(results.MethodCalls[0].Class).To(Equal("console"))
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

	g.Expect(results.Class).To(Equal("Person"))
}