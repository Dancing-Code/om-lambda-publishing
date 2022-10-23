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
	generateBaseHTML(writer, nil, "layout", "header", "content", "footer")
	//}
}

func corpInfo(writer http.ResponseWriter, request *http.Request) {
	log.Info("Entering Lexical corp info page.")
	generateBaseHTML(writer, nil, "layout", "header", "lexical-corp-info", "footer")
}

func disclaimer(writer http.ResponseWriter, request *http.Request) {
	log.Info("Entering Disclaimer page.")
	generateBaseHTML(writer, nil, "layout", "header", "disclaimer", "footer")
}

func contact(writer http.ResponseWriter, request *http.Request) {
	log.Info("Entering Contact page.")
	generateBaseHTML(writer, nil, "layout", "header", "contact", "footer")
}

func topOne(writer http.ResponseWriter, request *http.Request) {
	log.Info("Entering Contents page.")
	generateContentsHTML(writer, nil, "templates/layout", "templates/header", "contents/top_one", "templates/footer")
}
