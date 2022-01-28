package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/tmthrgd/httputils"
	"go.tmthrgd.dev/spork/dbus"
	"go.tmthrgd.dev/spork/web"
)

func actionHandlerResponse(w http.ResponseWriter, r *http.Request, ok, redirect string) error {
	if httputils.Negotiate(r.Header, "Accept", "text/html", "text/plain") == "text/plain" {
		io.WriteString(w, ok)
	} else {
		http.Redirect(w, r, redirect, http.StatusTemporaryRedirect)
	}

	return nil
}

func jumpHandler() http.HandlerFunc {
	return web.ErrorHandler(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		pos, err := strconv.ParseUint(chi.URLParam(r, "pos"), 10, 32)
		if err != nil {
			return err
		}

		if err := dbus.PlaylistJump(ctx, uint32(pos)); err != nil {
			return err
		}

		if err := dbus.Play(ctx); err != nil {
			return err
		}

		return actionHandlerResponse(w, r, "ok", "/playlist#current")
	})
}

func setVolumeHandler() http.HandlerFunc {
	return web.ErrorHandler(func(w http.ResponseWriter, r *http.Request) error {
		vol, err := strconv.ParseUint(chi.URLParam(r, "vol"), 10, 7)
		if err != nil {
			return err
		} else if vol > 100 {
			return errors.New("spork: volume out of range")
		}

		if err := dbus.SetVolume(r.Context(), int32(vol)); err != nil {
			return err
		}

		return actionHandlerResponse(w, r, "ok", "/")
	})
}

func controlHandler(fn func(context.Context) error, statusFn func(context.Context) (bool, error)) http.HandlerFunc {
	return web.ErrorHandler(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		if err := fn(ctx); err != nil {
			return err
		}

		ok := "ok"
		if statusFn != nil {
			on, err := statusFn(ctx)
			if err != nil {
				return err
			}

			ok = strconv.FormatBool(on)
		}

		return actionHandlerResponse(w, r, ok, "/")
	})
}

func playPauseHandler() http.HandlerFunc { return controlHandler(dbus.PlayPause, nil) }
func stopHandler() http.HandlerFunc      { return controlHandler(dbus.Stop, nil) }
func prevHandler() http.HandlerFunc      { return controlHandler(dbus.Reverse, nil) }
func nextHandler() http.HandlerFunc      { return controlHandler(dbus.Advance, nil) }
func repeatHandler() http.HandlerFunc    { return controlHandler(dbus.ToggleRepeat, dbus.GetRepeat) }
func shuffleHandler() http.HandlerFunc   { return controlHandler(dbus.ToggleShuffle, dbus.GetShuffle) }

func downloadHandler() http.HandlerFunc {
	var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

	return web.ErrorHandler(func(w http.ResponseWriter, r *http.Request) error {
		pos, err := strconv.ParseUint(chi.URLParam(r, "pos"), 10, 32)
		if err != nil {
			return err
		}

		filename, err := dbus.GetSongFilename(r.Context(), uint32(pos))
		if err != nil {
			return err
		}
		if filename == "" {
			return os.ErrNotExist
		}

		u, err := url.Parse(filename)
		if err != nil {
			return err
		}
		if u.Scheme != "file" {
			return errors.New("spork: song has unsupported filename")
		}

		base := filepath.Base(u.Path)
		w.Header().Set("Content-Disposition",
			fmt.Sprintf(`inline; filename="%s"`, quoteEscaper.Replace(base)))

		http.ServeFile(w, r, u.Path)
		return nil
	})
}
