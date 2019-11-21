package main

import (
	"html/template"
	"net/http"

	"go.tmthrgd.dev/spork/dbus"
	"go.tmthrgd.dev/spork/web"
	"golang.org/x/sync/errgroup"
)

var playlistTmpl = web.NewTemplate("playlist.tmpl", template.FuncMap{
	"FormatLength": formatLength,
})

func playlistHandler() http.HandlerFunc {
	return web.ErrorHandler(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		entries, err := dbus.GetPlaylistLength(ctx)
		if err != nil {
			return err
		}

		type playlistEntry struct {
			Title, Name, Artist, Album string
			Length, Year               int32
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

		return web.WriteTemplateResponse(w, playlistTmpl, &struct {
			Name    string
			Entries []playlistEntry
			Active  uint32
		}{name, playlist, active})
	})
}
