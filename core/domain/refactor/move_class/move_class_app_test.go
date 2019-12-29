package move_class

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestNewMoveClassApp(t *testing.T) {
	g := NewGomegaWithT(t)

	config := "../../../../_fixtures/refactor/move.config"
	path := "../../../../_fixtures/refactor/"
	app := NewMoveClassApp(config, path)
	app.Analysis()

	g.Expect(true).To(Equal(true))
}