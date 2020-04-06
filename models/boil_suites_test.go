// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Albums", testAlbums)
	t.Run("Artists", testArtists)
	t.Run("Genres", testGenres)
	t.Run("Tracks", testTracks)
}

func TestDelete(t *testing.T) {
	t.Run("Albums", testAlbumsDelete)
	t.Run("Artists", testArtistsDelete)
	t.Run("Genres", testGenresDelete)
	t.Run("Tracks", testTracksDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Albums", testAlbumsQueryDeleteAll)
	t.Run("Artists", testArtistsQueryDeleteAll)
	t.Run("Genres", testGenresQueryDeleteAll)
	t.Run("Tracks", testTracksQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Albums", testAlbumsSliceDeleteAll)
	t.Run("Artists", testArtistsSliceDeleteAll)
	t.Run("Genres", testGenresSliceDeleteAll)
	t.Run("Tracks", testTracksSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Albums", testAlbumsExists)
	t.Run("Artists", testArtistsExists)
	t.Run("Genres", testGenresExists)
	t.Run("Tracks", testTracksExists)
}

func TestFind(t *testing.T) {
	t.Run("Albums", testAlbumsFind)
	t.Run("Artists", testArtistsFind)
	t.Run("Genres", testGenresFind)
	t.Run("Tracks", testTracksFind)
}

func TestBind(t *testing.T) {
	t.Run("Albums", testAlbumsBind)
	t.Run("Artists", testArtistsBind)
	t.Run("Genres", testGenresBind)
	t.Run("Tracks", testTracksBind)
}

func TestOne(t *testing.T) {
	t.Run("Albums", testAlbumsOne)
	t.Run("Artists", testArtistsOne)
	t.Run("Genres", testGenresOne)
	t.Run("Tracks", testTracksOne)
}

func TestAll(t *testing.T) {
	t.Run("Albums", testAlbumsAll)
	t.Run("Artists", testArtistsAll)
	t.Run("Genres", testGenresAll)
	t.Run("Tracks", testTracksAll)
}

func TestCount(t *testing.T) {
	t.Run("Albums", testAlbumsCount)
	t.Run("Artists", testArtistsCount)
	t.Run("Genres", testGenresCount)
	t.Run("Tracks", testTracksCount)
}

func TestInsert(t *testing.T) {
	t.Run("Albums", testAlbumsInsert)
	t.Run("Albums", testAlbumsInsertWhitelist)
	t.Run("Artists", testArtistsInsert)
	t.Run("Artists", testArtistsInsertWhitelist)
	t.Run("Genres", testGenresInsert)
	t.Run("Genres", testGenresInsertWhitelist)
	t.Run("Tracks", testTracksInsert)
	t.Run("Tracks", testTracksInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("AlbumToTracks", testAlbumToManyTracks)
	t.Run("ArtistToTracks", testArtistToManyTracks)
	t.Run("GenreToTracks", testGenreToManyTracks)
	t.Run("TrackToAlbums", testTrackToManyAlbums)
	t.Run("TrackToArtists", testTrackToManyArtists)
	t.Run("TrackToGenres", testTrackToManyGenres)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("AlbumToTracks", testAlbumToManyAddOpTracks)
	t.Run("ArtistToTracks", testArtistToManyAddOpTracks)
	t.Run("GenreToTracks", testGenreToManyAddOpTracks)
	t.Run("TrackToAlbums", testTrackToManyAddOpAlbums)
	t.Run("TrackToArtists", testTrackToManyAddOpArtists)
	t.Run("TrackToGenres", testTrackToManyAddOpGenres)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("AlbumToTracks", testAlbumToManySetOpTracks)
	t.Run("ArtistToTracks", testArtistToManySetOpTracks)
	t.Run("GenreToTracks", testGenreToManySetOpTracks)
	t.Run("TrackToAlbums", testTrackToManySetOpAlbums)
	t.Run("TrackToArtists", testTrackToManySetOpArtists)
	t.Run("TrackToGenres", testTrackToManySetOpGenres)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("AlbumToTracks", testAlbumToManyRemoveOpTracks)
	t.Run("ArtistToTracks", testArtistToManyRemoveOpTracks)
	t.Run("GenreToTracks", testGenreToManyRemoveOpTracks)
	t.Run("TrackToAlbums", testTrackToManyRemoveOpAlbums)
	t.Run("TrackToArtists", testTrackToManyRemoveOpArtists)
	t.Run("TrackToGenres", testTrackToManyRemoveOpGenres)
}

func TestReload(t *testing.T) {
	t.Run("Albums", testAlbumsReload)
	t.Run("Artists", testArtistsReload)
	t.Run("Genres", testGenresReload)
	t.Run("Tracks", testTracksReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Albums", testAlbumsReloadAll)
	t.Run("Artists", testArtistsReloadAll)
	t.Run("Genres", testGenresReloadAll)
	t.Run("Tracks", testTracksReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Albums", testAlbumsSelect)
	t.Run("Artists", testArtistsSelect)
	t.Run("Genres", testGenresSelect)
	t.Run("Tracks", testTracksSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Albums", testAlbumsUpdate)
	t.Run("Artists", testArtistsUpdate)
	t.Run("Genres", testGenresUpdate)
	t.Run("Tracks", testTracksUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Albums", testAlbumsSliceUpdateAll)
	t.Run("Artists", testArtistsSliceUpdateAll)
	t.Run("Genres", testGenresSliceUpdateAll)
	t.Run("Tracks", testTracksSliceUpdateAll)
}
