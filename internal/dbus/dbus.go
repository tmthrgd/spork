package dbus

import "github.com/godbus/dbus"

var (
	conn *dbus.Conn

	audObj dbus.BusObject
)

func BusConnect() error {
	var err error
	if conn, err = dbus.SessionBus(); err != nil {
		return err
	}

	audObj = conn.Object("org.atheme.audacious", "/org/atheme/audacious")
	return nil
}
