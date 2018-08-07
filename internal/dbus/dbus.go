package dbus

import "github.com/godbus/dbus"

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
