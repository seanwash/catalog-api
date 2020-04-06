// Package modext is for SQLBoiler helper methods. Because the models are
// generated we need to put additional functionality in a place that won't be
// removed if the schema changes and the models are regenerated.
package modex

import (
	"context"
	"database/sql"
	"strings"

	generatedmodel "github.com/seanwash/catalog-api/graph/model"

	"github.com/seanwash/catalog-api/models"
)

// TrackCreateRelationshipSetops handles inserting and adding new relationships
// for a given Track and TrackInput. If a relation contained in TrackInput has
// an ID then a new relation is inserted, and if an ID is missing then a new row
// and relation are inserted.
func TrackCreateRelationshipSetops(ctx context.Context, tx *sql.Tx, track *models.Track, input generatedmodel.TrackInput) error {
	// Construct Artists to be added and inserted. Adding will create new join
	// table associations for Tracks, and inserting will both create the joins
	// and insert new Artist records.
	var newArtistsToInsert []*models.Artist
	var existingArtistsToAdd []*models.Artist

	for _, artistInput := range input.Artists {
		if artistInput.ID == nil {
			// TODO: Can we automatically map the fields here?
			newArtistsToInsert = append(newArtistsToInsert, &models.Artist{
				Name: artistInput.Name,
			})
		} else {
			// Note that by simply using ID here, someone can send an ID with any
			// other attributes that don't match the resource with that ID and the
			// join will still be created.
			existingArtistsToAdd = append(existingArtistsToAdd, &models.Artist{
				ID: *artistInput.ID,
			})
		}
	}

	err := track.AddArtists(ctx, tx, true, newArtistsToInsert...)
	if err != nil {
		return err
	}

	err = track.AddArtists(ctx, tx, false, existingArtistsToAdd...)
	if err != nil {
		return err
	}

	// Construct Albums to be added and inserted. Adding will create new join
	// table associations for Tracks, and inserting will both create the joins
	// and insert new Album records.
	var newAlbumsToInsert []*models.Album
	var existingAlbumsToAdd []*models.Album

	for _, artistInput := range input.Albums {
		if artistInput.ID == nil {
			// TODO: Can we automatically map the fields here?
			newAlbumsToInsert = append(newAlbumsToInsert, &models.Album{
				Name:       artistInput.Name,
				ReleasedAt: artistInput.ReleasedAt,
				AlbumType:  strings.ToLower(artistInput.AlbumType.String()),
			})
		} else {
			// Note that by simply using ID here, someone can send an ID with any
			// other attributes that don't match the resource with that ID and the
			// join will still be created.
			existingAlbumsToAdd = append(existingAlbumsToAdd, &models.Album{
				ID: *artistInput.ID,
			})
		}
	}

	err = track.AddAlbums(ctx, tx, true, newAlbumsToInsert...)
	if err != nil {
		return err
	}

	err = track.AddAlbums(ctx, tx, false, existingAlbumsToAdd...)
	if err != nil {
		return err
	}

	// Construct Genres to be added and inserted. Adding will create new join
	// table associations for Tracks, and inserting will both create the joins
	// and insert new Genre records.
	var newGenresToInsert []*models.Genre
	var existingGenresToAdd []*models.Genre

	for _, artistInput := range input.Genres {
		if artistInput.ID == nil {
			// TODO: Can we automatically map the fields here?
			newGenresToInsert = append(newGenresToInsert, &models.Genre{
				Name: artistInput.Name,
			})
		} else {
			// Note that by simply using ID here, someone can send an ID with any
			// other attributes that don't match the resource with that ID and the
			// join will still be created.
			existingGenresToAdd = append(existingGenresToAdd, &models.Genre{
				ID: *artistInput.ID,
			})
		}
	}

	err = track.AddGenres(ctx, tx, true, newGenresToInsert...)
	if err != nil {
		return err
	}

	err = track.AddGenres(ctx, tx, false, existingGenresToAdd...)
	if err != nil {
		return err
	}

	return nil
}

