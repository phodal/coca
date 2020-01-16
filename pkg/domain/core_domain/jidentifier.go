package core_domain

type JIdentifier struct {
	Package     string
	ClassName   string
	ClassType   string
	ExtendsName string
	Extends     []string
	Implements  []string
	Methods     []CodeFunction
	Annotations []CodeAnnotation
}

func NewJIdentifier() *JIdentifier {
	return &JIdentifier{}
}

func (identifier *JIdentifier) GetClassFullName() string {
	return identifier.Package + "." + identifier.ClassName
}

func BuildIdentifierMap(identifiers []JIdentifier) map[string]JIdentifier {
	var identifiersMap = make(map[string]JIdentifier)

	for _, ident := range identifiers {
		identifiersMap[ident.Package+"."+ident.ClassName] = ident
	}
	return identifiersMap
}

func BuildDIMap(identifiers []JIdentifier, identifierMap map[string]JIdentifier) map[string]string {
	var diMap = make(map[string]string)
	for _, clz := range identifiers {
		if len(clz.Annotations) > 0 {
			for _, annotation := range clz.Annotations {
				if (annotation.IsComponentOrRepository()) && len(clz.Implements) > 0 {
					superClz := identifierMap[clz.Implements[0]]
					diMap[superClz.GetClassFullName()] = superClz.GetClassFullName()
				}
			}
		}
	}

	return diMap
}
