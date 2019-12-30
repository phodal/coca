package domain

type Annotation struct {
	QualifiedName string
	ValuePairs    []AnnotationKeyValue
}

type AnnotationKeyValue struct {
	Key   string
	Value string
}

func NewAnnotation() Annotation {
	return *&Annotation{
		QualifiedName: "",
		ValuePairs:    nil,
	}
}

