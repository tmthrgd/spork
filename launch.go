package main

import (
	"log"
	"net/http"
	"os/exec"
	"sync"
	"syscall"

	"tmthrgd.dev/go/spork/internal/dbus"
)

func launchHandler() http.HandlerFunc {
	redirect := func(w http.ResponseWriter, r *http.Request) error {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return nil
	}

	waitStarted, isRunning, err := dbus.WaitForAudacious()
	if err != nil {
		log.Printf("dbus.WaitForAudacious failed (launch endpoint impacted): %v", err)
	}

	var mu sync.Mutex

	return errorHandler(func(w http.ResponseWriter, r *http.Request) error {
		mu.Lock()
		defer mu.Unlock()

		if isRunning() {
			// Audacious is already running.
			return redirect(w, r)
		}

		cmd := exec.Command("audacious")
		cmd.SysProcAttr = &syscall.SysProcAttr{
			// Setpgid causes audacious to have a different process
			// group and not be terminated when the parent process
			// is killed with ^C.
			Setpgid: true,
		}

		if err := cmd.Start(); err != nil {
			return err
		}

		// Wait until Audacious has started and it's D-Bus interface is
		// active.
		waitStarted()

		return redirect(w, r)
	})
}
