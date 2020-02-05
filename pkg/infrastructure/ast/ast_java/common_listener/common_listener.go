package common_listener

import (
	"github.com/phodal/coca/languages/java"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"reflect"
)

func BuildAnnotation(ctx *parser.AnnotationContext) core_domain.CodeAnnotation {
	annotationName := ctx.QualifiedName().GetText()
	annotation := core_domain.NewAnnotation()
	annotation.Name = annotationName
	if ctx.ElementValuePairs() != nil {
		pairs := ctx.ElementValuePairs().(*parser.ElementValuePairsContext).AllElementValuePair()
		for _, pair := range pairs {
			pairCtx := pair.(*parser.ElementValuePairContext)

			key := pairCtx.IDENTIFIER().GetText()
			value := pairCtx.ElementValue().GetText()
			annotation.KeyValues = append(annotation.KeyValues, core_domain.NewAnnotationKeyValue(key, value))
		}
	} else if ctx.ElementValue() != nil {
		value := ctx.ElementValue().GetText()
		annotation.KeyValues = append(annotation.KeyValues, core_domain.NewAnnotationKeyValue(value, value))
	}

	return annotation
}

func BuildAnnotationForMethod(context *parser.ModifierContext, method *core_domain.CodeFunction) {
	if context.ClassOrInterfaceModifier() != nil {
		if reflect.TypeOf(context.ClassOrInterfaceModifier().GetChild(0)).String() == "*parser.AnnotationContext" {
			annotationCtx := context.ClassOrInterfaceModifier().GetChild(0).(*parser.AnnotationContext)

			annotation := BuildAnnotation(annotationCtx)
			method.Annotations = append(method.Annotations, annotation)
		}
	}
}
