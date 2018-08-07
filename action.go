package main

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/tmthrgd/spork/internal/dbus"
)

func jumpHandler(bus *dbus.Bus) http.HandlerFunc {
	return httpHandlerError(func(w http.ResponseWriter, r *http.Request) error {
		pos, err := strconv.ParseUint(chi.URLParam(r, "pos"), 10, 32)
		if err != nil {
			return err
		}

		if err := bus.PlaylistJump(r.Context(), uint32(pos)); err != nil {
			return err
		}

		http.Redirect(w, r, "/#current", http.StatusTemporaryRedirect)
		//io.WriteString(w, `<script>history.back(-1)</script>`)
		return nil
	})
}
