package hamlog

import (
	"context"
	"math/rand"
)

type key int64

var k key

// CtxWithID adds an ID to the context for each request.
func CtxWithID(ctx context.Context) context.Context {
	return context.WithValue(ctx, k, rand.Int63())
}
