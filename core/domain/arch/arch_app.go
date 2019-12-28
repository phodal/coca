package arch

import "github.com/phodal/coca/core/models"

type ArchApp struct {
}

func NewArchApp() ArchApp {
	return *&ArchApp{}
}

func (a ArchApp) Analysis(deps []models.JClassNode) string {
	return ""
}
