package common_listener

import (
	"github.com/phodal/coca/core/domain"
	"github.com/phodal/coca/languages/java"
	"reflect"
)

func BuildAnnotation(ctx *parser.AnnotationContext) domain.Annotation {
	annotationName := ctx.QualifiedName().GetText()
	annotation := domain.NewAnnotation()
	annotation.QualifiedName = annotationName
	if ctx.ElementValuePairs() != nil {
		pairs := ctx.ElementValuePairs().(*parser.ElementValuePairsContext).AllElementValuePair()
		for _, pair := range pairs {
			pairCtx := pair.(*parser.ElementValuePairContext)
			pairCtx.ElementValue()
			annotation.ValuePairs = append(annotation.ValuePairs, *&domain.AnnotationKeyValue{
				Key:   pairCtx.IDENTIFIER().GetText(),
				Value: pairCtx.ElementValue().GetText(),
			})
		}
	} else if ctx.ElementValue() != nil {
		value := ctx.ElementValue().GetText()
		annotation.ValuePairs = append(annotation.ValuePairs, *&domain.AnnotationKeyValue{
			Key:   value,
			Value: value,
		})
	}

	return annotation
}

func BuildAnnotationForMethod(context *parser.ModifierContext, method *domain.JMethod) {
	if context.ClassOrInterfaceModifier() != nil {
		if reflect.TypeOf(context.ClassOrInterfaceModifier().GetChild(0)).String() == "*parser.AnnotationContext" {
			annotationCtx := context.ClassOrInterfaceModifier().GetChild(0).(*parser.AnnotationContext)

			annotation := BuildAnnotation(annotationCtx)
			method.Annotations = append(method.Annotations, annotation)
		}
	}
}

