package main

import (
	"net/http"
)

func topPage(writer http.ResponseWriter, request *http.Request) {

	threads, err := db.Threads()

	generateHTML(writer, threads, "layout", "private.navbar", "index")
}
