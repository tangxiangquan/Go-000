package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
)

func serverStart(addr string,cancelFunc context.CancelFunc) error{
	defer func() {
		cancelFunc()
	}()
	server := &http.Server{ Addr:addr}

	return server.ListenAndServe()
}

func test()  {
	ctx , cancel := context.WithCancel(context.Background())
	g ,_ := errgroup.WithContext(ctx)

	g.Go(func() error {
		return serverStart(":8080" , cancel)
	})

	g.Go(func() error {
		return serverStart(":8081" , cancel)
	})

	err := g.Wait()
	log.Println("server shut down ....." ,err)
}

func main() {
	test()
}
