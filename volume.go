package main

import (
	"html/template"
	"net/http"

	"github.com/tmthrgd/spork/internal/dbus"
)

var volumeTmpl = template.Must(template.New("volume").Parse(`<!doctype html>
<meta charset=utf-8>
<title>Volume Control</title>
<link rel=stylesheet href=/assets/style.css>
<h1>Volume Control</h1>
Current volume: <span class=volume>{{.}}</span>%<br>
<input type=range min=0 max=100 value="{{.}}" class=slider>
<script defer src=/assets/volume.js></script>`))

func volumeHandler() http.HandlerFunc {
	return errorHandler(func(w http.ResponseWriter, r *http.Request) error {
		volume, err := dbus.GetVolume(r.Context())
		if err != nil {
			return err
		}

		return templateExecute(w, volumeTmpl, volume)
	})
}
