package main

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/tmthrgd/spork/internal/dbus"
)

func jumpHandler() http.HandlerFunc {
	return httpHandlerError(func(w http.ResponseWriter, r *http.Request) error {
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

		http.Redirect(w, r, "/#current", http.StatusTemporaryRedirect)
		return nil
	})
}
