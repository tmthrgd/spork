package dbus

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/godbus/dbus"
)

// WaitForAudacious monitors the running state of Audacious. It returns two
// funcs, the first of which waits until Audacious has started, and the second
// of which can be used to query it's running state.
//
// Neither func will ever be nil, even in the case of err being non-nil.
func WaitForAudacious() (waitStarted func(), isRunning func() bool, err error) {
	waitStarted = waitStartedFallback
	isRunning = isRunningFallback

	var running int32
	cond := sync.NewCond(new(sync.Mutex))

	ch := make(chan *dbus.Signal, 1)
	conn.Signal(ch)

	go func() {
		var name, oldOwner, newOwner string
		for sig := range ch {
			if sig.Name != "org.freedesktop.DBus.NameOwnerChanged" ||
				dbus.Store(sig.Body, &name, &oldOwner, &newOwner) != nil ||
				name != "org.atheme.audacious" {
				continue
			}

			if newOwner != "" {
				atomic.StoreInt32(&running, 1)
			} else {
				atomic.StoreInt32(&running, 0)
			}

			cond.Broadcast()
		}
	}()

	// See the D-Bus specification for the format of this rule:
	// https://dbus.freedesktop.org/doc/dbus-specification.html#message-bus-routing-match-rules
	const matchRule = "type='signal'," +
		"sender='org.freedesktop.DBus'," +
		"interface='org.freedesktop.DBus'," +
		"member='NameOwnerChanged'," +
		"path='/org/freedesktop/DBus'," +
		"arg0='org.atheme.audacious'"

	if err = conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, matchRule).Err; err != nil {
		conn.RemoveSignal(ch)
		close(ch)
		return
	}

	// Prime the running state in case Audacious was already started.
	if isRunningFallback() {
		atomic.StoreInt32(&running, 1)
	}

	isRunning = func() bool {
		return atomic.LoadInt32(&running) != 0
	}

	waitStarted = func() {
		cond.L.Lock()
		for !isRunning() {
			cond.Wait()
		}
		cond.L.Unlock()
	}

	return
}

func waitStartedFallback() { /* NOOP */ }

func isRunningFallback() bool {
	// The fallback should never be allowed to block for too long.
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	_, err := GetVolume(ctx)
	return err == nil
}
