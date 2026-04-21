package middlewares

import (
	"log"
	"net/http"
)

// Goのミドルウェアの形は以下のようにhttp.Handlerを受け取ってhttp.Handlerを返す
// func myMiddleware(next http.Handler) http.Handler

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// リクエスト情報をロギング
		log.Println(req.RequestURI, req.Method)

		rlw := NewResLoggingWriter(w)
		next.ServeHTTP(rlw, req)
		log.Println(rlw.code)
	})
}

type resLoggingWriter struct {
	http.ResponseWriter
	code int
}

func (rsw *resLoggingWriter) WriteHeader(code int) {
	rsw.code = code
	rsw.ResponseWriter.WriteHeader(code)
}

func NewResLoggingWriter(w http.ResponseWriter) *resLoggingWriter {
	return &resLoggingWriter{ResponseWriter: w, code: http.StatusOK}
}
