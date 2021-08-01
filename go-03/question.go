// 1. 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。
package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("pong"))
	})

	srv := http.Server{
		Handler: mux,
		Addr:    ":7777",
	}

	// http server start
	g.Go(func() error {
		fmt.Println("http server start")
		return srv.ListenAndServe()
	})

	// http server shutdown
	g.Go(func() error {
		<-ctx.Done()
		fmt.Println("http server shutdown")
		return srv.Shutdown(ctx)
	})

	// linux signal
	g.Go(func() error {
		// 创建监听处理 buffer channel
		c := make(chan os.Signal, 1)
		// 监听所有信号
		signal.Notify(c)

		for {
			select {
			case <-ctx.Done():
				fmt.Println("http ctx done")
				return ctx.Err()
			case s := <-c:
				return errors.Errorf("get os signal: %v", s)
			}
		}
	})

	err := g.Wait()
	if err != nil {
		fmt.Println(err)
	}
}
