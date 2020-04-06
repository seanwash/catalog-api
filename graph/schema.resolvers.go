package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strings"

	"github.com/seanwash/catalog-api/graph/generated"
	generatedmodel "github.com/seanwash/catalog-api/graph/model"
	"github.com/seanwash/catalog-api/models"
	"github.com/seanwash/catalog-api/modex"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func (r *albumResolver) AlbumType(ctx context.Context, obj *models.Album) (generatedmodel.AlbumType, error) {
	// Postgres enums are lowercase, GraphQL enums are uppercase.
	albumType := generatedmodel.AlbumType(strings.ToUpper(obj.AlbumType))

	return albumType, nil
}

func (r *albumResolver) Tracks(ctx context.Context, obj *models.Album) ([]*models.Track, error) {
	album, err := models.Albums(models.AlbumWhere.ID.EQ(obj.ID), qm.Load(models.AlbumRels.Tracks)).One(ctx, r.DB)
	if err != nil {
		return nil, err
	}

	return album.R.Tracks, nil
}

func (r *mutationResolver) TrackCreate(ctx context.Context, input generatedmodel.TrackInput) (*models.Track, error) {
	if err := ensureUserIsAdmin(ctx); err != nil {
		return nil, err
	}

	track := &models.Track{
		Name:       strings.TrimSpace(input.Name),
		DurationMS: input.DurationMs,
	}

	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	err = track.Insert(ctx, tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = modex.TrackRelationshipSetops(ctx, tx, track, input)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return track, nil
}

func (r *mutationResolver) TrackUpdate(ctx context.Context, id int, input generatedmodel.TrackInput) (*models.Track, error) {
	if err := ensureUserIsAdmin(ctx); err != nil {
		return nil, err
	}

	track, err := models.FindTrack(ctx, r.DB, id)
	if track == nil {
		return nil, gqlerror.Errorf("No Tracks found with ID %d", id)
	}
	if err != nil {
		return nil, err
	}

	track.Name = strings.TrimSpace(input.Name)
	track.DurationMS = input.DurationMs

	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	_, err = track.Update(ctx, tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = modex.TrackRelationshipSetops(ctx, tx, track, input)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return track, nil
}

func (r *mutationResolver) TrackDelete(ctx context.Context, id int) (*models.Track, error) {
	if err := ensureUserIsAdmin(ctx); err != nil {
		return nil, err
	}

	track, err := models.FindTrack(ctx, r.DB, id)
	if track == nil {
		return nil, gqlerror.Errorf("No Tracks found with ID %d", id)
	}
	if err != nil {
		return nil, err
	}

	_, err = track.Delete(ctx, r.DB)
	if err != nil {
		return nil, err
	}

	return track, nil
}

func (r *mutationResolver) ArtistCreate(ctx context.Context, input generatedmodel.ArtistInput) (*models.Artist, error) {
	if err := ensureUserIsAdmin(ctx); err != nil {
		return nil, err
	}

	artist := &models.Artist{
		Name: strings.TrimSpace(input.Name),
	}

	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	err = artist.Insert(ctx, tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return artist, nil
}

func (r *mutationResolver) ArtistUpdate(ctx context.Context, id int, input generatedmodel.ArtistInput) (*models.Artist, error) {
	if err := ensureUserIsAdmin(ctx); err != nil {
		return nil, err
	}

	artist, err := models.FindArtist(ctx, r.DB, id)
	if artist == nil {
		return nil, gqlerror.Errorf("No Artists found with ID %d", id)
	}
	if err != nil {
		return nil, err
	}

	artist.Name = strings.TrimSpace(input.Name)

	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	_, err = artist.Update(ctx, tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return artist, nil
}

func (r *mutationResolver) ArtistDelete(ctx context.Context, id int) (*models.Artist, error) {
	if err := ensureUserIsAdmin(ctx); err != nil {
		return nil, err
	}

	artist, err := models.FindArtist(ctx, r.DB, id)
	if artist == nil {
		return nil, gqlerror.Errorf("No Artist found with ID %d", id)
	}
	if err != nil {
		return nil, err
	}

	_, err = artist.Delete(ctx, r.DB)
	if err != nil {
		return nil, err
	}

	return artist, nil
}

func (r *mutationResolver) AlbumCreate(ctx context.Context, input generatedmodel.AlbumInput) (*models.Album, error) {
	if err := ensureUserIsAdmin(ctx); err != nil {
		return nil, err
	}

	album := &models.Album{
		Name:       strings.TrimSpace(input.Name),
		ReleasedAt: input.ReleasedAt,
		// Postgres enums are lowercase, GraphQL enums are uppercase.
		AlbumType: strings.ToLower(input.AlbumType.String()),
	}

	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	err = album.Insert(ctx, tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = modex.AlbumRelationshipSetops(ctx, tx, album, input)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return album, nil
}

func (r *mutationResolver) AlbumUpdate(ctx context.Context, id int, input generatedmodel.AlbumInput) (*models.Album, error) {
	if err := ensureUserIsAdmin(ctx); err != nil {
		return nil, err
	}

	album, err := models.FindAlbum(ctx, r.DB, id)
	if album == nil {
		return nil, gqlerror.Errorf("No Albums found with ID %d", id)
	}
	if err != nil {
		return nil, err
	}

	album.Name = strings.TrimSpace(input.Name)
	album.ReleasedAt = input.ReleasedAt
	// Postgres enums are lowercase, GraphQL enums are uppercase.
	album.AlbumType = strings.ToLower(input.AlbumType.String())

	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	_, err = album.Update(ctx, tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = modex.AlbumRelationshipSetops(ctx, tx, album, input)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return album, nil
}

func (r *mutationResolver) AlbumDelete(ctx context.Context, id int) (*models.Album, error) {
	if err := ensureUserIsAdmin(ctx); err != nil {
		return nil, err
	}

	album, err := models.FindAlbum(ctx, r.DB, id)
	if album == nil {
		return nil, gqlerror.Errorf("No Albums found with ID %d", id)
	}
	if err != nil {
		return nil, err
	}

	// Report error in the case of a failure other than an invalid `id`.
	if err != nil {
		return nil, err
	}

	_, err = album.Delete(ctx, r.DB)
	if err != nil {
		return nil, err
	}

	return album, nil
}

func (r *mutationResolver) GenreCreate(ctx context.Context, input generatedmodel.GenreInput) (*models.Genre, error) {
	if err := ensureUserIsAdmin(ctx); err != nil {
		return nil, err
	}

	genre := &models.Genre{
		Name: strings.TrimSpace(input.Name),
	}

	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	err = genre.Insert(ctx, tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return genre, nil
}

func (r *mutationResolver) GenreUpdate(ctx context.Context, id int, input generatedmodel.GenreInput) (*models.Genre, error) {
	if err := ensureUserIsAdmin(ctx); err != nil {
		return nil, err
	}

	genre, err := models.FindGenre(ctx, r.DB, id)
	if genre == nil {
		return nil, gqlerror.Errorf("No Genres found with ID %d", id)
	}
	if err != nil {
		return nil, err
	}

	genre.Name = strings.TrimSpace(input.Name)

	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	_, err = genre.Update(ctx, tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return genre, nil
}

func (r *mutationResolver) GenreDelete(ctx context.Context, id int) (*models.Genre, error) {
	if err := ensureUserIsAdmin(ctx); err != nil {
		return nil, err
	}

	genre, err := models.FindGenre(ctx, r.DB, id)
	if genre == nil {
		return nil, gqlerror.Errorf("No Genres found with ID %d", id)
	}
	if err != nil {
		return nil, err
	}

	_, err = genre.Delete(ctx, r.DB)
	if err != nil {
		return nil, err
	}

	return genre, nil
}

func (r *queryResolver) Tracks(ctx context.Context, limit *int, offset *int, filter *generatedmodel.TrackFilter) ([]*models.Track, error) {
	// Default query mods
	mods := []qm.QueryMod{
		qm.OrderBy("created_at desc"),
	}

	// Set default limit if no limit is specified.
	if limit == nil {
		mods = append(mods, qm.Limit(25))
	} else {
		mods = append(mods, qm.Limit(*limit))
	}

	if offset != nil {
		mods = append(mods, qm.Offset(*offset))
	}

	if filter != nil && filter.Name != nil {
		// Build an inclusive search string.
		query := strings.Join(strings.Split(*filter.Name, " "), " &")
		mods = append(mods, qm.Where("to_tsvector(name) @@ to_tsquery(?)", query))
	}

	if filter != nil && filter.ArtistID != nil {
		mods = append(mods,
			qm.InnerJoin("track_artists ta on ta.track_id = tracks.id"),
			qm.Where("ta.artist_id = ?", *filter.ArtistID))
	}

	if filter != nil && filter.ArtistName != nil {
		query := strings.Join(strings.Split(*filter.ArtistName, " "), " &")
		mods = append(mods,
			qm.InnerJoin("track_artists ta on ta.track_id = tracks.id"),
			qm.InnerJoin("artists a on ta.artist_id = a.id"),
			qm.Where("to_tsvector(a.name) @@ to_tsquery(?)", query))
	}

	tracks, err := models.Tracks(mods...).All(ctx, r.DB)
	if err != nil {
		return nil, err
	}

	return tracks, nil
}

func (r *queryResolver) Track(ctx context.Context, id int) (*models.Track, error) {
	track, err := models.FindTrack(ctx, r.DB, id)
	// In most cases here, `err` will be reporting that no rows were found if
	// an invalid `id` is given. This is just a friendlier message.
	if track == nil {
		return nil, gqlerror.Errorf("No Track found with ID %d", id)
	}

	// Report error in the case of a failure other than an invalid `id`.
	if err != nil {
		return nil, err
	}

	return track, nil
}

func (r *queryResolver) Artists(ctx context.Context, limit *int, offset *int, filter *generatedmodel.ArtistFilter) ([]*models.Artist, error) {
	mods := []qm.QueryMod{
		qm.OrderBy("created_at desc"),
	}

	// Set default limit if no limit is specified.
	if limit == nil {
		mods = append(mods, qm.Limit(25))
	} else {
		mods = append(mods, qm.Limit(*limit))
	}

	if offset != nil {
		mods = append(mods, qm.Offset(*offset))
	}

	if filter != nil && filter.Name != nil {
		// Build an inclusive search string.
		query := strings.Join(strings.Split(*filter.Name, " "), " &")
		mods = append(mods, qm.Where("to_tsvector(name) @@ to_tsquery(?)", query))
	}

	artists, err := models.Artists(mods...).All(ctx, r.DB)
	if err != nil {
		return nil, err
	}

	return artists, nil
}

func (r *queryResolver) Artist(ctx context.Context, id int) (*models.Artist, error) {
	artist, err := models.FindArtist(ctx, r.DB, id)
	// In most cases here, `err` will be reporting that no rows were found if
	// an invalid `id` is given. This is just a friendlier message.
	if artist == nil {
		return nil, gqlerror.Errorf("No Artist found with ID %d", id)
	}

	// Report error in the case of a failure other than an invalid `id`.
	if err != nil {
		return nil, err
	}

	return artist, nil
}

func (r *queryResolver) Albums(ctx context.Context, limit *int, offset *int, filter *generatedmodel.AlbumFilter) ([]*models.Album, error) {
	mods := []qm.QueryMod{
		qm.OrderBy("created_at desc"),
	}
	// Set default limit if no limit is specified.
	if limit == nil {
		mods = append(mods, qm.Limit(25))
	} else {
		mods = append(mods, qm.Limit(*limit))
	}

	if offset != nil {
		mods = append(mods, qm.Offset(*offset))
	}

	if filter != nil && filter.Name != nil {
		// Build an inclusive search string.
		query := strings.Join(strings.Split(*filter.Name, " "), " &")
		mods = append(mods, qm.Where("to_tsvector(name) @@ to_tsquery(?)", query))
	}

	if filter != nil && filter.AlbumType != nil {
		// Postgres enums are lower case and filter.AlbumType is uppercase.
		mods = append(mods, qm.Where("album_type = ?", strings.ToLower(filter.AlbumType.String())))
	}

	artists, err := models.Albums(mods...).All(ctx, r.DB)
	if err != nil {
		return nil, err
	}

	return artists, nil
}

func (r *queryResolver) Album(ctx context.Context, id int) (*models.Album, error) {
	album, err := models.FindAlbum(ctx, r.DB, id)
	// In most cases here, `err` will be reporting that no rows were found if
	// an invalid `id` is given. This is just a friendlier message.
	if album == nil {
		return nil, gqlerror.Errorf("No Track found with ID %d", id)
	}

	// Report error in the case of a failure other than an invalid `id`.
	if err != nil {
		return nil, err
	}

	return album, nil
}

func (r *queryResolver) Genres(ctx context.Context, limit *int, offset *int) ([]*models.Genre, error) {
	mods := []qm.QueryMod{
		qm.OrderBy("created_at desc"),
	}

	// Set default limit if no limit is specified.
	if limit == nil {
		mods = append(mods, qm.Limit(25))
	} else {
		mods = append(mods, qm.Limit(*limit))
	}

	if offset != nil {
		mods = append(mods, qm.Offset(*offset))
	}

	genres, err := models.Genres(mods...).All(ctx, r.DB)
	if err != nil {
		return nil, err
	}

	return genres, nil
}

func (r *queryResolver) Genre(ctx context.Context, id int) (*models.Genre, error) {
	genre, err := models.FindGenre(ctx, r.DB, id)
	// In most cases here, `err` will be reporting that no rows were found if
	// an invalid `id` is given. This is just a friendlier message.
	if genre == nil {
		return nil, gqlerror.Errorf("No Genre found with ID %d", id)
	}

	// Report error in the case of a failure other than an invalid `id`.
	if err != nil {
		return nil, err
	}

	return genre, nil
}

func (r *trackResolver) Artists(ctx context.Context, obj *models.Track) ([]*models.Artist, error) {
	track, err := models.Tracks(models.TrackWhere.ID.EQ(obj.ID), qm.Load(models.TrackRels.Artists)).One(ctx, r.DB)
	if err != nil {
		return nil, err
	}

	return track.R.Artists, nil
}

func (r *trackResolver) Albums(ctx context.Context, obj *models.Track) ([]*models.Album, error) {
	track, err := models.Tracks(models.TrackWhere.ID.EQ(obj.ID), qm.Load(models.TrackRels.Albums)).One(ctx, r.DB)
	if err != nil {
		return nil, err
	}

	return track.R.Albums, nil
}

func (r *trackResolver) Genres(ctx context.Context, obj *models.Track) ([]*models.Genre, error) {
	track, err := models.Tracks(models.TrackWhere.ID.EQ(obj.ID), qm.Load(models.TrackRels.Genres)).One(ctx, r.DB)
	if err != nil {
		return nil, err
	}

	return track.R.Genres, nil
}

// Album returns generated.AlbumResolver implementation.
func (r *Resolver) Album() generated.AlbumResolver { return &albumResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Track returns generated.TrackResolver implementation.
func (r *Resolver) Track() generated.TrackResolver { return &trackResolver{r} }

type albumResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type trackResolver struct{ *Resolver }
