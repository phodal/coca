package jdomain

type Annotation struct {
	QualifiedName string
	ValuePairs    []AnnotationKeyValue
}

type AnnotationKeyValue struct {
	Key   string
	Value string
}

func NewAnnotation() Annotation {
	return Annotation{
		QualifiedName: "",
		ValuePairs:    nil,
	}
}

func (n * Annotation) IsComponentOrRepository() bool {
	return n.QualifiedName == "Component" || n.QualifiedName == "Repository"
}

func (n * Annotation) IsTest() bool {
	return n.QualifiedName == "Test"
}

func (n * Annotation) IsIgnoreTest() bool {
	return n.QualifiedName == "Ignore"
}

func (n * Annotation) IsIgnoreOrTest() bool {
	return n.IsTest() || n.IsIgnoreTest()
}