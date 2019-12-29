package common_listener

import (
	"github.com/phodal/coca/core/models"
	"github.com/phodal/coca/languages/java"
)

func BuildAnnotation(ctx *parser.AnnotationContext) models.Annotation {
	annotationName := ctx.QualifiedName().GetText()
	annotation := models.NewAnnotation()
	annotation.QualifiedName = annotationName
	if ctx.ElementValuePairs() != nil {
		pairs := ctx.ElementValuePairs().(*parser.ElementValuePairsContext).AllElementValuePair()
		for _, pair := range pairs {
			pairCtx := pair.(*parser.ElementValuePairContext)
			pairCtx.ElementValue()
			annotation.ValuePairs = append(annotation.ValuePairs, *&models.AnnotationKeyValue{
				Key:   pairCtx.IDENTIFIER().GetText(),
				Value: pairCtx.ElementValue().GetText(),
			})
		}
	}

	if ctx.ElementValue() != nil {
		value := ctx.ElementValue().GetText()
		annotation.ValuePairs = append(annotation.ValuePairs, *&models.AnnotationKeyValue{
			Key:   value,
			Value: value,
		})
	}

	return annotation
}

