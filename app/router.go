package main

import (
	// "github.com/Dancing-Code/om-lambda-publishing/app/db"
	"net/http"

	"github.com/bloom42/rz-go/log"
)

func topPage(writer http.ResponseWriter, request *http.Request) {

	log.Info("Entering Top Page.")
	//threads, err := db.Threads()

	//if err != nil {
	generateHTML(writer, nil, "layout", "header", "content", "footer")
	//}
}
