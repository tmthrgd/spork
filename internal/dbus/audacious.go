package dbus

import "context"

func GetPlaylistLength(ctx context.Context) (length int32, err error) {
	return length, audObj.CallWithContext(ctx,
		"org.atheme.audacious.Length", 0).Store(&length)
}

func GetSongTitle(ctx context.Context, entry uint32) (title string, err error) {
	return title, audObj.CallWithContext(ctx,
		"org.atheme.audacious.SongTitle", 0, entry).Store(&title)
}

func GetSongLength(ctx context.Context, entry uint32) (length int32, err error) {
	return length, audObj.CallWithContext(ctx,
		"org.atheme.audacious.SongLength", 0, entry).Store(&length)
}

func GetPlaylistPosition(ctx context.Context) (position uint32, err error) {
	return position, audObj.CallWithContext(ctx,
		"org.atheme.audacious.Position", 0).Store(&position)
}

func GetPlaylistName(ctx context.Context) (name string, err error) {
	return name, audObj.CallWithContext(ctx,
		"org.atheme.audacious.GetActivePlaylistName", 0).Store(&name)
}

func PlaylistJump(ctx context.Context, pos uint32) error {
	return audObj.CallWithContext(ctx, "org.atheme.audacious.Jump", 0, pos).Err
}

func Play(ctx context.Context) error {
	return audObj.CallWithContext(ctx, "org.atheme.audacious.Play", 0).Err
}
