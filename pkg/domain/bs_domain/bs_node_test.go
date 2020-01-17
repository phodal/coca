package bs_domain

import (
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"testing"
)

func Test_IsGetter(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	position := core_domain.CodePosition{
		StartLine:         0,
		StartLinePosition: 0,
		StopLine:          0,
		StopLinePosition:  0,
	}

	bs := &BsJMethod{
		Name:       "getHome",
		Type:       "",
		MethodBody: "",
		Modifier:   "",
		Parameters: nil,
		MethodBs:   MethodBadSmellInfo{},
		Position:   position,
	}

	g.Expect(bs.IsGetterSetter()).To(Equal(true))
}

func Test_IsSetter(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	position := core_domain.CodePosition{
		StartLine:         0,
		StartLinePosition: 0,
		StopLine:          0,
		StopLinePosition:  0,
	}

	bs := &BsJMethod{
		Name:       "setHome",
		Type:       "",
		MethodBody: "",
		Modifier:   "",
		Parameters: nil,
		MethodBs:   MethodBadSmellInfo{},
		Position:   position,
	}

	g.Expect(bs.IsGetterSetter()).To(Equal(true))
}
