package main

import (
	"fmt"
	"html/template"
	"net/http"

	"go.tmthrgd.dev/spork/dbus"
	"go.tmthrgd.dev/spork/web"
)

var controlTmpl = web.NewTemplate("control.tmpl", template.FuncMap{
	"FormatLength": formatLength,
})

func controlsHandler() http.HandlerFunc {
	return web.ErrorHandler(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		volume, err := dbus.GetVolume(ctx)
		if err != nil {
			return err
		}

		pos, err := dbus.GetPlaylistPosition(ctx)
		if err != nil {
			return err
		}

		title, err := dbus.GetSongTitle(ctx, pos)
		if err != nil {
			return err
		}

		length, err := dbus.GetSongLength(ctx, pos)
		if err != nil {
			return err
		}

		shuffle, err := dbus.GetShuffle(ctx)
		if err != nil {
			return err
		}

		repeat, err := dbus.GetRepeat(ctx)
		if err != nil {
			return err
		}

		return web.WriteTemplateResponse(w, controlTmpl, &struct {
			Volume          int32
			Title           string
			Length          int32
			Shuffle, Repeat bool
		}{volume, title, length, shuffle, repeat})
	})
}

// formatLength formats a given song length in seconds into a display friendly
// minute:seconds format.
func formatLength(length int32) string {
	return fmt.Sprintf("%d:%02d", length/60, length%60)
}
