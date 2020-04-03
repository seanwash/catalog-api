package web

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/go-chi/chi"
	"github.com/seanwash/catalog-api/models"
	"github.com/volatiletech/sqlboiler/boil"
)

// HealthCheck responds with a 200 status. Used for zero downtime deploys and
// uptime monitors.
func (env *Env) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// TODO: Note that these are views and can be used to control the presentation
// in case there are fields that shouldn't be exposed to the client.
type trackWithRels struct {
	ID      int         `boil:"id" json:"id"`
	Name    string      `boil:"name" json:"name"`
	Artists []artistRel `boil:"artists" json:"artists"`
}

type artistRel struct {
	ID   int
	Name string
}

// TracksIndex returns a listing of Tracks.
func (env *Env) TracksIndex(w http.ResponseWriter, r *http.Request) {
	validateRequestMethod(w, r, http.MethodGet)

	tracks, err := models.Tracks(qm.Load(models.TrackRels.Artists)).All(context.Background(), env.DB)
	if err != nil {
		// TODO: Change this
		log.Fatal(err)
	}

	// TODO: There has to be a better way of doing this 😱
	var tracksWithRels []trackWithRels
	for _, track := range tracks {
		// Build associated ArtistRel structs.
		var artistRels []artistRel
		for _, artist := range track.R.Artists {
			artistRels = append(artistRels, artistRel{
				ID:   artist.ID,
				Name: artist.Name,
			})
		}

		// Build trackWithRels struct and assign Artists.
		tracksWithRels = append(tracksWithRels, trackWithRels{
			ID:      track.ID,
			Name:    track.Name,
			Artists: artistRels,
		})
	}

	json.NewEncoder(w).Encode(tracksWithRels)
}

// TracksShow returns a single Track matching a given ID.
func (env *Env) TracksShow(w http.ResponseWriter, r *http.Request) {
	validateRequestMethod(w, r, http.MethodGet)

	var track *models.Track

	if id := chi.URLParam(r, "id"); id != "" {
		intID, _ := strconv.Atoi(id)
		track, _ = models.FindTrack(context.Background(), env.DB, intID)
	}

	if track == nil {
		EncodeJSONError(w, errors.New("not found"), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(track)
}

// TracksCreate inserts and returns a new Track.
func (env *Env) TracksCreate(w http.ResponseWriter, r *http.Request) {
	validateRequestMethod(w, r, http.MethodPost)

	var payload map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		// This is a bit naive because not all errors returned by Decode are
		// related to an invalid payload. Also, err messages from Decode
		// probably aren't ideal for sending to the client.
		EncodeJSONError(w, err, http.StatusBadRequest)
	}

	// Initialize an empty Track struct & assign values from given Form vars.
	var track models.Track
	// Set the new Track's Name field, since payload's values are a blank
	// `interface{}` we need to use type assertions to specify that name is a
	// string.
	track.Name = payload["name"].(string)

	// Insert Track.
	if err := track.Insert(context.Background(), env.DB, boil.Infer()); err != nil {
		EncodeJSONError(w, err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(track)
}

// TracksUpdate mutates an existing Track and returns it.
func (env *Env) TracksUpdate(w http.ResponseWriter, r *http.Request) {
	validateRequestMethod(w, r, http.MethodPut)

	var track *models.Track

	if id := chi.URLParam(r, "id"); id != "" {
		intID, _ := strconv.Atoi(id)
		track, _ = models.FindTrack(context.Background(), env.DB, intID)
	}

	// If the Track to be updated is not found the client should be notified.
	if track == nil {
		EncodeJSONError(w, errors.New("not found"), http.StatusNotFound)
		return
	}

	var payload map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		// This is a bit naive because not all errors returned by Decode are
		// related to an invalid payload. Also, err messages from Decode
		// probably aren't ideal for sending to the client.
		EncodeJSONError(w, err, http.StatusBadRequest)
	}

	// Set the new Track's Name field, since payload's values are a blank
	// `interface{}` we need to use type assertions to specify that name is a
	// string.
	track.Name = payload["name"].(string)

	// Update the existing Track.
	if _, err := track.Update(context.Background(), env.DB, boil.Infer()); err != nil {
		EncodeJSONError(w, err, http.StatusBadRequest)
		return
	}

	// Status default to 200 OK.
	json.NewEncoder(w).Encode(track)
}

// TracksDelete removes an existing Track.
func (env *Env) TracksDelete(w http.ResponseWriter, r *http.Request) {
	validateRequestMethod(w, r, http.MethodDelete)

	var track *models.Track

	if id := chi.URLParam(r, "id"); id != "" {
		intID, _ := strconv.Atoi(id)
		track, _ = models.FindTrack(context.Background(), env.DB, intID)
	}

	// If the Track to be deleted is not found the client should be notified.
	if track == nil {
		EncodeJSONError(w, errors.New("not found"), http.StatusNotFound)
		return
	}

	if _, err := track.Delete(context.Background(), env.DB); err != nil {
		EncodeJSONError(w, err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func validateRequestMethod(w http.ResponseWriter, r *http.Request, validMethod string) {
	if r.Method != validMethod {
		w.Header().Set("Allow", validMethod)
		EncodeJSONError(w, errors.New("method not allowed"), http.StatusMethodNotAllowed)
		return
	}
}
