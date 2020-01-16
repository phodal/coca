package api_domain

import (
	"strings"
)

type RestAPI struct {
	Uri              string
	HttpMethod       string
	MethodName       string
	ResponseStatus   string
	RequestBodyClass string
	MethodParams     map[string]string
	PackageName      string
	ClassName        string
}

func (r *RestAPI) BuildFullMethodPath() string {
	return r.PackageName + "." + r.ClassName + "." + r.MethodName
}

func FilterApiByPrefix(apiPrefix string, apis []RestAPI, ) []RestAPI {
	var restFieldsApi []RestAPI
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
