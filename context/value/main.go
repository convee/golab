package main

import (
	"context"
	"fmt"
	"math/rand"
)

func process(ctx context.Context) {
	ret, ok := ctx.Value("trace_id").(int)
	if !ok {
		traceId := rand.Intn(100)
		ret = traceId
	}
	s, _ := ctx.Value("session").(string)
	fmt.Println("ret:", ret)
	fmt.Println("s:", s)

}
func main() {
	traceId := rand.Intn(100)
	ctx := context.WithValue(context.Background(), "trace_id", traceId)
	ctx = context.WithValue(ctx, "session", "asfsf")
	process(ctx)

}
