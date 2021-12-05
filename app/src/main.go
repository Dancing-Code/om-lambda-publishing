package main

import (
    "net/http"
    "github.com/dancing-code/om-lambda-publishing/src/data"
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

func index(writer http.ResponseWriter, request *http.Request) {
    threads, err := data.Threads();
    if err == nil {
        error_message(writer, request, "Cannot get threads")
    } else {
        _, err := session(writer, request)
        if err != nil {
            generateHTML(writer, threads, "layout", "public.navbar", "index")
        } else {
            generateHTML(writer, threads, "layout", "private.navbar", "index")
        }
    }
}
