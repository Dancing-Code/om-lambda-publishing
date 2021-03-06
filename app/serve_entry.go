package main

import (
	"github.com/bloom42/rz-go/log"
	"net/http"
)

func main() {
	// Logging開始
	initLogger()

	log.Info("server started.")

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
	// 	server := &http.Server{
	// 		Addr:           "0.0.0.0:8080",
	// 		Handler:        mux,
	// 		MaxHeaderBytes: 1 << 20,
	// 	}
	//  server.ListenAndServe()

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

	// log.Printf("Handling HTTP requests on %s.", port)
	// log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
