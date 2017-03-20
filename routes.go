package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Login Email",
		"POST",
		"/auth/email/login",
		LoginEmail,
	},
	Route{
		"Login Facebook",
		"POST",
		"/auth/facebook/login",
		LoginFacebook,
	},
	Route{
		"Callback Facebook",
		"GET",
		"/auth/email/login",
		CallbackFacebook,
	},
	Route{
		"Login Google",
		"POST",
		"/auth/google/callback",
		LoginGoogle,
	},
	Route{
		"Callback Google",
		"GET",
		"/auth/google/callback",
		CallbackGoogle,
	},
	Route{
		"Login Google",
		"POST",
		"/auth/email/register",
		RegisterEmail,
	},
}
