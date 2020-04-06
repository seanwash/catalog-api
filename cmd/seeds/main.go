package main

import (
	"context"
	"log"
	"time"

	"github.com/seanwash/catalog-api/models"
	"github.com/volatiletech/sqlboiler/boil"

	"github.com/joho/godotenv"
	"github.com/seanwash/catalog-api/db"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	dbConn := db.Connection()
	defer dbConn.Close()

	tx, err := dbConn.BeginTx(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	//
	// Genre
	//

	progressiveJazzGenre := &models.Genre{
		Name: "Progressive Jazz",
	}
	err = progressiveJazzGenre.Insert(context.Background(), tx, boil.Infer())

	instrumentalGenre := &models.Genre{
		Name: "Instrumental",
	}
	err = instrumentalGenre.Insert(context.Background(), tx, boil.Infer())

	//
	// Artists
	//

	pliniArtist := &models.Artist{
		Name: "Plini",
	}

	err = pliniArtist.Insert(context.Background(), tx, boil.Infer())

	nollyArtist := &models.Artist{
		Name: "Adam Nolly Getgood",
	}

	err = nollyArtist.Insert(context.Background(), tx, boil.Infer())

	anomalieArtist := &models.Artist{
		Name: "Anomalie",
	}

	err = anomalieArtist.Insert(context.Background(), tx, boil.Infer())

	//
	// Albums
	//

	sunheadAlbum := &models.Album{
		Name:       "Sunhead",
		AlbumType:  models.AlbumTypeAlbum,
		ReleasedAt: time.Now(),
	}

	err = sunheadAlbum.Insert(context.Background(), tx, boil.Infer())

	blueAngelAlbum := &models.Album{
		Name:       "Blue Angel",
		AlbumType:  models.AlbumTypeSingle,
		ReleasedAt: time.Now(),
	}

	err = blueAngelAlbum.Insert(context.Background(), tx, boil.Infer())

	//
	// Tracks - Sunhead
	//

	kindTrack := &models.Track{
		Name:       "Kind",
		DurationMS: 240000,
	}

	err = kindTrack.Insert(context.Background(), tx, boil.Infer())
	err = kindTrack.AddArtists(context.Background(), tx, false, []*models.Artist{pliniArtist}...)
	err = kindTrack.AddGenres(context.Background(), tx, false, []*models.Genre{progressiveJazzGenre, instrumentalGenre}...)

	saltAndCharcoalTrack := &models.Track{
		Name:       "Salt & Charcoal",
		DurationMS: 240000,
	}

	err = saltAndCharcoalTrack.Insert(context.Background(), tx, boil.Infer())
	err = saltAndCharcoalTrack.AddArtists(context.Background(), tx, false, []*models.Artist{pliniArtist}...)
	err = saltAndCharcoalTrack.AddGenres(context.Background(), tx, false, []*models.Genre{progressiveJazzGenre, instrumentalGenre}...)

	flaneurTrack := &models.Track{
		Name:       "Flaneur",
		DurationMS: 240000,
	}

	err = flaneurTrack.Insert(context.Background(), tx, boil.Infer())
	err = flaneurTrack.AddArtists(context.Background(), tx, false, []*models.Artist{pliniArtist, anomalieArtist}...)
	err = flaneurTrack.AddGenres(context.Background(), tx, false, []*models.Genre{progressiveJazzGenre, instrumentalGenre}...)

	sunheadTrack := &models.Track{
		Name:       "Sunhead",
		DurationMS: 240000,
	}

	err = sunheadTrack.Insert(context.Background(), tx, boil.Infer())
	err = sunheadTrack.AddArtists(context.Background(), tx, false, []*models.Artist{pliniArtist}...)
	err = sunheadTrack.AddGenres(context.Background(), tx, false, []*models.Genre{progressiveJazzGenre, instrumentalGenre}...)

	err = sunheadAlbum.AddTracks(context.Background(), tx, false, []*models.Track{kindTrack, saltAndCharcoalTrack, flaneurTrack, sunheadTrack}...)

	//
	// Tracks - Blue Angel
	//

	blueAngelTrack := &models.Track{
		Name:       "Blue Angel",
		DurationMS: 240000,
	}

	err = blueAngelTrack.Insert(context.Background(), tx, boil.Infer())
	err = blueAngelTrack.AddArtists(context.Background(), tx, false, []*models.Artist{pliniArtist, nollyArtist}...)
	err = blueAngelTrack.AddGenres(context.Background(), tx, false, []*models.Genre{progressiveJazzGenre, instrumentalGenre}...)

	err = blueAngelAlbum.AddTracks(context.Background(), tx, false, []*models.Track{blueAngelTrack}...)

	tx.Commit()
}
