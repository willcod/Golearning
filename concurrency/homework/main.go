package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
)

func startHttpServer(wg *sync.WaitGroup) *http.Server {
	server := &http.Server{Addr: ":8080"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world")
	})

	go func() {
		defer wg.Done()

		if err := server.ListenAndServe(); err != nil {
			fmt.Printf("error: %v\n", err)
		}

		fmt.Println("server exit")
	}()

	return server
}

func main() {
	fmt.Println("Start")
	var wg sync.WaitGroup
	wg.Add(1)
	server := startHttpServer(&wg)

	wg.Wait()
	if err := server.Shutdown(context.Background()); err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Println("End")
}
