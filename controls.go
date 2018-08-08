package main

import (
	"net/http"
	"time"

	"github.com/tmthrgd/httphandlers"
)

const controlPage = `<!doctype html>
<meta charset=utf-8>
<title>Play Controls</title>
<link rel=stylesheet href=/assets/style.css>
<body class=page-controls>
<h1>Play Controls</h1>
<a href=/controls/prev>⏮</a> <a href=/controls/stop>⏹</a> <a href=/controls/pause>⏸</a> <a href=/controls/play>▶️</a> <a href=/controls/next>⏭</a>`

func controlsHandler() http.HandlerFunc {
	return handlers.ServeString("controls.html", time.Now(), controlPage).ServeHTTP
}
