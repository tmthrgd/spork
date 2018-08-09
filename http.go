package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"reflect"

	"github.com/go-chi/chi"
	"github.com/tmthrgd/httphandlers"
	"github.com/tmthrgd/spork/internal/dbus"
)

const error404 = `<!doctype html>
<meta charset=utf-8>
<title>404 Not Found</title>
<link rel=stylesheet href=/assets/style.css>
<h1>404 Not Found</h1>
<p>The requested file was not found.</p>`

var error500 = newTemplate(`<!doctype html>
<meta charset=utf-8>
<title>500 Internal Server Error</title>
<link rel=stylesheet href=/assets/style.css>
<h1>500 Internal Server Error</h1>
<p>{{.Type}}: {{- if .Name}} {{.Name}}: {{- end}} {{.Message}}</p>`)

const error502NoAudacious = `<!doctype html>
<meta charset=utf-8>
<title>502 Bad Gateway</title>
<link rel=stylesheet href=/assets/style.css>
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

		if dbus.IsUnknownServiceError(err) {
			w.WriteHeader(http.StatusBadGateway)

			io.WriteString(w, error502NoAudacious)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)

		derr, _ := err.(dbus.Error)
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

// newTemplate parses source and returns a new html/template.Template. It
// panics if source is invalid.
func newTemplate(source string) *template.Template {
	return template.Must(template.New("").Funcs(templateFuncs).Parse(source))
}

var templateFuncs = template.FuncMap{
	"FormatLength": formatLength,
}

// formatLength formats a given song length in seconds into a display friendly
// minute:seconds format.
func formatLength(length int32) string {
	return fmt.Sprintf("%d:%02d", length/60, length%60)
}

type noDirFileSystem struct{ http.FileSystem }

func (fs *noDirFileSystem) Open(name string) (http.File, error) {
	f, err := fs.FileSystem.Open(name)
	if err != nil {
		return nil, err
	}

	if stat, err := f.Stat(); err == nil && stat.IsDir() {
		f.Close()
		return nil, os.ErrNotExist
	}

	return f, nil
}
