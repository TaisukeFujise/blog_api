package middlewares

import (
	"sync"
)

var (
	logNo int = 1
	mu    sync.Mutex
)

// type traceIDKey struct{}

func newTraceID() int {
	var no int

	mu.Lock()
	no = logNo
	logNo += 1
	mu.Unlock()

	return no
}

// func SetTraceID(ctx context.Context, traceID int) context.Context {
// 	return context.WithValue(ctx, traceIDKey{}, traceID) // WithValue(ctx, key, value)
// }

// func GetTraceID(ctx context.Context) int {
// 	id := ctx.Value(traceIDKey{}) // ctx.Value(key)

// 	if idInt, ok := id.(int); ok {
// 		return idInt
// 	}
// 	return 0
// }
