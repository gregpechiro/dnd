package main

import "github.com/cagnosolutions/web"

var ADMIN = web.Auth{
	Roles:    []string{"ADMIN"},
	Redirect: "/login",
	Msg:      "You Are not authorized",
}

var USER = web.Auth{
	Roles:    []string{"ADMIN", "USER"},
	Redirect: "/login",
	Msg:      "Please login",
}
