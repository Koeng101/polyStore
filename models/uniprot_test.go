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

func testUniprots(t *testing.T) {
	t.Parallel()

	query := Uniprots()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testUniprotsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Uniprot{}
	if err = randomize.Struct(seed, o, uniprotDBTypes, true, uniprotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Uniprot struct: %s", err)
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

	count, err := Uniprots().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUniprotsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Uniprot{}
	if err = randomize.Struct(seed, o, uniprotDBTypes, true, uniprotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Uniprot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Uniprots().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Uniprots().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUniprotsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Uniprot{}
	if err = randomize.Struct(seed, o, uniprotDBTypes, true, uniprotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Uniprot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UniprotSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Uniprots().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUniprotsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Uniprot{}
	if err = randomize.Struct(seed, o, uniprotDBTypes, true, uniprotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Uniprot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := UniprotExists(ctx, tx, o.Accession)
	if err != nil {
		t.Errorf("Unable to check if Uniprot exists: %s", err)
	}
	if !e {
		t.Errorf("Expected UniprotExists to return true, but got false.")
	}
}

func testUniprotsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Uniprot{}
	if err = randomize.Struct(seed, o, uniprotDBTypes, true, uniprotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Uniprot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	uniprotFound, err := FindUniprot(ctx, tx, o.Accession)
	if err != nil {
		t.Error(err)
	}

	if uniprotFound == nil {
		t.Error("want a record, got nil")
	}
}

func testUniprotsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Uniprot{}
	if err = randomize.Struct(seed, o, uniprotDBTypes, true, uniprotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Uniprot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Uniprots().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testUniprotsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Uniprot{}
	if err = randomize.Struct(seed, o, uniprotDBTypes, true, uniprotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Uniprot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Uniprots().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testUniprotsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	uniprotOne := &Uniprot{}
	uniprotTwo := &Uniprot{}
	if err = randomize.Struct(seed, uniprotOne, uniprotDBTypes, false, uniprotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Uniprot struct: %s", err)
	}
	if err = randomize.Struct(seed, uniprotTwo, uniprotDBTypes, false, uniprotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Uniprot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = uniprotOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = uniprotTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Uniprots().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testUniprotsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	uniprotOne := &Uniprot{}
	uniprotTwo := &Uniprot{}
	if err = randomize.Struct(seed, uniprotOne, uniprotDBTypes, false, uniprotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Uniprot struct: %s", err)
	}
	if err = randomize.Struct(seed, uniprotTwo, uniprotDBTypes, false, uniprotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Uniprot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = uniprotOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = uniprotTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Uniprots().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func uniprotBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Uniprot) error {
	*o = Uniprot{}
	return nil
}

func uniprotAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Uniprot) error {
	*o = Uniprot{}
	return nil
}

func uniprotAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Uniprot) error {
	*o = Uniprot{}
	return nil
}

func uniprotBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Uniprot) error {
	*o = Uniprot{}
	return nil
}

func uniprotAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Uniprot) error {
	*o = Uniprot{}
	return nil
}

func uniprotBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Uniprot) error {
	*o = Uniprot{}
	return nil
}

func uniprotAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Uniprot) error {
	*o = Uniprot{}
	return nil
}

func uniprotBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Uniprot) error {
	*o = Uniprot{}
	return nil
}

func uniprotAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Uniprot) error {
	*o = Uniprot{}
	return nil
}

func testUniprotsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Uniprot{}
	o := &Uniprot{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, uniprotDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Uniprot object: %s", err)
	}

	AddUniprotHook(boil.BeforeInsertHook, uniprotBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	uniprotBeforeInsertHooks = []UniprotHook{}

	AddUniprotHook(boil.AfterInsertHook, uniprotAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	uniprotAfterInsertHooks = []UniprotHook{}

	AddUniprotHook(boil.AfterSelectHook, uniprotAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	uniprotAfterSelectHooks = []UniprotHook{}

	AddUniprotHook(boil.BeforeUpdateHook, uniprotBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	uniprotBeforeUpdateHooks = []UniprotHook{}

	AddUniprotHook(boil.AfterUpdateHook, uniprotAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	uniprotAfterUpdateHooks = []UniprotHook{}

	AddUniprotHook(boil.BeforeDeleteHook, uniprotBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	uniprotBeforeDeleteHooks = []UniprotHook{}

	AddUniprotHook(boil.AfterDeleteHook, uniprotAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	uniprotAfterDeleteHooks = []UniprotHook{}

	AddUniprotHook(boil.BeforeUpsertHook, uniprotBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	uniprotBeforeUpsertHooks = []UniprotHook{}

	AddUniprotHook(boil.AfterUpsertHook, uniprotAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	uniprotAfterUpsertHooks = []UniprotHook{}
}

func testUniprotsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Uniprot{}
	if err = randomize.Struct(seed, o, uniprotDBTypes, true, uniprotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Uniprot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Uniprots().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUniprotsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Uniprot{}
	if err = randomize.Struct(seed, o, uniprotDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Uniprot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(uniprotColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Uniprots().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUniprotToOneSeqhashUsingSeqhash(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Uniprot
	var foreign Seqhash

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, uniprotDBTypes, false, uniprotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Uniprot struct: %s", err)
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

	slice := UniprotSlice{&local}
	if err = local.L.LoadSeqhash(ctx, tx, false, (*[]*Uniprot)(&slice), nil); err != nil {
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

func testUniprotToOneSetOpSeqhashUsingSeqhash(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Uniprot
	var b, c Seqhash

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, uniprotDBTypes, false, strmangle.SetComplement(uniprotPrimaryKeyColumns, uniprotColumnsWithoutDefault)...); err != nil {
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

		if x.R.Uniprots[0] != &a {
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

func testUniprotsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Uniprot{}
	if err = randomize.Struct(seed, o, uniprotDBTypes, true, uniprotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Uniprot struct: %s", err)
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

func testUniprotsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Uniprot{}
	if err = randomize.Struct(seed, o, uniprotDBTypes, true, uniprotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Uniprot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UniprotSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testUniprotsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Uniprot{}
	if err = randomize.Struct(seed, o, uniprotDBTypes, true, uniprotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Uniprot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Uniprots().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	uniprotDBTypes = map[string]string{`Accession`: `text`, `SeqhashID`: `text`, `Database`: `text`}
	_              = bytes.MinRead
)

func testUniprotsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(uniprotPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(uniprotAllColumns) == len(uniprotPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Uniprot{}
	if err = randomize.Struct(seed, o, uniprotDBTypes, true, uniprotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Uniprot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Uniprots().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, uniprotDBTypes, true, uniprotPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Uniprot struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testUniprotsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(uniprotAllColumns) == len(uniprotPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Uniprot{}
	if err = randomize.Struct(seed, o, uniprotDBTypes, true, uniprotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Uniprot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Uniprots().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, uniprotDBTypes, true, uniprotPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Uniprot struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(uniprotAllColumns, uniprotPrimaryKeyColumns) {
		fields = uniprotAllColumns
	} else {
		fields = strmangle.SetComplement(
			uniprotAllColumns,
			uniprotPrimaryKeyColumns,
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

	slice := UniprotSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testUniprotsUpsert(t *testing.T) {
	t.Parallel()

	if len(uniprotAllColumns) == len(uniprotPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Uniprot{}
	if err = randomize.Struct(seed, &o, uniprotDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Uniprot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Uniprot: %s", err)
	}

	count, err := Uniprots().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, uniprotDBTypes, false, uniprotPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Uniprot struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Uniprot: %s", err)
	}

	count, err = Uniprots().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
