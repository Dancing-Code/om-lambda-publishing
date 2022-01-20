// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/bloom42/rz-go"
	"github.com/bloom42/rz-go/log"
	"net/http"
)

func main() {
	// Logging開始
	initLogger()

	log.Info("info from logger")

	mux := http.NewServeMux()

	files := http.FileServer(http.Dir("pages/styles"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", topPage)

	server := &http.Server{
		Addr:           "0.0.0.0:8080",
		Handler:        mux,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()

    /*
    サーバーレス用
    */
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Cloud Build Testing. Lambda Publishing LTD.")
	//})

	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "8080"
	// }

	// log.Printf("Handling HTTP requests on %s.", port)
	// log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}