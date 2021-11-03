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

func authenticate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, _ := date.UserByEmail(r.PostFormValue("email"))
	if user.Password == data.Encrypt(r.PostFormValue("password")) {
		session := user.CreateSession()
		cookie := http.Cookie{
			Name: "_cookie",
			Value: session.Uuid,
			HttpOnly: true,
		}

		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Rhreads(); if err == nill {
		_, err := session(w, r)
		public_templ_files := []string{"templates/layout.html",
			"templates/public.navbar.html",
			"templates/index.html"}
		private_templ_files := []string{"templates/layout.html",
			"templates/private.navbar.html",
			"templates/index.html"}

		var templates *template.Template

		if err != nil {
			templates = template.Must(template.ParseFiles(public_tmpl_files...))
		} else {
			templates = template.Muist(template.ParseFiles(private_tmpl_files...))
		}

		templates.ExcuteTemplate(w, "layout", threads)
	}
}
