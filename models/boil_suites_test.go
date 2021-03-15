// Code generated by SQLBoiler 4.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	t.Run("Authors", testAuthors)
}

func TestDelete(t *testing.T) {
	t.Run("Authors", testAuthorsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Authors", testAuthorsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Authors", testAuthorsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Authors", testAuthorsExists)
}

func TestFind(t *testing.T) {
	t.Run("Authors", testAuthorsFind)
}

func TestBind(t *testing.T) {
	t.Run("Authors", testAuthorsBind)
}

func TestOne(t *testing.T) {
	t.Run("Authors", testAuthorsOne)
}

func TestAll(t *testing.T) {
	t.Run("Authors", testAuthorsAll)
}

func TestCount(t *testing.T) {
	t.Run("Authors", testAuthorsCount)
}

func TestHooks(t *testing.T) {
	t.Run("Authors", testAuthorsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Authors", testAuthorsInsert)
	t.Run("Authors", testAuthorsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {}

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
func TestToManyAdd(t *testing.T) {}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Authors", testAuthorsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Authors", testAuthorsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Authors", testAuthorsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Authors", testAuthorsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Authors", testAuthorsSliceUpdateAll)
}
