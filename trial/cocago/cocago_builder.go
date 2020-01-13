package cocago

import (
	"github.com/phodal/coca/pkg/domain/trial"
	"go/ast"
)

func AddStructType(currentStruct trial.CodeDataStruct, x *ast.StructType, currentFile *trial.CodeFile) {
	member := trial.CodeMember{
		DataStructID: currentStruct.Name,
		Type:         "struct",
	}
	for _, field := range x.Fields.List {
		typeName, typeType := BuildPropertyField(field)
		property := trial.CodeProperty{
			Modifiers: nil,
			Name:      getFieldName(field),
			TypeType:  typeType,
			TypeName:  typeName,
		}
		member.FileID = currentFile.FullName
		currentStruct.Properties = append(currentStruct.Properties, property)
	}
	currentFile.Members = append(currentFile.Members, &member)
	currentFile.DataStructures = append(currentFile.DataStructures, currentStruct)
}

