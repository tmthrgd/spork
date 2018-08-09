package main

import (
	"net/http"

	"github.com/tmthrgd/spork/internal/dbus"
)

var controlTmpl = newTemplate(`<!doctype html>
<meta charset=utf-8>
<title>Audacious Control Panel – Spork</title>
<link rel=stylesheet href=/assets/style.css>
<body class=page-controls>
<h1>Audacious Control Panel</h1>
<div class=controls>
<a href=/controls/prev title=Previous>⏮</a>
<a href=/controls/stop title=Stop>⏹</a>
<a href=/controls/playpause title=Play/Pause>⏯</a>
<a href=/controls/next title=Next>⏭</a>
&nbsp;
<a href=/controls/repeat title=Repeat data-toggle {{- if .Repeat}} class="active"{{end}}>🔁</a>
<a href=/controls/shuffle title=Shuffle data-toggle {{- if .Shuffle}} class="active"{{end}}>🔀</a>
<br>
<input type=range min=0 max=100 value="{{.Volume}}" title="{{.Volume}}%" class=volume>
<p class=song>{{if .Title}}{{.Title}} ({{FormatLength .Length}}){{end}}</p>
<p class=error></p>
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

		shuffle, err := dbus.GetShuffle(ctx)
		if err != nil {
			return err
		}

		repeat, err := dbus.GetRepeat(ctx)
		if err != nil {
			return err
		}

		return templateExecute(w, controlTmpl, &struct {
			Volume          int32
			Title           string
			Length          int32
			Shuffle, Repeat bool
		}{volume, title, length, shuffle, repeat})
	})
}
