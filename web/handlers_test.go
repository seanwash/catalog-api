package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/seanwash/catalog-api/db"
)

func TestTracksCreate(t *testing.T) {
	// TODO: Stub out actual DB writes.
	env := &Env{db: db.Connection()}
	payload, err := json.Marshal(map[string]interface{}{"name": "test"})

	r, err := http.NewRequest("POST", "http://localhost:4000/tracks", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatalf("Could not perform request: %v", err)
	}

	// TODO: Document why?
	w := httptest.NewRecorder()

	// Call handler that we're wanting to test.
	env.tracksCreate(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		t.Fatalf("Expected 201 status, got %v", res.StatusCode)
	}
}
