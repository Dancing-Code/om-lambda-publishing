package main

import (
	"fmt"
	"net/http"
)

func main() {
	// マルチプレクサの生成
	mux := http.NewServeMux()

	// 静的ファイルの返送
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", hhtp.StripPrefix("/staic/", files))

	// ルートURLをハンドラ関数にリダイレクトする
	// 第１引数：URL、第２引数：ハンドラ関数名
	mux.HandleFunc("/", index)
	mux.HandleFunc("/err", err)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	// threads method
	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	server := &http.server{
		Addr: "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()

}
