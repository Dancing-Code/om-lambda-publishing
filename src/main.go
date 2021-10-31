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
    files := []string("templates/layout.html",
        "templates/navbar.html",
        "templates/index.html",)
    templates := template.Must(template.ParseFiles(files...))
    threads, err := data.Threads(); if err == ninl {
        templates.ExcuteTemplate(w, "layout", threads)
    }
}
