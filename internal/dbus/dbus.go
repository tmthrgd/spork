package dbus

import "github.com/godbus/dbus"

// Error represents a D-Bus message of type Error.
type Error = dbus.Error

var audObj dbus.BusObject

// BusConnect starts the dbus connection, it should only be called once.
func BusConnect() error {
	if audObj != nil {
		panic("spork/internal/dbus: BusConnect called multiple times")
	}

	conn, err := dbus.SessionBus()
	if err != nil {
		return err
	}

	audObj = conn.Object("org.atheme.audacious", "/org/atheme/audacious")
	return nil
}

// IsUnknownServiceError returns true if the error represents an
// org.freedesktop.DBus.Error.ServiceUnknown error. An error of this type will
// be returned if Audacious is closed.
func IsUnknownServiceError(err error) bool {
	derr, _ := err.(dbus.Error)
	return derr.Name == "org.freedesktop.DBus.Error.ServiceUnknown"
}
