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

func testSeqhashlinks(t *testing.T) {
	t.Parallel()

	query := Seqhashlinks()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testSeqhashlinksDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Seqhashlink{}
	if err = randomize.Struct(seed, o, seqhashlinkDBTypes, true, seqhashlinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Seqhashlink struct: %s", err)
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

	count, err := Seqhashlinks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSeqhashlinksQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Seqhashlink{}
	if err = randomize.Struct(seed, o, seqhashlinkDBTypes, true, seqhashlinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Seqhashlink struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Seqhashlinks().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Seqhashlinks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSeqhashlinksSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Seqhashlink{}
	if err = randomize.Struct(seed, o, seqhashlinkDBTypes, true, seqhashlinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Seqhashlink struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SeqhashlinkSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Seqhashlinks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSeqhashlinksExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Seqhashlink{}
	if err = randomize.Struct(seed, o, seqhashlinkDBTypes, true, seqhashlinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Seqhashlink struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := SeqhashlinkExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Seqhashlink exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SeqhashlinkExists to return true, but got false.")
	}
}

func testSeqhashlinksFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Seqhashlink{}
	if err = randomize.Struct(seed, o, seqhashlinkDBTypes, true, seqhashlinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Seqhashlink struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	seqhashlinkFound, err := FindSeqhashlink(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if seqhashlinkFound == nil {
		t.Error("want a record, got nil")
	}
}

func testSeqhashlinksBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Seqhashlink{}
	if err = randomize.Struct(seed, o, seqhashlinkDBTypes, true, seqhashlinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Seqhashlink struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Seqhashlinks().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testSeqhashlinksOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Seqhashlink{}
	if err = randomize.Struct(seed, o, seqhashlinkDBTypes, true, seqhashlinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Seqhashlink struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Seqhashlinks().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSeqhashlinksAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	seqhashlinkOne := &Seqhashlink{}
	seqhashlinkTwo := &Seqhashlink{}
	if err = randomize.Struct(seed, seqhashlinkOne, seqhashlinkDBTypes, false, seqhashlinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Seqhashlink struct: %s", err)
	}
	if err = randomize.Struct(seed, seqhashlinkTwo, seqhashlinkDBTypes, false, seqhashlinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Seqhashlink struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = seqhashlinkOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = seqhashlinkTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Seqhashlinks().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSeqhashlinksCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	seqhashlinkOne := &Seqhashlink{}
	seqhashlinkTwo := &Seqhashlink{}
	if err = randomize.Struct(seed, seqhashlinkOne, seqhashlinkDBTypes, false, seqhashlinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Seqhashlink struct: %s", err)
	}
	if err = randomize.Struct(seed, seqhashlinkTwo, seqhashlinkDBTypes, false, seqhashlinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Seqhashlink struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = seqhashlinkOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = seqhashlinkTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Seqhashlinks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func seqhashlinkBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Seqhashlink) error {
	*o = Seqhashlink{}
	return nil
}

func seqhashlinkAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Seqhashlink) error {
	*o = Seqhashlink{}
	return nil
}

func seqhashlinkAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Seqhashlink) error {
	*o = Seqhashlink{}
	return nil
}

func seqhashlinkBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Seqhashlink) error {
	*o = Seqhashlink{}
	return nil
}

func seqhashlinkAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Seqhashlink) error {
	*o = Seqhashlink{}
	return nil
}

func seqhashlinkBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Seqhashlink) error {
	*o = Seqhashlink{}
	return nil
}

func seqhashlinkAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Seqhashlink) error {
	*o = Seqhashlink{}
	return nil
}

func seqhashlinkBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Seqhashlink) error {
	*o = Seqhashlink{}
	return nil
}

func seqhashlinkAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Seqhashlink) error {
	*o = Seqhashlink{}
	return nil
}

func testSeqhashlinksHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Seqhashlink{}
	o := &Seqhashlink{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, seqhashlinkDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Seqhashlink object: %s", err)
	}

	AddSeqhashlinkHook(boil.BeforeInsertHook, seqhashlinkBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	seqhashlinkBeforeInsertHooks = []SeqhashlinkHook{}

	AddSeqhashlinkHook(boil.AfterInsertHook, seqhashlinkAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	seqhashlinkAfterInsertHooks = []SeqhashlinkHook{}

	AddSeqhashlinkHook(boil.AfterSelectHook, seqhashlinkAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	seqhashlinkAfterSelectHooks = []SeqhashlinkHook{}

	AddSeqhashlinkHook(boil.BeforeUpdateHook, seqhashlinkBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	seqhashlinkBeforeUpdateHooks = []SeqhashlinkHook{}

	AddSeqhashlinkHook(boil.AfterUpdateHook, seqhashlinkAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	seqhashlinkAfterUpdateHooks = []SeqhashlinkHook{}

	AddSeqhashlinkHook(boil.BeforeDeleteHook, seqhashlinkBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	seqhashlinkBeforeDeleteHooks = []SeqhashlinkHook{}

	AddSeqhashlinkHook(boil.AfterDeleteHook, seqhashlinkAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	seqhashlinkAfterDeleteHooks = []SeqhashlinkHook{}

	AddSeqhashlinkHook(boil.BeforeUpsertHook, seqhashlinkBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	seqhashlinkBeforeUpsertHooks = []SeqhashlinkHook{}

	AddSeqhashlinkHook(boil.AfterUpsertHook, seqhashlinkAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	seqhashlinkAfterUpsertHooks = []SeqhashlinkHook{}
}

func testSeqhashlinksInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Seqhashlink{}
	if err = randomize.Struct(seed, o, seqhashlinkDBTypes, true, seqhashlinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Seqhashlink struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Seqhashlinks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSeqhashlinksInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Seqhashlink{}
	if err = randomize.Struct(seed, o, seqhashlinkDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Seqhashlink struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(seqhashlinkColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Seqhashlinks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSeqhashlinkToOneSeqhashUsingChildSeqhash(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Seqhashlink
	var foreign Seqhash

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, seqhashlinkDBTypes, false, seqhashlinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Seqhashlink struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, seqhashDBTypes, false, seqhashColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Seqhash struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ChildSeqhashID = foreign.Seqhash
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.ChildSeqhash().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.Seqhash != foreign.Seqhash {
		t.Errorf("want: %v, got %v", foreign.Seqhash, check.Seqhash)
	}

	slice := SeqhashlinkSlice{&local}
	if err = local.L.LoadChildSeqhash(ctx, tx, false, (*[]*Seqhashlink)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ChildSeqhash == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.ChildSeqhash = nil
	if err = local.L.LoadChildSeqhash(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ChildSeqhash == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testSeqhashlinkToOneSeqhashUsingParentSeqhash(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Seqhashlink
	var foreign Seqhash

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, seqhashlinkDBTypes, false, seqhashlinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Seqhashlink struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, seqhashDBTypes, false, seqhashColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Seqhash struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ParentSeqhashID = foreign.Seqhash
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.ParentSeqhash().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.Seqhash != foreign.Seqhash {
		t.Errorf("want: %v, got %v", foreign.Seqhash, check.Seqhash)
	}

	slice := SeqhashlinkSlice{&local}
	if err = local.L.LoadParentSeqhash(ctx, tx, false, (*[]*Seqhashlink)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ParentSeqhash == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.ParentSeqhash = nil
	if err = local.L.LoadParentSeqhash(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ParentSeqhash == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testSeqhashlinkToOneSetOpSeqhashUsingChildSeqhash(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Seqhashlink
	var b, c Seqhash

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, seqhashlinkDBTypes, false, strmangle.SetComplement(seqhashlinkPrimaryKeyColumns, seqhashlinkColumnsWithoutDefault)...); err != nil {
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
		err = a.SetChildSeqhash(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.ChildSeqhash != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ChildSeqhashSeqhashlinks[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ChildSeqhashID != x.Seqhash {
			t.Error("foreign key was wrong value", a.ChildSeqhashID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ChildSeqhashID))
		reflect.Indirect(reflect.ValueOf(&a.ChildSeqhashID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ChildSeqhashID != x.Seqhash {
			t.Error("foreign key was wrong value", a.ChildSeqhashID, x.Seqhash)
		}
	}
}
func testSeqhashlinkToOneSetOpSeqhashUsingParentSeqhash(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Seqhashlink
	var b, c Seqhash

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, seqhashlinkDBTypes, false, strmangle.SetComplement(seqhashlinkPrimaryKeyColumns, seqhashlinkColumnsWithoutDefault)...); err != nil {
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
		err = a.SetParentSeqhash(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.ParentSeqhash != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ParentSeqhashSeqhashlinks[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ParentSeqhashID != x.Seqhash {
			t.Error("foreign key was wrong value", a.ParentSeqhashID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ParentSeqhashID))
		reflect.Indirect(reflect.ValueOf(&a.ParentSeqhashID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ParentSeqhashID != x.Seqhash {
			t.Error("foreign key was wrong value", a.ParentSeqhashID, x.Seqhash)
		}
	}
}

func testSeqhashlinksReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Seqhashlink{}
	if err = randomize.Struct(seed, o, seqhashlinkDBTypes, true, seqhashlinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Seqhashlink struct: %s", err)
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

func testSeqhashlinksReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Seqhashlink{}
	if err = randomize.Struct(seed, o, seqhashlinkDBTypes, true, seqhashlinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Seqhashlink struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SeqhashlinkSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSeqhashlinksSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Seqhashlink{}
	if err = randomize.Struct(seed, o, seqhashlinkDBTypes, true, seqhashlinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Seqhashlink struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Seqhashlinks().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	seqhashlinkDBTypes = map[string]string{`ID`: `integer`, `ChildSeqhashID`: `text`, `ParentSeqhashID`: `text`}
	_                  = bytes.MinRead
)

func testSeqhashlinksUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(seqhashlinkPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(seqhashlinkAllColumns) == len(seqhashlinkPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Seqhashlink{}
	if err = randomize.Struct(seed, o, seqhashlinkDBTypes, true, seqhashlinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Seqhashlink struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Seqhashlinks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, seqhashlinkDBTypes, true, seqhashlinkPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Seqhashlink struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testSeqhashlinksSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(seqhashlinkAllColumns) == len(seqhashlinkPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Seqhashlink{}
	if err = randomize.Struct(seed, o, seqhashlinkDBTypes, true, seqhashlinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Seqhashlink struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Seqhashlinks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, seqhashlinkDBTypes, true, seqhashlinkPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Seqhashlink struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(seqhashlinkAllColumns, seqhashlinkPrimaryKeyColumns) {
		fields = seqhashlinkAllColumns
	} else {
		fields = strmangle.SetComplement(
			seqhashlinkAllColumns,
			seqhashlinkPrimaryKeyColumns,
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

	slice := SeqhashlinkSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testSeqhashlinksUpsert(t *testing.T) {
	t.Parallel()

	if len(seqhashlinkAllColumns) == len(seqhashlinkPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Seqhashlink{}
	if err = randomize.Struct(seed, &o, seqhashlinkDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Seqhashlink struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Seqhashlink: %s", err)
	}

	count, err := Seqhashlinks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, seqhashlinkDBTypes, false, seqhashlinkPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Seqhashlink struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Seqhashlink: %s", err)
	}

	count, err = Seqhashlinks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
