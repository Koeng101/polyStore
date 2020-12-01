// Code generated by SQLBoiler 4.3.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Basemodels", testBasemodels)
	t.Run("Genbanks", testGenbanks)
	t.Run("Seqhashes", testSeqhashes)
	t.Run("Seqhashlinks", testSeqhashlinks)
	t.Run("Uniprots", testUniprots)
}

func TestDelete(t *testing.T) {
	t.Run("Basemodels", testBasemodelsDelete)
	t.Run("Genbanks", testGenbanksDelete)
	t.Run("Seqhashes", testSeqhashesDelete)
	t.Run("Seqhashlinks", testSeqhashlinksDelete)
	t.Run("Uniprots", testUniprotsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Basemodels", testBasemodelsQueryDeleteAll)
	t.Run("Genbanks", testGenbanksQueryDeleteAll)
	t.Run("Seqhashes", testSeqhashesQueryDeleteAll)
	t.Run("Seqhashlinks", testSeqhashlinksQueryDeleteAll)
	t.Run("Uniprots", testUniprotsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Basemodels", testBasemodelsSliceDeleteAll)
	t.Run("Genbanks", testGenbanksSliceDeleteAll)
	t.Run("Seqhashes", testSeqhashesSliceDeleteAll)
	t.Run("Seqhashlinks", testSeqhashlinksSliceDeleteAll)
	t.Run("Uniprots", testUniprotsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Basemodels", testBasemodelsExists)
	t.Run("Genbanks", testGenbanksExists)
	t.Run("Seqhashes", testSeqhashesExists)
	t.Run("Seqhashlinks", testSeqhashlinksExists)
	t.Run("Uniprots", testUniprotsExists)
}

func TestFind(t *testing.T) {
	t.Run("Basemodels", testBasemodelsFind)
	t.Run("Genbanks", testGenbanksFind)
	t.Run("Seqhashes", testSeqhashesFind)
	t.Run("Seqhashlinks", testSeqhashlinksFind)
	t.Run("Uniprots", testUniprotsFind)
}

func TestBind(t *testing.T) {
	t.Run("Basemodels", testBasemodelsBind)
	t.Run("Genbanks", testGenbanksBind)
	t.Run("Seqhashes", testSeqhashesBind)
	t.Run("Seqhashlinks", testSeqhashlinksBind)
	t.Run("Uniprots", testUniprotsBind)
}

func TestOne(t *testing.T) {
	t.Run("Basemodels", testBasemodelsOne)
	t.Run("Genbanks", testGenbanksOne)
	t.Run("Seqhashes", testSeqhashesOne)
	t.Run("Seqhashlinks", testSeqhashlinksOne)
	t.Run("Uniprots", testUniprotsOne)
}

func TestAll(t *testing.T) {
	t.Run("Basemodels", testBasemodelsAll)
	t.Run("Genbanks", testGenbanksAll)
	t.Run("Seqhashes", testSeqhashesAll)
	t.Run("Seqhashlinks", testSeqhashlinksAll)
	t.Run("Uniprots", testUniprotsAll)
}

func TestCount(t *testing.T) {
	t.Run("Basemodels", testBasemodelsCount)
	t.Run("Genbanks", testGenbanksCount)
	t.Run("Seqhashes", testSeqhashesCount)
	t.Run("Seqhashlinks", testSeqhashlinksCount)
	t.Run("Uniprots", testUniprotsCount)
}

func TestHooks(t *testing.T) {
	t.Run("Basemodels", testBasemodelsHooks)
	t.Run("Genbanks", testGenbanksHooks)
	t.Run("Seqhashes", testSeqhashesHooks)
	t.Run("Seqhashlinks", testSeqhashlinksHooks)
	t.Run("Uniprots", testUniprotsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Basemodels", testBasemodelsInsert)
	t.Run("Basemodels", testBasemodelsInsertWhitelist)
	t.Run("Genbanks", testGenbanksInsert)
	t.Run("Genbanks", testGenbanksInsertWhitelist)
	t.Run("Seqhashes", testSeqhashesInsert)
	t.Run("Seqhashes", testSeqhashesInsertWhitelist)
	t.Run("Seqhashlinks", testSeqhashlinksInsert)
	t.Run("Seqhashlinks", testSeqhashlinksInsertWhitelist)
	t.Run("Uniprots", testUniprotsInsert)
	t.Run("Uniprots", testUniprotsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("GenbankToSeqhashUsingSeqhash", testGenbankToOneSeqhashUsingSeqhash)
	t.Run("SeqhashlinkToSeqhashUsingChildSeqhash", testSeqhashlinkToOneSeqhashUsingChildSeqhash)
	t.Run("SeqhashlinkToSeqhashUsingParentSeqhash", testSeqhashlinkToOneSeqhashUsingParentSeqhash)
	t.Run("UniprotToSeqhashUsingSeqhash", testUniprotToOneSeqhashUsingSeqhash)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("SeqhashToGenbanks", testSeqhashToManyGenbanks)
	t.Run("SeqhashToChildSeqhashSeqhashlinks", testSeqhashToManyChildSeqhashSeqhashlinks)
	t.Run("SeqhashToParentSeqhashSeqhashlinks", testSeqhashToManyParentSeqhashSeqhashlinks)
	t.Run("SeqhashToUniprots", testSeqhashToManyUniprots)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("GenbankToSeqhashUsingGenbanks", testGenbankToOneSetOpSeqhashUsingSeqhash)
	t.Run("SeqhashlinkToSeqhashUsingChildSeqhashSeqhashlinks", testSeqhashlinkToOneSetOpSeqhashUsingChildSeqhash)
	t.Run("SeqhashlinkToSeqhashUsingParentSeqhashSeqhashlinks", testSeqhashlinkToOneSetOpSeqhashUsingParentSeqhash)
	t.Run("UniprotToSeqhashUsingUniprots", testUniprotToOneSetOpSeqhashUsingSeqhash)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("SeqhashToGenbanks", testSeqhashToManyAddOpGenbanks)
	t.Run("SeqhashToChildSeqhashSeqhashlinks", testSeqhashToManyAddOpChildSeqhashSeqhashlinks)
	t.Run("SeqhashToParentSeqhashSeqhashlinks", testSeqhashToManyAddOpParentSeqhashSeqhashlinks)
	t.Run("SeqhashToUniprots", testSeqhashToManyAddOpUniprots)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Basemodels", testBasemodelsReload)
	t.Run("Genbanks", testGenbanksReload)
	t.Run("Seqhashes", testSeqhashesReload)
	t.Run("Seqhashlinks", testSeqhashlinksReload)
	t.Run("Uniprots", testUniprotsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Basemodels", testBasemodelsReloadAll)
	t.Run("Genbanks", testGenbanksReloadAll)
	t.Run("Seqhashes", testSeqhashesReloadAll)
	t.Run("Seqhashlinks", testSeqhashlinksReloadAll)
	t.Run("Uniprots", testUniprotsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Basemodels", testBasemodelsSelect)
	t.Run("Genbanks", testGenbanksSelect)
	t.Run("Seqhashes", testSeqhashesSelect)
	t.Run("Seqhashlinks", testSeqhashlinksSelect)
	t.Run("Uniprots", testUniprotsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Basemodels", testBasemodelsUpdate)
	t.Run("Genbanks", testGenbanksUpdate)
	t.Run("Seqhashes", testSeqhashesUpdate)
	t.Run("Seqhashlinks", testSeqhashlinksUpdate)
	t.Run("Uniprots", testUniprotsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Basemodels", testBasemodelsSliceUpdateAll)
	t.Run("Genbanks", testGenbanksSliceUpdateAll)
	t.Run("Seqhashes", testSeqhashesSliceUpdateAll)
	t.Run("Seqhashlinks", testSeqhashlinksSliceUpdateAll)
	t.Run("Uniprots", testUniprotsSliceUpdateAll)
}
