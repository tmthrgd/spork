package main

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/tmthrgd/httputils"
	"github.com/tmthrgd/spork/internal/dbus"
)

func actionHandlerResponse(w http.ResponseWriter, r *http.Request, redirect string) error {
	if httputils.Negotiate(r.Header, "Accept", "text/html", "text/plain") == "text/plain" {
		io.WriteString(w, "ok")
	} else {
		http.Redirect(w, r, redirect, http.StatusTemporaryRedirect)
	}

	return nil
}

func jumpHandler() http.HandlerFunc {
	return errorHandler(func(w http.ResponseWriter, r *http.Request) error {
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

		return actionHandlerResponse(w, r, "/playlist#current")
	})
}

func setVolumeHandler() http.HandlerFunc {
	return errorHandler(func(w http.ResponseWriter, r *http.Request) error {
		vol, err := strconv.ParseUint(chi.URLParam(r, "vol"), 10, 7)
		if err != nil {
			return err
		} else if vol > 100 {
			return errors.New("spork: volume out of range")
		}

		if err := dbus.SetVolume(r.Context(), int32(vol)); err != nil {
			return err
		}

		return actionHandlerResponse(w, r, "/")
	})
}

func controlHandler(fn func(context.Context) error) http.HandlerFunc {
	return errorHandler(func(w http.ResponseWriter, r *http.Request) error {
		if err := fn(r.Context()); err != nil {
			return err
		}

		return actionHandlerResponse(w, r, "/")
	})
}

func playHandler() http.HandlerFunc      { return controlHandler(dbus.Play) }
func playPauseHandler() http.HandlerFunc { return controlHandler(dbus.PlayPause) }
func stopHandler() http.HandlerFunc      { return controlHandler(dbus.Stop) }
func prevHandler() http.HandlerFunc      { return controlHandler(dbus.Reverse) }
func nextHandler() http.HandlerFunc      { return controlHandler(dbus.Advance) }
func repeatHandler() http.HandlerFunc    { return controlHandler(dbus.ToggleRepeat) }
func shuffleHandler() http.HandlerFunc   { return controlHandler(dbus.ToggleShuffle) }
