package main

import (
	"net/http"
	"os/exec"
	"syscall"
)

func launchHandler() http.HandlerFunc {
	return errorHandler(func(w http.ResponseWriter, r *http.Request) error {
		cmd := exec.Command("audacious")
		cmd.SysProcAttr = &syscall.SysProcAttr{
			// Setpgid causes audacious to have a different process group and
			// not be terminated when the parent process is killed with ^C.
			Setpgid: true,
		}

		if err := cmd.Start(); err != nil {
			return err
		}

		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return nil
	})
}