// TrackUpdateRelationshipSetops handles inserting and adding new relationships
// for a given Track and TrackInput. If a relation contained in TrackInput has
// an ID then a new relation is inserted, and if an ID is missing then a new row
// and relation are inserted. Any already existing relationships that are not
// included in the TrackInput will be removed.
func TrackUpdateRelationshipSetops(ctx context.Context, tx *sql.Tx, track *models.Track, input generatedmodel.TrackInput) error {
	// Construct Artists to be added and inserted. Adding will create new join
	// table associations for Tracks, and inserting will both create the joins
	// and insert new Artist records.
	var newArtistsToInsert []*models.Artist
	var existingArtistsToSet []*models.Artist

	for _, artistInput := range input.Artists {
		if artistInput.ID == nil {
			// TODO: Can we automatically map the fields here?
			newArtistsToInsert = append(newArtistsToInsert, &models.Artist{
				Name: artistInput.Name,
			})
		} else {
			// Note that by simply using ID here, someone can send an ID with any
			// other attributes that don't match the resource with that ID and the
			// join will still be created.
			existingArtistsToSet = append(existingArtistsToSet, &models.Artist{
				ID: *artistInput.ID,
			})
		}
	}

	// Note that order matters here. We need to set before insert otherwise Setting
	// will remove any previously inserted artists since they're not included in the
	// `existingArtistsToSet` slice.
	err := track.SetArtists(ctx, tx, false, existingArtistsToSet...)
	if err != nil {
		return err
	}

	err = track.AddArtists(ctx, tx, true, newArtistsToInsert...)
	if err != nil {
		return err
	}

	// Construct Albums to be added and inserted. Adding will create new join
	// table associations for Tracks, and inserting will both create the joins
	// and insert new Album records.
	var newAlbumsToInsert []*models.Album
	var existingAlbumsToSet []*models.Album

	for _, artistInput := range input.Albums {
		if artistInput.ID == nil {
			// TODO: Can we automatically map the fields here?
			newAlbumsToInsert = append(newAlbumsToInsert, &models.Album{
				Name:       artistInput.Name,
				ReleasedAt: artistInput.ReleasedAt,
				AlbumType:  strings.ToLower(artistInput.AlbumType.String()),
			})
		} else {
			// Note that by simply using ID here, someone can send an ID with any
			// other attributes that don't match the resource with that ID and the
			// join will still be created.
			existingAlbumsToSet = append(existingAlbumsToSet, &models.Album{
				ID: *artistInput.ID,
			})
		}
	}

	// Note that order matters here. We need to set before insert otherwise Setting
	// will remove any previously inserted artists since they're not included in the
	// `existingAlbumsToSet` slice.
	err = track.SetAlbums(ctx, tx, false, existingAlbumsToSet...)
	if err != nil {
		return err
	}

	err = track.AddAlbums(ctx, tx, true, newAlbumsToInsert...)
	if err != nil {
		return err
	}

	// Construct Genres to be added and inserted. Adding will create new join
	// table associations for Tracks, and inserting will both create the joins
	// and insert new Genre records.
	var newGenresToInsert []*models.Genre
	var existingGenresToSet []*models.Genre

	for _, artistInput := range input.Genres {
		if artistInput.ID == nil {
			// TODO: Can we automatically map the fields here?
			newGenresToInsert = append(newGenresToInsert, &models.Genre{
				Name: artistInput.Name,
			})
		} else {
			// Note that by simply using ID here, someone can send an ID with any
			// other attributes that don't match the resource with that ID and the
			// join will still be created.
			existingGenresToSet = append(existingGenresToSet, &models.Genre{
				ID: *artistInput.ID,
			})
		}
	}

	// Note that order matters here. We need to set before insert otherwise Setting
	// will remove any previously inserted artists since they're not included in the
	// `existingGenresToSet` slice.
	err = track.SetGenres(ctx, tx, false, existingGenresToSet...)
	if err != nil {
		return err
	}

	err = track.AddGenres(ctx, tx, true, newGenresToInsert...)
	if err != nil {
		return err
	}

	return nil
}
