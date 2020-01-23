package core_domain

type CodeDependency struct {
	GroupId    string
	ArtifactId string
	Scope      string
	Type       string
	Version    string
	Optional   bool
}

func NewCodeDependency(group string, artifact string) *CodeDependency {
	return &CodeDependency{
		GroupId:    group,
		ArtifactId: artifact,
		Optional:   false,
	}
}
