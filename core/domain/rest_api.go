package domain

type RestApi struct {
	Uri              string
	HttpMethod       string
	MethodName       string
	ResponseStatus   string
	RequestBodyClass string
	MethodParams     map[string]string
	PackageName      string
	ClassName        string
}

func (r *RestApi) BuildFullMethodPath() string {
	return r.PackageName + "." + r.ClassName + "." + r.MethodName
}

