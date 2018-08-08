package main

import (
	"net/http"
	"time"

	"github.com/tmthrgd/httphandlers"
)

const controlPage = `<!doctype html>
<meta charset=utf-8>
<title>Play Controls</title>
<style>body{margin:40px auto;max-width:650px;line-height:1.6;font-size:18px;color:#444;padding:0 10px}h1,h2,h3{line-height:1.2}a{text-decoration:none}</style>
<h1>Play Controls</h1>
<a href=/controls/prev>⏮</a> <a href=/controls/stop>⏹</a> <a href=/controls/pause>⏸</a> <a href=/controls/play>▶️</a> <a href=/controls/next>⏭</a>`

func controlsHandler() http.HandlerFunc {
	return handlers.ServeString("controls.html", time.Now(), controlPage).ServeHTTP
}
