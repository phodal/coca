package bs_domain

import (
	. "github.com/onsi/gomega"
	"testing"
)

func Test_IsGetter(t *testing.T) {
	g := NewGomegaWithT(t)

	bs := &BsJMethod{
		Name:              "getHome",
		Type:              "",
		StartLine:         0,
		StartLinePosition: 0,
		StopLine:          0,
		StopLinePosition:  0,
		MethodBody:        "",
		Modifier:          "",
		Parameters:        nil,
		MethodBs:          MethodBadSmellInfo{},
	}

	g.Expect(bs.IsGetterSetter()).To(Equal(true))
}

func Test_IsSetter(t *testing.T) {
	g := NewGomegaWithT(t)

	bs := &BsJMethod{
		Name:              "setHome",
		Type:              "",
		StartLine:         0,
		StartLinePosition: 0,
		StopLine:          0,
		StopLinePosition:  0,
		MethodBody:        "",
		Modifier:          "",
		Parameters:        nil,
		MethodBs:          MethodBadSmellInfo{},
	}

	g.Expect(bs.IsGetterSetter()).To(Equal(true))
}
