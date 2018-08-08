package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/tmthrgd/spork/internal/dbus"
)

var playlistTmpl = template.Must(template.New("playlist").Parse(`<!doctype html>
<meta charset=utf-8>
<title>{{.Name}}</title>
<link rel=stylesheet href=/assets/style.css>
<body class=page-playlist>
<h1>{{.Name}}</h1>
<ol>
{{- range $idx, $entry := .Entries}}
<li{{if eq $idx $.Active}} class="active" id="current"{{end}}><a href="/jump/{{$idx}}">{{.Title}}</a> – {{.Length}}</li>
{{- end}}
</ol>`))

type playlistData struct {
	Name    string
	Entries []playlistEntry
	Active  uint32
}

type playlistEntry struct {
	Title  string
	length int32
}

func (e playlistEntry) Length() string {
	return fmt.Sprintf("%d:%02d", e.length/60, e.length%60)
}

func playlistHandler() http.HandlerFunc {
	return errorHandler(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		entries, err := dbus.GetPlaylistLength(ctx)
		if err != nil {
			return err
		}

		playlist := make([]playlistEntry, 0, entries)

		for entry := uint32(0); entry < uint32(entries); entry++ {
			title, err := dbus.GetSongTitle(ctx, entry)
			if err != nil {
				return err
			}

			length, err := dbus.GetSongLength(ctx, entry)
			if err != nil {
				return err
			}

			playlist = append(playlist, playlistEntry{title, length})
		}

		name, err := dbus.GetPlaylistName(ctx)
		if err != nil {
			return err
		}

		active, err := dbus.GetPlaylistPosition(ctx)
		if err != nil {
			return err
		}

		return templateExecute(w, playlistTmpl, &playlistData{
			Name:    name,
			Entries: playlist,
			Active:  active,
		})
	})
}
