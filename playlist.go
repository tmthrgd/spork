package main

import (
	"fmt"
	"net/http"

	"github.com/tmthrgd/spork/internal/dbus"
	"golang.org/x/sync/errgroup"
)

var playlistTmpl = newTemplate(`<!doctype html>
<meta charset=utf-8>
<title>{{.Name}}</title>
<link rel=stylesheet href=/assets/style.css>
<body class=page-playlist>
<h1>{{.Name}}</h1>
<p class=error></p>
<ol class=playlist>
{{- range $idx, $entry := .Entries}}
<li{{if eq $idx $.Active}} class="active" id="current"{{end}}>
<details>
<summary><a href="/jump/{{$idx}}">{{.Title}}</a> ({{.Length}})</summary>
{{if .Name -}}  <p>Title:  {{.Name}}</p>{{end}}
{{if .Artist -}}<p>Artist: {{.Artist}}</p>{{end}}
{{if .Album -}} <p>Album:  {{.Album}}</p>{{end}}
{{if .Year -}}  <p>Year:   {{.Year}}</p>{{end}}
<p>Length: {{.Length}}</p>
</details>
</li>
{{- end}}
</ol>
<script defer src=/assets/fetch-helpers.js></script>
<script defer src=/assets/playlist.js></script>`)

type playlistData struct {
	Name    string
	Entries []playlistEntry
	Active  uint32
}

type playlistEntry struct {
	Title, Name, Artist, Album string
	length, Year               int32
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

		playlist := make([]playlistEntry, entries)

		g, gctx := errgroup.WithContext(ctx)
		sem := make(chan struct{}, 8)

		for entry := uint32(0); entry < uint32(entries); entry++ {
			select {
			case sem <- struct{}{}:
			case <-gctx.Done():
				return g.Wait()
			}

			entry := entry
			g.Go(func() error {
				defer func() { <-sem }()

				title, err := dbus.GetSongTitle(gctx, entry)
				if err != nil {
					return err
				}

				length, err := dbus.GetSongLength(gctx, entry)
				if err != nil {
					return err
				}

				name, err := dbus.GetSongName(gctx, entry)
				if err != nil {
					return err
				}

				artist, err := dbus.GetSongArtist(gctx, entry)
				if err != nil {
					return err
				}

				album, err := dbus.GetSongAlbum(gctx, entry)
				if err != nil {
					return err
				}

				year, err := dbus.GetSongYear(gctx, entry)
				if err != nil {
					return err
				}

				playlist[entry] = playlistEntry{
					title, name, artist, album,
					length, year,
				}
				return nil
			})
		}

		if err := g.Wait(); err != nil {
			return err
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
