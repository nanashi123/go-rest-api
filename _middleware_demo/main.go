package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	hellowHandler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	})

	//http.Handle("/", hellowHandler)
	//http.Handle("/", myMiddleware1(hellowHandler))
	http.Handle("/", myMiddleware2(myMiddleware1(hellowHandler)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// ハンドラを受け取って、新しいハンドラを作る
func myMiddleware1(next http.Handler) http.Handler {
	// ハンドラ関数func(w jhttp.ResoponseWriter, r *http.Request)をhttp.HandlerFunc型にキャストすることで、戻り値であるhttp.Handlerインターフェースを満たすようにしている
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 戻り値にしたい新しいハンドラの中身

		// 元のハンドラを実行する前に行いたい前処理
		io.WriteString(w, "Pre-process1\n")

		// 元のハンドラを実行
		next.ServeHTTP(w, r)

		// 元のハンドラを実行した後に行いたい後処理
		io.WriteString(w, "Post-process1\n")
	})
}

func myMiddleware2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Pre-process2\n")
		next.ServeHTTP(w, r)
		io.WriteString(w, "Post-process2\n")
	})
}
