package main

import (
	"net/http"

	"github.com/tmthrgd/spork/internal/dbus"
)

var controlTmpl = newTemplate(`<!doctype html>
<meta charset=utf-8>
<title>Play Controls</title>
<link rel=stylesheet href=/assets/style.css>
<body class=page-controls>
<h1>Play Controls</h1>
<div class=controls>
<a href=/controls/prev>⏮</a> <a href=/controls/stop>⏹</a> <a href=/controls/playpause>⏯</a> <a href=/controls/next>⏭</a><br>
<input type=range min=0 max=100 value="{{.Volume}}" title="{{.Volume}}%" class=volume>
<p class=song>{{if .Title}}{{.Title}} ({{FormatLength .Length}}){{end}}</p>
<span class=error></span>
</div>
<a href=/playlist>Active Playlist</a> – <a href=/playlist#current>Current Song</a>
<script defer src=/assets/fetch-helpers.js></script>
<script defer src=/assets/controls.js></script>`)

func controlsHandler() http.HandlerFunc {
	return errorHandler(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		volume, err := dbus.GetVolume(ctx)
		if err != nil {
			return err
		}

		pos, err := dbus.GetPlaylistPosition(ctx)
		if err != nil {
			return err
		}

		title, err := dbus.GetSongTitle(ctx, pos)
		if err != nil {
			return err
		}

		length, err := dbus.GetSongLength(ctx, pos)
		if err != nil {
			return err
		}

		return templateExecute(w, controlTmpl, &struct {
			Volume int32
			Title  string
			Length int32
		}{volume, title, length})
	})
}
