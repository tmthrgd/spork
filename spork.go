package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tmthrgd/spork/internal/dbus"
)

func main() {
	bus, err := dbus.SessionBus()
	if err != nil {
		log.Fatal(err)
	}
	defer bus.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	entries, err := bus.GetPlaylistLength(ctx)
	if err != nil {
		log.Fatal(err)
	}

	plural := "s"
	if entries == 1 {
		plural = ""
	}

	fmt.Printf("%d track%s.\n", entries, plural)

	var total uint64

	for entry := uint32(0); entry < uint32(entries); entry++ {
		title, err := bus.GetSongTitle(ctx, entry)
		if err != nil {
			log.Fatal(err)
		}

		length, err := bus.GetSongLength(ctx, entry)
		if err != nil {
			log.Fatal(err)
		}

		total += uint64(length)

		fmt.Printf("%4d | %-60s | %d:%02d\n", entry+1, title, length/60, length%60)
	}

	fmt.Printf("Total length: %d:%02d\n", total/60, total%60)
}
