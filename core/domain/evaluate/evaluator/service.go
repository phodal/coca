package evaluator

import (
	"github.com/phodal/coca/core/models"
)

type Service struct {
}

func (s Service) Evaluate(models.JClassNode) {
	if s.hasLifeCycle() {

	}
	if s.hasSameBehavior() {

	}
	if s.hasModelLike() {

	}
	if s.hasSameReturnType() {

	}
}

func (s Service) hasLifeCycle() bool {
	return false
}

func (s Service) hasSameBehavior() bool {
	return false
}

func (s Service) hasModelLike() bool {
	return false
}

func (s Service) hasSameReturnType() bool {
	return false
}
