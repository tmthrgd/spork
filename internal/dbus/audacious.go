package dbus

import "context"

func (b *Bus) GetPlaylistLength(ctx context.Context) (length int32, err error) {
	return length, b.audaciousObject.CallWithContext(ctx,
		"org.atheme.audacious.Length", 0).Store(&length)
}

func (b *Bus) GetSongTitle(ctx context.Context, entry uint32) (title string, err error) {
	return title, b.audaciousObject.CallWithContext(ctx,
		"org.atheme.audacious.SongTitle", 0, entry).Store(&title)
}

func (b *Bus) GetSongLength(ctx context.Context, entry uint32) (length int32, err error) {
	return length, b.audaciousObject.CallWithContext(ctx,
		"org.atheme.audacious.SongLength", 0, entry).Store(&length)
}
