package common_listener

import (
	"github.com/phodal/coca/languages/java"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"github.com/phodal/coca/pkg/domain/jdomain"
	"reflect"
)

func BuildAnnotation(ctx *parser.AnnotationContext) jdomain.Annotation {
	annotationName := ctx.QualifiedName().GetText()
	annotation := jdomain.NewAnnotation()
	annotation.QualifiedName = annotationName
	if ctx.ElementValuePairs() != nil {
		pairs := ctx.ElementValuePairs().(*parser.ElementValuePairsContext).AllElementValuePair()
		for _, pair := range pairs {
			pairCtx := pair.(*parser.ElementValuePairContext)
			pairCtx.ElementValue()

			key := pairCtx.IDENTIFIER().GetText()
			value := pairCtx.ElementValue().GetText()
			annotation.ValuePairs = append(annotation.ValuePairs, core_domain.NewAnnotationKeyValue(key, value))
		}
	} else if ctx.ElementValue() != nil {
		value := ctx.ElementValue().GetText()
		annotation.ValuePairs = append(annotation.ValuePairs, core_domain.NewAnnotationKeyValue(value, value))
	}

	return annotation
}

func BuildAnnotationForMethod(context *parser.ModifierContext, method *jdomain.JMethod) {
	if context.ClassOrInterfaceModifier() != nil {
		if reflect.TypeOf(context.ClassOrInterfaceModifier().GetChild(0)).String() == "*parser.AnnotationContext" {
			annotationCtx := context.ClassOrInterfaceModifier().GetChild(0).(*parser.AnnotationContext)

			annotation := BuildAnnotation(annotationCtx)
			method.Annotations = append(method.Annotations, annotation)
		}
	}
}

