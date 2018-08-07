package main

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"net/http"

	"github.com/tmthrgd/spork/internal/dbus"
)

var playlistTmpl = template.Must(template.New("playlist").Parse(`<!doctype html>
<meta charset=utf-8>
<title>{{.Name}}</title>
<style>body{margin:40px auto;max-width:650px;line-height:1.6;font-size:18px;color:#444;padding:0 10px}h1,h2,h3{line-height:1.2}.active{font-weight:bold}li a{color:#444}</style>
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
	Length string
}

func playlistHandler(bus *dbus.Bus) http.HandlerFunc {
	return httpHandlerError(func(w http.ResponseWriter, r *http.Request) error {
		ctx, cancel := context.WithCancel(r.Context())
		defer cancel()

		entries, err := bus.GetPlaylistLength(ctx)
		if err != nil {
			return err
		}

		playlist := make([]playlistEntry, 0, entries)

		for entry := uint32(0); entry < uint32(entries); entry++ {
			title, err := bus.GetSongTitle(ctx, entry)
			if err != nil {
				return err
			}

			length, err := bus.GetSongLength(ctx, entry)
			if err != nil {
				return err
			}

			playlist = append(playlist, playlistEntry{
				title,
				fmt.Sprintf("%d:%02d", length/60, length%60),
			})
		}

		name, err := bus.GetPlaylistName(ctx)
		if err != nil {
			return err
		}

		active, err := bus.GetPlaylistPosition(ctx)
		if err != nil {
			return err
		}

		var buf bytes.Buffer
		if err := playlistTmpl.Execute(&buf, &playlistData{
			Name:    name,
			Entries: playlist,
			Active:  active,
		}); err != nil {
			return err
		}

		buf.WriteTo(w)
		return nil
	})
}
