// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testGenres(t *testing.T) {
	t.Parallel()

	query := Genres()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testGenresDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Genre{}
	if err = randomize.Struct(seed, o, genreDBTypes, true, genreColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genre struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Genres().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGenresQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Genre{}
	if err = randomize.Struct(seed, o, genreDBTypes, true, genreColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genre struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Genres().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Genres().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGenresSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Genre{}
	if err = randomize.Struct(seed, o, genreDBTypes, true, genreColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genre struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := GenreSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Genres().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGenresExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Genre{}
	if err = randomize.Struct(seed, o, genreDBTypes, true, genreColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genre struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := GenreExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Genre exists: %s", err)
	}
	if !e {
		t.Errorf("Expected GenreExists to return true, but got false.")
	}
}

func testGenresFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Genre{}
	if err = randomize.Struct(seed, o, genreDBTypes, true, genreColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genre struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	genreFound, err := FindGenre(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if genreFound == nil {
		t.Error("want a record, got nil")
	}
}

func testGenresBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Genre{}
	if err = randomize.Struct(seed, o, genreDBTypes, true, genreColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genre struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Genres().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testGenresOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Genre{}
	if err = randomize.Struct(seed, o, genreDBTypes, true, genreColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genre struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Genres().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testGenresAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	genreOne := &Genre{}
	genreTwo := &Genre{}
	if err = randomize.Struct(seed, genreOne, genreDBTypes, false, genreColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genre struct: %s", err)
	}
	if err = randomize.Struct(seed, genreTwo, genreDBTypes, false, genreColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genre struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = genreOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = genreTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Genres().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testGenresCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	genreOne := &Genre{}
	genreTwo := &Genre{}
	if err = randomize.Struct(seed, genreOne, genreDBTypes, false, genreColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genre struct: %s", err)
	}
	if err = randomize.Struct(seed, genreTwo, genreDBTypes, false, genreColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genre struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = genreOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = genreTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Genres().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func genreBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Genre) error {
	*o = Genre{}
	return nil
}

func genreAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Genre) error {
	*o = Genre{}
	return nil
}

func genreAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Genre) error {
	*o = Genre{}
	return nil
}

func genreBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Genre) error {
	*o = Genre{}
	return nil
}

func genreAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Genre) error {
	*o = Genre{}
	return nil
}

func genreBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Genre) error {
	*o = Genre{}
	return nil
}

func genreAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Genre) error {
	*o = Genre{}
	return nil
}

func genreBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Genre) error {
	*o = Genre{}
	return nil
}

func genreAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Genre) error {
	*o = Genre{}
	return nil
}

func testGenresHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Genre{}
	o := &Genre{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, genreDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Genre object: %s", err)
	}

	AddGenreHook(boil.BeforeInsertHook, genreBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	genreBeforeInsertHooks = []GenreHook{}

	AddGenreHook(boil.AfterInsertHook, genreAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	genreAfterInsertHooks = []GenreHook{}

	AddGenreHook(boil.AfterSelectHook, genreAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	genreAfterSelectHooks = []GenreHook{}

	AddGenreHook(boil.BeforeUpdateHook, genreBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	genreBeforeUpdateHooks = []GenreHook{}

	AddGenreHook(boil.AfterUpdateHook, genreAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	genreAfterUpdateHooks = []GenreHook{}

	AddGenreHook(boil.BeforeDeleteHook, genreBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	genreBeforeDeleteHooks = []GenreHook{}

	AddGenreHook(boil.AfterDeleteHook, genreAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	genreAfterDeleteHooks = []GenreHook{}

	AddGenreHook(boil.BeforeUpsertHook, genreBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	genreBeforeUpsertHooks = []GenreHook{}

	AddGenreHook(boil.AfterUpsertHook, genreAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	genreAfterUpsertHooks = []GenreHook{}
}

func testGenresInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Genre{}
	if err = randomize.Struct(seed, o, genreDBTypes, true, genreColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genre struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Genres().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGenresInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Genre{}
	if err = randomize.Struct(seed, o, genreDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Genre struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(genreColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Genres().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGenreToManyTracks(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Genre
	var b, c Track

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, genreDBTypes, true, genreColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genre struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, trackDBTypes, false, trackColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, trackDBTypes, false, trackColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	_, err = tx.Exec("insert into \"track_genres\" (\"genre_id\", \"track_id\") values ($1, $2)", a.ID, b.ID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = tx.Exec("insert into \"track_genres\" (\"genre_id\", \"track_id\") values ($1, $2)", a.ID, c.ID)
	if err != nil {
		t.Fatal(err)
	}

	check, err := a.Tracks().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ID == b.ID {
			bFound = true
		}
		if v.ID == c.ID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := GenreSlice{&a}
	if err = a.L.LoadTracks(ctx, tx, false, (*[]*Genre)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Tracks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Tracks = nil
	if err = a.L.LoadTracks(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Tracks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testGenreToManyAddOpTracks(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Genre
	var b, c, d, e Track

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, genreDBTypes, false, strmangle.SetComplement(genrePrimaryKeyColumns, genreColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Track{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, trackDBTypes, false, strmangle.SetComplement(trackPrimaryKeyColumns, trackColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Track{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddTracks(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if first.R.Genres[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}
		if second.R.Genres[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}

		if a.R.Tracks[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Tracks[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Tracks().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testGenreToManySetOpTracks(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Genre
	var b, c, d, e Track

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, genreDBTypes, false, strmangle.SetComplement(genrePrimaryKeyColumns, genreColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Track{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, trackDBTypes, false, strmangle.SetComplement(trackPrimaryKeyColumns, trackColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetTracks(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Tracks().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetTracks(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Tracks().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	// The following checks cannot be implemented since we have no handle
	// to these when we call Set(). Leaving them here as wishful thinking
	// and to let people know there's dragons.
	//
	// if len(b.R.Genres) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	// if len(c.R.Genres) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	if d.R.Genres[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}
	if e.R.Genres[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}

	if a.R.Tracks[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Tracks[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testGenreToManyRemoveOpTracks(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Genre
	var b, c, d, e Track

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, genreDBTypes, false, strmangle.SetComplement(genrePrimaryKeyColumns, genreColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Track{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, trackDBTypes, false, strmangle.SetComplement(trackPrimaryKeyColumns, trackColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddTracks(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Tracks().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveTracks(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Tracks().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if len(b.R.Genres) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if len(c.R.Genres) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if d.R.Genres[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Genres[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if len(a.R.Tracks) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Tracks[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Tracks[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testGenresReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Genre{}
	if err = randomize.Struct(seed, o, genreDBTypes, true, genreColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genre struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testGenresReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Genre{}
	if err = randomize.Struct(seed, o, genreDBTypes, true, genreColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genre struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := GenreSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testGenresSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Genre{}
	if err = randomize.Struct(seed, o, genreDBTypes, true, genreColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genre struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Genres().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	genreDBTypes = map[string]string{`ID`: `integer`, `Name`: `text`}
	_            = bytes.MinRead
)

func testGenresUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(genrePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(genreAllColumns) == len(genrePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Genre{}
	if err = randomize.Struct(seed, o, genreDBTypes, true, genreColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genre struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Genres().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, genreDBTypes, true, genrePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Genre struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testGenresSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(genreAllColumns) == len(genrePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Genre{}
	if err = randomize.Struct(seed, o, genreDBTypes, true, genreColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genre struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Genres().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, genreDBTypes, true, genrePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Genre struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(genreAllColumns, genrePrimaryKeyColumns) {
		fields = genreAllColumns
	} else {
		fields = strmangle.SetComplement(
			genreAllColumns,
			genrePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := GenreSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testGenresUpsert(t *testing.T) {
	t.Parallel()

	if len(genreAllColumns) == len(genrePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Genre{}
	if err = randomize.Struct(seed, &o, genreDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Genre struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Genre: %s", err)
	}

	count, err := Genres().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, genreDBTypes, false, genrePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Genre struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Genre: %s", err)
	}

	count, err = Genres().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
