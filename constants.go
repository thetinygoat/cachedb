package main

import "net/url"

type Route struct {
	Path      string
	variables map[string]string
}

// Constants used throughout the project
const (
	MalformedParams = "MALFORMED_PARAMS"
	KeyDoesNotExist = "KEY_DOES_NOT_EXIST_ERROR"
	KeyExpired      = "KEY_EXPIRED_ERROR"
	OK              = "OK"

	DefaultPort = "9898"
)

var GetRoute = Route{
	Path: "/get",
	variables: map[string]string{
		"key": "key",
	},
}

var SetRoute = Route{
	Path: "/set",
	variables: map[string]string{
		"key":    "key",
		"value":  "value",
		"ttlRaw": "ttl",
	},
}

var DelRoute = Route{
	Path: "/del",
	variables: map[string]string{
		"key": "key",
	},
}

var FlushRoute = Route{
	Path:      "/flush",
	variables: map[string]string{},
}

func (route Route) getVars(values url.Values) map[string]string {
	result := make(map[string]string)
	for k, v := range route.variables {
		result[k] = values.Get(v)
	}
	return result
}
