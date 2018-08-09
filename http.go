package main

import (
	"bytes"
	"html/template"
	"io"
	"net/http"
	"reflect"

	"github.com/go-chi/chi"
	"github.com/godbus/dbus"
	"github.com/tmthrgd/httphandlers"
)

const error404 = `<!doctype html>
<meta charset=utf-8>
<title>404 Not Found</title>
<style>body{margin:40px auto;max-width:650px;line-height:1.6;font-size:18px;color:#444;padding:0 10px}h1,h2,h3{line-height:1.2}</style>
<h1>404 Not Found</h1>
<p>The requested file was not found.</p>`

var error500 = template.Must(template.New("error500").Parse(`<!doctype html>
<meta charset=utf-8>
<title>500 Internal Server Error</title>
<style>body{margin:40px auto;max-width:650px;line-height:1.6;font-size:18px;color:#444;padding:0 10px}h1,h2,h3{line-height:1.2}</style>
<h1>500 Internal Server Error</h1>
<p>{{.Type}}: {{- if .Name}} {{.Name}}: {{- end}} {{.Message}}</p>`))

const error502NoAudacious = `<!doctype html>
<meta charset=utf-8>
<title>502 Bad Gateway</title>
<style>body{margin:40px auto;max-width:650px;line-height:1.6;font-size:18px;color:#444;padding:0 10px}h1,h2,h3{line-height:1.2}</style>
<h1>502 Bad Gateway</h1>
<p>The Audacious Media Player is not currently running. Please start Audacious and try again.</p>`

// notFoundHandler returns a handler that serves a 404 error page.
func notFoundHandler() http.HandlerFunc {
	return handlers.ServeError(http.StatusNotFound, []byte(error404), "text/html; charset=utf-8").ServeHTTP
}

// errorHandler converts a handler with an error return to a http.HandlerFunc,
// sending a 500 Internal Server Error, or a 502 Bad Gateway where appropriate,
// to the client when an error is returned.
func errorHandler(handler func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err == nil {
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		derr, ok := err.(dbus.Error)
		if ok && derr.Name == "org.freedesktop.DBus.Error.ServiceUnknown" {
			w.WriteHeader(http.StatusBadGateway)

			io.WriteString(w, error502NoAudacious)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)

		error500.Execute(w, &struct {
			Type, Name, Message string
		}{
			reflect.ValueOf(err).Type().String(),
			derr.Name,
			err.Error(),
		})
	}
}

// templateExecute calls Execute on the given template, only writing out the result if
// execution was successful.
func templateExecute(w http.ResponseWriter, tmpl *template.Template, data interface{}) error {
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return err
	}

	if n, err := buf.WriteTo(w); n == 0 {
		// Only forward the error if nothing was written.
		return err
	}

	return nil
}

// undoGetHead is a middleware that undoes the effect of chi/middleware.GetHead.
func undoGetHead(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		if rctx != nil && rctx.RouteMethod != r.Method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		next.ServeHTTP(w, r)
	})
}
