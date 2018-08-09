package dbus

import (
	"context"

	"github.com/godbus/dbus"
)

// GetPlaylistLength returns the number of songs in the active playlist.
func GetPlaylistLength(ctx context.Context) (length int32, err error) {
	return length, audObj.CallWithContext(ctx,
		"org.atheme.audacious.Length", 0).Store(&length)
}

// GetSongTitle returns the formatted title of the given song in the active playlist.
func GetSongTitle(ctx context.Context, entry uint32) (title string, err error) {
	return title, audObj.CallWithContext(ctx,
		"org.atheme.audacious.SongTitle", 0, entry).Store(&title)
}

// GetSongLength returns the length, in seconds, of the given song in the active playlist.
func GetSongLength(ctx context.Context, entry uint32) (length int32, err error) {
	return length, audObj.CallWithContext(ctx,
		"org.atheme.audacious.SongLength", 0, entry).Store(&length)
}

// GetSongName returns the unformatted title of the given song in the active playlist.
func GetSongName(ctx context.Context, entry uint32) (title string, err error) {
	return title, audObj.CallWithContext(ctx,
		"org.atheme.audacious.SongTuple", 0, entry, "title").Store(&title)
}

// GetSongArtist returns the artist of the given song in the active playlist.
func GetSongArtist(ctx context.Context, entry uint32) (artist string, err error) {
	return artist, audObj.CallWithContext(ctx,
		"org.atheme.audacious.SongTuple", 0, entry, "artist").Store(&artist)
}

// GetSongAlbum returns the album of the given song in the active playlist.
func GetSongAlbum(ctx context.Context, entry uint32) (album string, err error) {
	return album, audObj.CallWithContext(ctx,
		"org.atheme.audacious.SongTuple", 0, entry, "album").Store(&album)
}

// GetSongYear returns the year of the given song in the active playlist.
func GetSongYear(ctx context.Context, entry uint32) (year int32, err error) {
	// Audacious may return a D-Bus string when a year is missing, so
	// cast manually here.
	var yearVar dbus.Variant
	err = audObj.CallWithContext(ctx,
		"org.atheme.audacious.SongTuple", 0, entry, "year").Store(&yearVar)
	year, _ = yearVar.Value().(int32)
	return year, err
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

// PlayPause either starts playback of the currently selected song or pauses
// playback of the currently playing song.
func PlayPause(ctx context.Context) error {
	return audObj.CallWithContext(ctx, "org.atheme.audacious.PlayPause", 0).Err
}

// Stop stops playback of the currently playing song.
func Stop(ctx context.Context) error {
	return audObj.CallWithContext(ctx, "org.atheme.audacious.Stop", 0).Err
}

// GetVolume returns the volume of the media player.
func GetVolume(ctx context.Context) (volume int32, err error) {
	var unused int32
	return volume, audObj.CallWithContext(ctx,
		"org.atheme.audacious.Volume", 0).Store(&volume, &unused)
}

// SetVolume sets the volume of the media player to the given value.
func SetVolume(ctx context.Context, volume int32) error {
	return audObj.CallWithContext(ctx,
		"org.atheme.audacious.SetVolume", 0, volume, volume).Err
}

// Advance skips ahead one song in the current playlist.
func Advance(ctx context.Context) error {
	return audObj.CallWithContext(ctx, "org.atheme.audacious.Advance", 0).Err
}

// Reverse skips backwards one song in the current playlist.
func Reverse(ctx context.Context) error {
	return audObj.CallWithContext(ctx, "org.atheme.audacious.Reverse", 0).Err
}

// ToggleRepeat toggles the value of the repeat button.
func ToggleRepeat(ctx context.Context) error {
	return audObj.CallWithContext(ctx, "org.atheme.audacious.ToggleRepeat", 0).Err
}

// ToggleShuffle toggles the value of the shuffle button.
func ToggleShuffle(ctx context.Context) error {
	return audObj.CallWithContext(ctx, "org.atheme.audacious.ToggleShuffle", 0).Err
}

// GetRepeat toggles the value of the repeat button.
func GetRepeat(ctx context.Context) (repeat bool, err error) {
	return repeat, audObj.CallWithContext(ctx,
		"org.atheme.audacious.Repeat", 0).Store(&repeat)
}

// GetShuffle toggles the value of the shuffle button.
func GetShuffle(ctx context.Context) (shuffle bool, err error) {
	return shuffle, audObj.CallWithContext(ctx,
		"org.atheme.audacious.Shuffle", 0).Store(&shuffle)
}
