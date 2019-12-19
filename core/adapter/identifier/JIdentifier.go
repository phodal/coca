package identifier

import "coca/core/models"

var methods []models.JMethod

type JIdentifier struct {
	Package     string
	Name        string
	Type        string
	ExtendsName string
	Extends     []JIdentifier
	Methods     []models.JMethod
}

func NewJIdentifier() *JIdentifier {
	identifier := &JIdentifier{"", "", "", "", nil, nil}
	methods = nil
	return identifier
}

func (identifier *JIdentifier) AddMethod(method models.JMethod) {
	methods = append(methods, method)
}

func (identifier *JIdentifier) GetMethods() []models.JMethod {
	return methods
}
