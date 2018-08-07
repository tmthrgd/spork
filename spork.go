package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/tmthrgd/httphandlers"
	"github.com/tmthrgd/spork/internal/dbus"
)

var (
	favicon = "\x89PNG\r\n\x1a\n\x00\x00\x00\rIHDR\x00\x00\x00\x10\x00\x00\x00\x10\b\x06\x00\x00\x00\x1f\xf3\xffa\x00\x00\x00\x16IDATx\xdacd\xa0\x100\x8e\x1a0j\xc0\xa8\x01\xc3\xc5\x00\x00&\x98\x00\x11\x9b\x92AZ\x00\x00\x00\x00IEND\xaeB`\x82"
	robots  = "User-agent: *\nDisallow: /"
)

var error404 = `<!doctype html>
<meta charset=utf-8>
<title>404 Not Found</title>
<style>body{margin:40px auto;max-width:650px;line-height:1.6;font-size:18px;color:#444;padding:0 10px}h1,h2,h3{line-height:1.2}</style>
<h1>404 Not Found</h1>
<p>The requested file was not found.</p>`

func main() {
	port := flag.Int("port", 8080, "the port to listen on")
	flag.Parse()

	bus, err := dbus.SessionBus()
	if err != nil {
		log.Fatal(err)
	}
	defer bus.Close()

	notFoundHandler := handlers.ServeError(http.StatusNotFound, []byte(error404), "text/html; charset=utf-8")

	router := chi.NewRouter()
	router.Use(middleware.GetHead)
	router.NotFound(notFoundHandler.ServeHTTP)

	now := time.Now()
	router.Get("/favicon.ico", handlers.ServeString("favicon.png", now, favicon).ServeHTTP)
	router.Get("/robots.txt", handlers.ServeString("robots.txt", now, robots).ServeHTTP)

	router.Get("/", playlistHandler(bus))

	handler := handlers.AccessLog(router, nil)
	handler = &handlers.SecurityHeaders{
		Handler: router,
	}
	handler = handlers.SetHeader(handler, "Server", "spork (audacious control panel)")

	fmt.Printf("Listening on :%d\n", *port)

	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(*port),
		Handler: handler,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// termination handler
	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	<-term

	// gracefull shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("error shutting down: %v", err)
	}
}
