package main

import (
	"html/template"
	"net/http"

	"github.com/tmthrgd/spork/internal/dbus"
)

var volumeTmpl = template.Must(template.New("volume").Parse(`<!doctype html>
<meta charset=utf-8>
<title>Volume Control</title>
<style>body{margin:40px auto;max-width:650px;line-height:1.6;font-size:18px;color:#444;padding:0 10px}h1,h2,h3{line-height:1.2}input[type=range]::-moz-range-track{background-color:lightgrey}</style>
<h1>Volume Control</h1>
Current volume: <span class=volume>{{.}}</span>%<br>
<input type=range min=0 max=100 value="{{.}}" class=slider>
<script>document.querySelector('.slider').addEventListener('input',e=>{document.querySelector('.volume').textContent=e.target.value;fetch('/volume/'+encodeURIComponent(e.target.value)).catch(console.error)},!1)</script>`))

func volumeHandler() http.HandlerFunc {
	return errorHandler(func(w http.ResponseWriter, r *http.Request) error {
		volume, err := dbus.GetVolume(r.Context())
		if err != nil {
			return err
		}

		return templateExecute(w, volumeTmpl, volume)
	})
}
