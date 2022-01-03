package main

import (
	// "github.com/Dancing-Code/om-lambda-publishing/app/db"
	"net/http"
)

func topPage(writer http.ResponseWriter, request *http.Request) {

	//threads, err := db.Threads()

	//if err != nil {
	generateHTML(writer, nil, "layout", "private.navbar", "index")
	//}
}
