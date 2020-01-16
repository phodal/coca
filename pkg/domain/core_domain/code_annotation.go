package core_domain

type CodeAnnotation struct {
	Name       string
	Properties []AnnotationKeyValue
}

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

