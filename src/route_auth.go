package main

import (
	"github.com/behike56/om-lambda-publishing/data"
	"net/http"
)

// /login
func login(writer http.ResponseWriter, request *http.Request) {
	t := parseTemplateFiles("login.layout", "public.navbar", "login")
	t.Excute(writer, nil)
}
