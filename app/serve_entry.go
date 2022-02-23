package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bloom42/rz-go/log"
)

func main() {
	// Logging開始
	initLogger()

	log.Info("Server Initiatializing......")

	mux := http.NewServeMux()

	files := http.FileServer(http.Dir("pages"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", topPage)
	mux.HandleFunc("/lexical-corp-info", corpInfo)
	mux.HandleFunc("/disclaimer", disclaimer)
	mux.HandleFunc("/contact", contact)

	/*
	   開発用
	*/
	// server := &http.Server{
	// 	Addr:           os.Getenv("PORT"),
	// 	Handler:        mux,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	// err := server.ListenAndServe()
	// if err == nil {
	// 	log.Info("Server Started Normaly.")
	// }
	/*
	   サーバーレス用
	*/
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Cloud Build Testing. Lambda Publishing LTD.")
	//})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	//	server.ListenAndServe()
	http.ListenAndServe(fmt.Sprintf(":%s", port), mux)

	//log.Printf("Handling HTTP requests on %s.", port)
	//log.Fatal\\(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))

}
