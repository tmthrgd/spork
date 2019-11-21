package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	handlers "github.com/tmthrgd/httphandlers"
	"go.tmthrgd.dev/spork/web"
)

func handler(ctx context.Context) http.Handler {
	router := chi.NewRouter()
	router.Use(
		handlers.AccessLogWrap(nil),
		handlers.SecurityHeadersWrap(nil),
		handlers.SetHeaderWrap("Server", "spork (audacious control panel)"),
	)
	router.NotFound(web.NotFoundHandler())
	router.MethodNotAllowed(web.MethodNotAllowedHandler())

	// assets routes
	web.MountAssets(router)

	// pprof handler
	router.Mount("/debug", middleware.Profiler())

	// HTML page routes
	router.Group(func(r chi.Router) {
		r.Use(middleware.NoCache)

		r.Get("/", controlsHandler())
		r.Get("/playlist", playlistHandler())
	})

	// API routes
	router.Group(func(r chi.Router) {
		r.Use(middleware.NoCache)

		r.Get("/jump/{pos}", jumpHandler())
		r.Get("/controls/playpause", playPauseHandler())
		r.Get("/controls/stop", stopHandler())
		r.Get("/controls/prev", prevHandler())
		r.Get("/controls/next", nextHandler())
		r.Get("/controls/repeat", repeatHandler())
		r.Get("/controls/shuffle", shuffleHandler())
		r.Get("/controls/volume/{vol}", setVolumeHandler())
		r.Get("/playlist/updates", playlistUpdateHandler(ctx))
		r.Get("/launch", launchHandler())
	})

	return router
}
