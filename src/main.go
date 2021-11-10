package main

import (
    "fmt"
    "net/http"
    "os"
    "strconv"
    "html/template"
)

/* 開始起点  */
func main() {
    // マルチプレクサの生成
    mux := http.NewServeMux()

    // マルチプレクサで静的なファイルの返送にも使える
    files := http.FileServer(http.Dir("/public"))
    // URLパスからプレフィックスを削除するにはStripPrefixを使う。
    mux.Handle("/static/", http.StripPrefix("/static/", files))

    // ルートURLをハンドラ関数にリダイレクト
    mux.HandleFunc("/", index)

    server := &http.Server{
        Addr: "0.0.0.0:8080",
        Handler: mux,
    }

    server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
    threds, err := data.Threads(); if err == nil {
        _, err := session(writer, request)
        if err != nil {
            generateHTML(writer, threads, "layout", "public.navbar", "index")
        } else {
            generateHTML(writer, threads, "layout", "private.navbar", "index")
        }
    }
}
