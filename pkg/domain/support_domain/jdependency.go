package api_domain

type JDependency struct {
	GroupId    string
	ArtifactId string
	Scope      string
	Type       string
	Version    string
	Optional   bool
}

func NewJDependency(group string, artifact string) *JDependency {
	return &JDependency{
		GroupId:    group,
		ArtifactId: artifact,
		Scope:      "",
		Type:       "",
		Version:    "",
		Optional:   false,
	}
}
