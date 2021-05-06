package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func startHttpServer(wg *sync.WaitGroup) *http.Server {
	server := &http.Server{Addr: ":8080"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world")
	})

	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Printf("error: %v\n", err)
		}

		fmt.Println("server exit")
	}()

	return server
}

func main() {
	fmt.Println("Start")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	var wg sync.WaitGroup
	wg.Add(1)
	server := startHttpServer(&wg)

	s := <-sig
	switch s {
	case syscall.SIGTERM:
		fmt.Println("Received SIGTERM")
		wg.Done()
	case syscall.SIGINT:
		fmt.Println("Received SIGINT")
		wg.Done()
	}

	wg.Wait()
	if err := server.Shutdown(context.Background()); err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Println("End")
}
