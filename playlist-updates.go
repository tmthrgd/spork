package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"sync/atomic"
	"time"

	"go.tmthrgd.dev/spork/dbus"
	"go.tmthrgd.dev/spork/web"
)

type updateJSON struct {
	Position uint32 `json:"pos"`
	Title    string `json:"title,omitempty"`
	Length   int32  `json:"length,omitempty"`
}

func (u *updateJSON) fill(ctx context.Context) error {
	var (
		data updateJSON
		err  error
	)

	if data.Position, err = dbus.GetPlaylistPosition(ctx); err != nil {
		return err
	}

	if data.Title, err = dbus.GetSongTitle(ctx, data.Position); err != nil {
		return err
	}

	if data.Length, err = dbus.GetSongLength(ctx, data.Position); err != nil {
		return err
	}

	*u = data
	return nil
}

func (u *updateJSON) marshal() []byte {
	data, err := json.Marshal(u)
	if err != nil {
		panic("spork: internal error: unable to marshal *updateJSON: " + err.Error())
	}

	return data
}

func playlistUpdateHandler(ctx context.Context) http.HandlerFunc {
	type updateData struct {
		pos  uint32
		data []byte

		nextCh chan struct{}
		next   *updateData
	}
	var update atomic.Value // *updateData

	go func() {
		var data updateJSON

		if err := data.fill(ctx); err != nil && !dbus.IsUnknownServiceError(err) {
			log.Printf("spork: playlist update poller error: %v", err)
		}

		old := &updateData{
			pos:  data.Position,
			data: data.marshal(),

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

			if err := data.fill(ctx); err != nil {
				if !dbus.IsUnknownServiceError(err) {
					log.Printf("spork: playlist update poller error: %v", err)
				}

				continue
			}

			nextData := data.marshal()
			if bytes.Equal(old.data, nextData) {
				continue
			}

			next := &updateData{
				pos:  data.Position,
				data: nextData,

				nextCh: make(chan struct{}),
			}
			update.Store(next)

			old.next = next
			close(old.nextCh)

			old = next
		}
	}()

	return web.ErrorHandler(func(w http.ResponseWriter, r *http.Request) error {
		f, ok := w.(http.Flusher)
		if !ok {
			return errors.New("spork: http.ResponseWriter does not implement http.Flusher")
		}

		hdr := w.Header()
		hdr.Set("Content-Type", "text/event-stream")
		hdr.Set("Connection", "keep-alive")

		w.WriteHeader(http.StatusOK)
		f.Flush()

		update := update.Load().(*updateData)
		for {
			select {
			case <-update.nextCh:
				update = update.next
			case <-r.Context().Done():
				return nil
			case <-shutdown:
				return nil
			}

			io.WriteString(w, "data: ")
			w.Write(update.data)
			io.WriteString(w, "\n\n")
			f.Flush()
		}
	})
}
