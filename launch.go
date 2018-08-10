package main

import (
	"net/http"
	"os/exec"
	"syscall"

	"github.com/tmthrgd/spork/internal/dbus"
)

func launchHandler() http.HandlerFunc {
	redirect := func(w http.ResponseWriter, r *http.Request) error {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return nil
	}

	return errorHandler(func(w http.ResponseWriter, r *http.Request) error {
		if _, err := dbus.GetVolume(r.Context()); err == nil {
			// Audacious is already running.
			return redirect(w, r)
		}

		cmd := exec.Command("audacious")
		cmd.SysProcAttr = &syscall.SysProcAttr{
			// Setpgid causes audacious to have a different process group and
			// not be terminated when the parent process is killed with ^C.
			Setpgid: true,
		}

		if err := cmd.Start(); err != nil {
			return err
		}

		return redirect(w, r)
	})
}
