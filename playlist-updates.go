package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/tmthrgd/spork/internal/dbus"
)

func playlistUpdateHandler(ctx context.Context) http.HandlerFunc {
	type positionUpdate struct {
		pos uint32

		nextCh chan struct{}
		next   *positionUpdate
	}
	var update atomic.Value // *positionUpdate

	go func() {
		pos, err := dbus.GetPlaylistPosition(ctx)
		if err != nil && !dbus.IsUnknownServiceError(err) {
			log.Printf("spork: playlist update poller error: %v", err)
		}

		old := &positionUpdate{
			pos: pos,

			nextCh: make(chan struct{}),
		}
		update.Store(old)

		t := time.NewTicker(time.Second)
		defer t.Stop()

		for {
			select {
			case <-t.C:
			case <-ctx.Done():
				return
			}

			pos, err := dbus.GetPlaylistPosition(ctx)
			if err != nil {
				if !dbus.IsUnknownServiceError(err) {
					log.Printf("spork: playlist update poller error: %v", err)
				}

				continue
			} else if pos == old.pos {
				continue
			}

			next := &positionUpdate{
				pos: pos,

				nextCh: make(chan struct{}),
			}
			update.Store(next)

			old.next = next
			close(old.nextCh)

			old = next
		}
	}()

	return errorHandler(func(w http.ResponseWriter, r *http.Request) error {
		f, ok := w.(http.Flusher)
		if !ok {
			return errors.New("spork: http.ResponseWriter does not implement http.Flusher")
		}

		hdr := w.Header()
		hdr.Set("Content-Type", "text/event-stream")
		hdr.Set("Cache-Control", "no-cache")
		hdr.Set("Connection", "keep-alive")

		w.WriteHeader(http.StatusOK)
		f.Flush()

		update := update.Load().(*positionUpdate)

		for {
			select {
			case <-update.nextCh:
				update = update.next
			case <-r.Context().Done():
				return nil
			case <-shutdown:
				return nil
			}

			fmt.Fprintf(w, "data: {\"pos\":%d}\n\n", update.pos)
			f.Flush()
		}
	})
}
