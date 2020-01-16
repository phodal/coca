package core_domain

var methods []JMethod

type JIdentifier struct {
	Package     string
	ClassName   string
	ClassType   string
	ExtendsName string
	Extends     []string
	Implements  []string
	Methods     []JMethod
	Annotations []CodeAnnotation
}

func NewJIdentifier() *JIdentifier {
	identifier := &JIdentifier{"", "", "", "", nil, nil, nil, nil}
	methods = nil
	return identifier
}

func (identifier *JIdentifier) AddMethod(method JMethod) {
	methods = append(methods, method)
}

func (identifier *JIdentifier) GetMethods() []JMethod {
	return methods
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
