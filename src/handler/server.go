package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

func welcome(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Welcome!")
}

func log(handler http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
		fmt.Println("Called Handler-Function - " + name)
		handler(writer, request)
	}
}

func main(){
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/welcome", log(welcome))
	server.ListenAndServe()
}
