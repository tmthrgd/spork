package main

//go:generate go run -tags=dev internal/assets/generate.go
//go:generate go run -tags=dev internal/templates/generate.go

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"go.tmthrgd.dev/spork/dbus"
)

var shutdown = make(chan struct{})

func init() {
	log.SetFlags(0)
}

func main() {
	runtime.SetBlockProfileRate(5)
	runtime.SetMutexProfileFraction(5)

	addr := flag.String("addr", ":8080", "the address to listen on")
	flag.Parse()

	if err := dbus.BusConnect(); err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fmt.Printf("Listening on %s\n", *addr)

	srv := &http.Server{
		Addr:    *addr,
		Handler: handler(ctx),
	}

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// termination handler
	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	<-term

	// gracefull shutdown
	sctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	go func() {
		<-term
		signal.Stop(term)

		cancel()
	}()

	close(shutdown)

	if err := srv.Shutdown(sctx); err != nil {
		log.Printf("spork: error shutting down: %v", err)
	}
}
