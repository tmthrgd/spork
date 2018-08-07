package dbus

import "context"

// GetPlaylistLength returns the number of songs in the active playlist.
func GetPlaylistLength(ctx context.Context) (length int32, err error) {
	return length, audObj.CallWithContext(ctx,
		"org.atheme.audacious.Length", 0).Store(&length)
}

// GetSongTitle returns the title of the given song in the active playlist.
func GetSongTitle(ctx context.Context, entry uint32) (title string, err error) {
	return title, audObj.CallWithContext(ctx,
		"org.atheme.audacious.SongTitle", 0, entry).Store(&title)
}

// GetSongLength returns the length, in seconds, of the given song in the active playlist.
func GetSongLength(ctx context.Context, entry uint32) (length int32, err error) {
	return length, audObj.CallWithContext(ctx,
		"org.atheme.audacious.SongLength", 0, entry).Store(&length)
}

// GetPlaylistPosition returns the currently selected song in the active playlist.
func GetPlaylistPosition(ctx context.Context) (position uint32, err error) {
	return position, audObj.CallWithContext(ctx,
		"org.atheme.audacious.Position", 0).Store(&position)
}

// GetPlaylistName returns the name of the active playlist.
func GetPlaylistName(ctx context.Context) (name string, err error) {
	return name, audObj.CallWithContext(ctx,
		"org.atheme.audacious.GetActivePlaylistName", 0).Store(&name)
}

// PlaylistJump switches the currently selected song in the active playlist.
func PlaylistJump(ctx context.Context, pos uint32) error {
	return audObj.CallWithContext(ctx, "org.atheme.audacious.Jump", 0, pos).Err
}

// Play starts playing the currently selected song in the active playlist.
func Play(ctx context.Context) error {
	return audObj.CallWithContext(ctx, "org.atheme.audacious.Play", 0).Err
}
