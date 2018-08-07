package dbus

import "github.com/godbus/dbus"

type Bus struct {
	conn *dbus.Conn

	audaciousObject dbus.BusObject
}

func SessionBus() (*Bus, error) {
	conn, err := dbus.SessionBus()
	if err != nil {
		return nil, err
	}

	return &Bus{
		conn,

		conn.Object("org.mpris.MediaPlayer2.audacious",
			"/org/atheme/audacious"),
	}, nil
}

func (b *Bus) Close() error {
	return b.conn.Close()
}
