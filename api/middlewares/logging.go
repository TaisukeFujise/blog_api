package middlewares

import (
	"log"
	"net/http"
)

// Goのミドルウェアの形は以下のようにhttp.Handlerを受け取ってhttp.Handlerを返す
// func myMiddleware(next http.Handler) http.Handler

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		traceID := newTraceID()

		// リクエスト情報をロギング
		// log.Println(req.RequestURI, req.Method)
		log.Printf("[%d]%s %s\n", traceID, req.RequestURI, req.Method)

		ctx := req.Context()
		ctx = SetTraceID(ctx, traceID)
		req = req.WithContext(ctx) // WithContextメソッドで、新しいコンテキストをreqに登録
		rlw := NewResLoggingWriter(w)
		next.ServeHTTP(rlw, req)
		// log.Println(rlw.code)
		log.Printf("[%d]res: %d", traceID, rlw.code)
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
