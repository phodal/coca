package core_domain

type AnnotationKeyValue struct {
	Key   string
	Value string
}

func NewAnnotationKeyValue(key string, value string) AnnotationKeyValue {
	return AnnotationKeyValue{
		Key:   key,
		Value: value,
	}
}

type CodeAnnotation struct {
	Name      string
	KeyValues []AnnotationKeyValue
}

func NewAnnotation() CodeAnnotation {
	return CodeAnnotation{
		Name:      "",
		KeyValues: nil,
	}
}

func (n *CodeAnnotation) IsComponentOrRepository() bool {
	return n.Name == "Component" || n.Name == "Repository"
}

func (n *CodeAnnotation) IsTest() bool {
	return n.Name == "Test"
}

func (n *CodeAnnotation) IsIgnoreTest() bool {
	return n.Name == "Ignore"
}

func (n *CodeAnnotation) IsIgnoreOrTest() bool {
	return n.IsTest() || n.IsIgnoreTest()
}