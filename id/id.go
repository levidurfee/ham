package id

import (
	"context"
	"log"
	"math/rand"
)

type key int

const k = key(0)

// CtxWithID adds an ID to the context for each request.
func CtxWithID(ctx context.Context) context.Context {
	return context.WithValue(ctx, k, rand.Int63())
}

// GetID retrieves the ID using the const k since it isn't exported, only funcs
// in this package can access the value.
func GetID(ctx context.Context) int64 {
	id, ok := ctx.Value(k).(int64)
	if !ok {
		log.Println("Could not get ID from context.")
		return 0
	}

	return id
}

// PrintID will display the ID in the logs
func PrintID(ctx context.Context) {
	id := GetID(ctx)

	log.Printf("Request ID [%d]\n", id)
}
