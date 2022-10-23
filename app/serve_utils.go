package main

import (
	"fmt"
	"html/template"
	"net/http"
)

/*
基本的なページ要素のHTMLを生成する。
*/
func generateBaseHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string

	for _, file := range filenames {
		files = append(files, fmt.Sprintf("pages/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}

/*
コンテンツページのHTMLを生成する。
*/
func generateContentsHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string

	for _, file := range filenames {
		files = append(files, fmt.Sprintf("pages/templates/%s.html", file))
	}

	for _, file := range filenames {
		files = append(files, fmt.Sprintf("pages/contents/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}
