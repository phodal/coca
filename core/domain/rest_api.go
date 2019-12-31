package domain

import "strings"

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

func FilterApiByPrefix(apiPrefix string, apis []RestApi, ) []RestApi {
	var restFieldsApi []RestApi
	if apiPrefix != "" {
		for _, api := range apis {
			if strings.HasPrefix(api.Uri, apiPrefix) {
				restFieldsApi = append(restFieldsApi, api)
			}
		}
	} else {
		restFieldsApi = apis
	}

	return restFieldsApi
}
