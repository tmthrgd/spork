// Package dbus implements a D-Bus interface to Audacious.
package dbus

import (
	"errors"

	"github.com/godbus/dbus"
)

// Error represents a D-Bus message of type Error.
type Error = dbus.Error

var (
	conn   *dbus.Conn
	audObj dbus.BusObject
)

// BusConnect starts the dbus connection, it should only be called once.
func BusConnect() error {
	if conn != nil {
		panic("spork/internal/dbus: BusConnect called multiple times")
	}

	var err error
	if conn, err = dbus.SessionBus(); err != nil {
		return err
	}

	audObj = conn.Object("org.atheme.audacious", "/org/atheme/audacious")
	return nil
}

// IsUnknownServiceError returns true if the error represents an
// org.freedesktop.DBus.Error.ServiceUnknown error. An error of this type will
// be returned if Audacious is closed.
func IsUnknownServiceError(err error) bool {
	var derr dbus.Error
	return errors.As(err, &derr) &&
		derr.Name == "org.freedesktop.DBus.Error.ServiceUnknown"
}
