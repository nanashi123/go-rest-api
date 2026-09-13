package middlewares

import (
	"log"
	"net/http"

	"github.com/Nanashi123/go-api-practice/common"
)

// 自作ResponseWriterを作る
type resLoggingWriter struct {
	http.ResponseWriter     // もともと使用していたhttp.ResponseWriterを格納するためのフィールド（フィールド名を省略しているので、型名と同じResPonseWrtierになる）
	code                int //ハンドラが使ったレスポンスコードを格納しておくためのフィールド
}

// コンストラクタを作る
// -> 内部フィールドに入れるResponseWriterを受け取ってresLoggingWriter構造体を作る
func NewResLoggingWriter(w http.ResponseWriter) *resLoggingWriter {
	return &resLoggingWriter{ResponseWriter: w, code: http.StatusOK}
}

// WriteHeaderメソッドを作る ハンドラがHTTPレスポンスコードを書き込むときに使うメソッド
func (rsw *resLoggingWriter) WriteHeader(code int) {
	// resLoggingWriter構造体のcodeフィールドに、使うレスポンスコードを保存する
	rsw.code = code

	// HTTPレスポンスに使うレスポンスコードを指定 (=WriteHeaderメソッド本来の機能を呼び出し)
	rsw.ResponseWriter.WriteHeader(code)
}

// ミドルウェアの中身
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		traceID := newTraceID()

		// リクエスト情報をロギング
		log.Printf("[%d]%s %s\n", traceID, req.RequestURI, req.Method)

		ctx := common.SetTraceID(req.Context(), traceID)
		req = req.WithContext(ctx)

		// 自作のResponseWriterを作って
		rlw := NewResLoggingWriter(w)

		// それをハンドラに渡す
		// -> メソッドの移譲によってrlwをServeHTTPの第一引数にできる
		next.ServeHTTP(rlw, req)

		// 自作ResponseWriterからロギングしたいデータを出す
		log.Printf("[%d]res: %d", traceID, rlw.code)
	})
}
