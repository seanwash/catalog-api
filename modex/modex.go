// Package modext is for SQLBoiler helper methods. Because the models are
// generated we need to put additional functionality in a place that won't be
// removed if the schema changes and the models are regenerated.
package modex

import (
	"context"
	"database/sql"

	generatedmodel "github.com/seanwash/catalog-api/graph/model"

	"github.com/seanwash/catalog-api/models"
)

// TrackRelationshipSetops handles adding new relationships for a given
// Track and TrackInput. Note that any IDs not supplied by the client will have
// their relationship removed.
func TrackRelationshipSetops(ctx context.Context, tx *sql.Tx, track *models.Track, input generatedmodel.TrackInput) error {
	artists, err := models.Artists(models.ArtistWhere.ID.IN(input.ArtistIds)).All(ctx, tx)

	if err != nil {
		return err
	}

	err = track.SetArtists(ctx, tx, false, artists...)
	if err != nil {
		return err
	}

	albums, err := models.Albums(models.AlbumWhere.ID.IN(input.AlbumIds)).All(ctx, tx)
	if err != nil {
		return err
	}

	err = track.SetAlbums(ctx, tx, false, albums...)
	if err != nil {
		return err
	}

	genres, err := models.Genres(models.GenreWhere.ID.IN(input.GenreIds)).All(ctx, tx)
	if err != nil {
		return err
	}

	err = track.SetGenres(ctx, tx, false, genres...)
	if err != nil {
		return err
	}

	return nil
}
