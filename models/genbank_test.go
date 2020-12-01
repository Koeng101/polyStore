// Code generated by SQLBoiler 4.3.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testGenbanks(t *testing.T) {
	t.Parallel()

	query := Genbanks()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testGenbanksDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Genbank{}
	if err = randomize.Struct(seed, o, genbankDBTypes, true, genbankColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genbank struct: %s", err)
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

	count, err := Genbanks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGenbanksQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Genbank{}
	if err = randomize.Struct(seed, o, genbankDBTypes, true, genbankColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genbank struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Genbanks().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Genbanks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGenbanksSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Genbank{}
	if err = randomize.Struct(seed, o, genbankDBTypes, true, genbankColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genbank struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := GenbankSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Genbanks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGenbanksExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Genbank{}
	if err = randomize.Struct(seed, o, genbankDBTypes, true, genbankColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genbank struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := GenbankExists(ctx, tx, o.Accession)
	if err != nil {
		t.Errorf("Unable to check if Genbank exists: %s", err)
	}
	if !e {
		t.Errorf("Expected GenbankExists to return true, but got false.")
	}
}

func testGenbanksFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Genbank{}
	if err = randomize.Struct(seed, o, genbankDBTypes, true, genbankColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genbank struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	genbankFound, err := FindGenbank(ctx, tx, o.Accession)
	if err != nil {
		t.Error(err)
	}

	if genbankFound == nil {
		t.Error("want a record, got nil")
	}
}

func testGenbanksBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Genbank{}
	if err = randomize.Struct(seed, o, genbankDBTypes, true, genbankColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genbank struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Genbanks().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testGenbanksOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Genbank{}
	if err = randomize.Struct(seed, o, genbankDBTypes, true, genbankColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genbank struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Genbanks().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testGenbanksAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	genbankOne := &Genbank{}
	genbankTwo := &Genbank{}
	if err = randomize.Struct(seed, genbankOne, genbankDBTypes, false, genbankColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genbank struct: %s", err)
	}
	if err = randomize.Struct(seed, genbankTwo, genbankDBTypes, false, genbankColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genbank struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = genbankOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = genbankTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Genbanks().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testGenbanksCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	genbankOne := &Genbank{}
	genbankTwo := &Genbank{}
	if err = randomize.Struct(seed, genbankOne, genbankDBTypes, false, genbankColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genbank struct: %s", err)
	}
	if err = randomize.Struct(seed, genbankTwo, genbankDBTypes, false, genbankColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genbank struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = genbankOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = genbankTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Genbanks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func genbankBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Genbank) error {
	*o = Genbank{}
	return nil
}

func genbankAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Genbank) error {
	*o = Genbank{}
	return nil
}

func genbankAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Genbank) error {
	*o = Genbank{}
	return nil
}

func genbankBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Genbank) error {
	*o = Genbank{}
	return nil
}

func genbankAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Genbank) error {
	*o = Genbank{}
	return nil
}

func genbankBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Genbank) error {
	*o = Genbank{}
	return nil
}

func genbankAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Genbank) error {
	*o = Genbank{}
	return nil
}

func genbankBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Genbank) error {
	*o = Genbank{}
	return nil
}

func genbankAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Genbank) error {
	*o = Genbank{}
	return nil
}

func testGenbanksHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Genbank{}
	o := &Genbank{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, genbankDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Genbank object: %s", err)
	}

	AddGenbankHook(boil.BeforeInsertHook, genbankBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	genbankBeforeInsertHooks = []GenbankHook{}

	AddGenbankHook(boil.AfterInsertHook, genbankAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	genbankAfterInsertHooks = []GenbankHook{}

	AddGenbankHook(boil.AfterSelectHook, genbankAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	genbankAfterSelectHooks = []GenbankHook{}

	AddGenbankHook(boil.BeforeUpdateHook, genbankBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	genbankBeforeUpdateHooks = []GenbankHook{}

	AddGenbankHook(boil.AfterUpdateHook, genbankAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	genbankAfterUpdateHooks = []GenbankHook{}

	AddGenbankHook(boil.BeforeDeleteHook, genbankBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	genbankBeforeDeleteHooks = []GenbankHook{}

	AddGenbankHook(boil.AfterDeleteHook, genbankAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	genbankAfterDeleteHooks = []GenbankHook{}

	AddGenbankHook(boil.BeforeUpsertHook, genbankBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	genbankBeforeUpsertHooks = []GenbankHook{}

	AddGenbankHook(boil.AfterUpsertHook, genbankAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	genbankAfterUpsertHooks = []GenbankHook{}
}

func testGenbanksInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Genbank{}
	if err = randomize.Struct(seed, o, genbankDBTypes, true, genbankColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genbank struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Genbanks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGenbanksInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Genbank{}
	if err = randomize.Struct(seed, o, genbankDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Genbank struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(genbankColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Genbanks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGenbankToOneSeqhashUsingSeqhash(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Genbank
	var foreign Seqhash

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, genbankDBTypes, false, genbankColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genbank struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, seqhashDBTypes, false, seqhashColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Seqhash struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.SeqhashID = foreign.Seqhash
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Seqhash().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.Seqhash != foreign.Seqhash {
		t.Errorf("want: %v, got %v", foreign.Seqhash, check.Seqhash)
	}

	slice := GenbankSlice{&local}
	if err = local.L.LoadSeqhash(ctx, tx, false, (*[]*Genbank)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Seqhash == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Seqhash = nil
	if err = local.L.LoadSeqhash(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Seqhash == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testGenbankToOneSetOpSeqhashUsingSeqhash(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Genbank
	var b, c Seqhash

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, genbankDBTypes, false, strmangle.SetComplement(genbankPrimaryKeyColumns, genbankColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, seqhashDBTypes, false, strmangle.SetComplement(seqhashPrimaryKeyColumns, seqhashColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, seqhashDBTypes, false, strmangle.SetComplement(seqhashPrimaryKeyColumns, seqhashColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Seqhash{&b, &c} {
		err = a.SetSeqhash(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Seqhash != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Genbanks[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.SeqhashID != x.Seqhash {
			t.Error("foreign key was wrong value", a.SeqhashID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.SeqhashID))
		reflect.Indirect(reflect.ValueOf(&a.SeqhashID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.SeqhashID != x.Seqhash {
			t.Error("foreign key was wrong value", a.SeqhashID, x.Seqhash)
		}
	}
}

func testGenbanksReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Genbank{}
	if err = randomize.Struct(seed, o, genbankDBTypes, true, genbankColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genbank struct: %s", err)
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

func testGenbanksReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Genbank{}
	if err = randomize.Struct(seed, o, genbankDBTypes, true, genbankColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genbank struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := GenbankSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testGenbanksSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Genbank{}
	if err = randomize.Struct(seed, o, genbankDBTypes, true, genbankColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genbank struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Genbanks().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	genbankDBTypes = map[string]string{`Accession`: `text`, `SeqhashID`: `text`, `BinaryJSON`: `jsonb`, `JSONHash`: `text`}
	_              = bytes.MinRead
)

func testGenbanksUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(genbankPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(genbankAllColumns) == len(genbankPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Genbank{}
	if err = randomize.Struct(seed, o, genbankDBTypes, true, genbankColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genbank struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Genbanks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, genbankDBTypes, true, genbankPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Genbank struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testGenbanksSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(genbankAllColumns) == len(genbankPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Genbank{}
	if err = randomize.Struct(seed, o, genbankDBTypes, true, genbankColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genbank struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Genbanks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, genbankDBTypes, true, genbankPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Genbank struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(genbankAllColumns, genbankPrimaryKeyColumns) {
		fields = genbankAllColumns
	} else {
		fields = strmangle.SetComplement(
			genbankAllColumns,
			genbankPrimaryKeyColumns,
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

	slice := GenbankSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testGenbanksUpsert(t *testing.T) {
	t.Parallel()

	if len(genbankAllColumns) == len(genbankPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Genbank{}
	if err = randomize.Struct(seed, &o, genbankDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Genbank struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Genbank: %s", err)
	}

	count, err := Genbanks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, genbankDBTypes, false, genbankPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Genbank struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Genbank: %s", err)
	}

	count, err = Genbanks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
