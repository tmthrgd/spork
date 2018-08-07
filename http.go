package main

import (
	"bytes"
	"html/template"
	"net/http"
)

var error500 = template.Must(template.New("error500").Parse(`<!doctype html>
<meta charset=utf-8>
<title>500 Internal Server Error</title>
<style>body{margin:40px auto;max-width:650px;line-height:1.6;font-size:18px;color:#444;padding:0 10px}h1,h2,h3{line-height:1.2}</style>
<h1>500 Internal Server Error</h1>
<p>{{.}}</p>`))

func httpHandlerError(handler func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err == nil {
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)

		error500.Execute(w, err.Error())
	}
}

func templateExecute(w http.ResponseWriter, tmpl *template.Template, data interface{}) error {
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return err
	}

	buf.WriteTo(w)
	return nil
}
