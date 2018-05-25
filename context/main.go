package main

import (
	"context"
	"net/http"
	"sync"
)

//context包实现http request 传值 超时 取消
//golang使用context管理管理goroutine
func HTTPAPIServe(ctx context.Context) {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	ctx, cancel := context.WithCancel(ctx)
	wg := &sync.WaitGroup{}
	defer wg.Wait()
	wg.Add(1)
	var server *http.Server
	go func(ctx context.Context) {
		defer wg.Done()
		defer cancel()
		server = &http.Server{Addr: "127.0.0.1:8000", Handler: nil}
		_ = server.ListenAndServe()
	}(ctx)
	select {
	case <-ctx.Done():
		server.Close()

	}
}
func main() {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		defer cancel()
		HTTPAPIServe(ctx)
	}(ctx)

	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		defer cancel()
		//todo
	}(ctx)

}
