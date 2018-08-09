package main

//go:generate go run -tags=dev internal/assets/generate.go

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

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/tmthrgd/httphandlers"
	"github.com/tmthrgd/spork/internal/assets"
	"github.com/tmthrgd/spork/internal/dbus"
)

const (
	favicon = "\x89PNG\r\n\x1a\n\x00\x00\x00\rIHDR\x00\x00\x00\x10\x00\x00\x00\x10\b\x06\x00\x00\x00\x1f\xf3\xffa\x00\x00\x00\x16IDATx\xdacd\xa0\x100\x8e\x1a0j\xc0\xa8\x01\xc3\xc5\x00\x00&\x98\x00\x11\x9b\x92AZ\x00\x00\x00\x00IEND\xaeB`\x82"
	robots  = "User-agent: *\nDisallow: /"
)

var shutdown = make(chan struct{})

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

	router := chi.NewRouter()
	router.Use(
		middleware.GetHead,
		handlers.AccessLogWrap(nil),
		handlers.SecurityHeadersWrap(nil),
		handlers.SetHeaderWrap("Server", "spork (audacious control panel)"),
	)
	router.NotFound(notFoundHandler())

	// pprof handler
	router.Mount("/debug", middleware.Profiler())

	// Asset routes
	router.Group(func(r chi.Router) {
		now := time.Now()
		r.Get("/favicon.ico", handlers.ServeString("favicon.png", now, favicon).ServeHTTP)
		r.Get("/robots.txt", handlers.ServeString("robots.txt", now, robots).ServeHTTP)

		r.Mount("/assets", http.StripPrefix("/assets", http.FileServer(&noDirFileSystem{assets.FileSystem})))
	})

	// HTML page routes
	router.Group(func(r chi.Router) {
		r.Use(middleware.NoCache)

		r.Get("/", controlsHandler())
		r.Get("/playlist", playlistHandler())
	})

	// API routes
	router.Group(func(r chi.Router) {
		r.Use(
			undoGetHead,
			middleware.NoCache,
		)

		r.Get("/jump/{pos}", jumpHandler())
		r.Get("/controls/playpause", playPauseHandler())
		r.Get("/controls/stop", stopHandler())
		r.Get("/controls/prev", prevHandler())
		r.Get("/controls/next", nextHandler())
		r.Get("/controls/repeat", repeatHandler())
		r.Get("/controls/shuffle", shuffleHandler())
		r.Get("/controls/volume/{vol}", setVolumeHandler())
		r.Get("/playlist/updates", playlistUpdateHandler(ctx))
	})

	fmt.Printf("Listening on %s\n", *addr)

	srv := &http.Server{
		Addr:    *addr,
		Handler: router,
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
	sctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	go func() {
		<-term
		signal.Stop(term)

		cancel()
	}()

	close(shutdown)

	if err := srv.Shutdown(sctx); err != nil {
		log.Printf("error shutting down: %v", err)
	}
}
