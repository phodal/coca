package domain

import "coca/src/adapter/models"


type CallGraph struct {

}

func NewCallGraph() CallGraph {
	return *&CallGraph{}
}

func (c CallGraph) Analysis(path string, clzs *[]models.JClassNode) {

}