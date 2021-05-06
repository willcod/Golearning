package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func startHttpServer() *http.Server {
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

	g, ctx := errgroup.WithContext(context.Background())
	server := startHttpServer()
	g.Go(func() error {

		select {
		case <-ctx.Done():
			fmt.Println("Canceled")
			return nil
		default:
			s := <-sig
			switch s {
			case syscall.SIGTERM:
				fmt.Println("Received SIGTERM")
				return fmt.Errorf("SIGTERM")

			case syscall.SIGINT:
				fmt.Println("Received SIGINT")
				return fmt.Errorf("SIGINT")
			}
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("%v\n", err)
	}

	if err := server.Shutdown(context.Background()); err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Println("End")
}
