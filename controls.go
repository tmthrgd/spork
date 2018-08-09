package main

import (
	"html/template"
	"net/http"

	"github.com/tmthrgd/spork/internal/dbus"
)

var controlTmpl = template.Must(template.New("control").Parse(`<!doctype html>
<meta charset=utf-8>
<title>Play Controls</title>
<link rel=stylesheet href=/assets/style.css>
<body class=page-controls>
<h1>Play Controls</h1>
<div class=controls>
<a href=/controls/prev>⏮</a> <a href=/controls/stop>⏹</a> <a href=/controls/playpause>⏯</a> <a href=/controls/next>⏭</a><br>
<input type=range min=0 max=100 value="{{.Volume}}" title="{{.Volume}}%" class=volume>
</div>
<a href=/playlist>Active Playlist</a> – <a href=/playlist#current>Current Song</a>
<script defer src=/assets/controls.js></script>`))

func controlsHandler() http.HandlerFunc {
	return errorHandler(func(w http.ResponseWriter, r *http.Request) error {
		volume, err := dbus.GetVolume(r.Context())
		if err != nil {
			return err
		}

		return templateExecute(w, controlTmpl, struct{ Volume int32 }{volume})
	})
}
