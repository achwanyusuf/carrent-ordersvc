// Code generated by SQLBoiler 4.16.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package psqlmodel

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

func testSchemaMigrations(t *testing.T) {
	t.Parallel()

	query := SchemaMigrations()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testSchemaMigrationsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SchemaMigration{}
	if err = randomize.Struct(seed, o, schemaMigrationDBTypes, true, schemaMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SchemaMigration struct: %s", err)
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

	count, err := SchemaMigrations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSchemaMigrationsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SchemaMigration{}
	if err = randomize.Struct(seed, o, schemaMigrationDBTypes, true, schemaMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SchemaMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := SchemaMigrations().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := SchemaMigrations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSchemaMigrationsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SchemaMigration{}
	if err = randomize.Struct(seed, o, schemaMigrationDBTypes, true, schemaMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SchemaMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SchemaMigrationSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := SchemaMigrations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSchemaMigrationsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SchemaMigration{}
	if err = randomize.Struct(seed, o, schemaMigrationDBTypes, true, schemaMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SchemaMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := SchemaMigrationExists(ctx, tx, o.Version)
	if err != nil {
		t.Errorf("Unable to check if SchemaMigration exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SchemaMigrationExists to return true, but got false.")
	}
}

func testSchemaMigrationsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SchemaMigration{}
	if err = randomize.Struct(seed, o, schemaMigrationDBTypes, true, schemaMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SchemaMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	schemaMigrationFound, err := FindSchemaMigration(ctx, tx, o.Version)
	if err != nil {
		t.Error(err)
	}

	if schemaMigrationFound == nil {
		t.Error("want a record, got nil")
	}
}

func testSchemaMigrationsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SchemaMigration{}
	if err = randomize.Struct(seed, o, schemaMigrationDBTypes, true, schemaMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SchemaMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = SchemaMigrations().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testSchemaMigrationsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SchemaMigration{}
	if err = randomize.Struct(seed, o, schemaMigrationDBTypes, true, schemaMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SchemaMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := SchemaMigrations().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSchemaMigrationsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	schemaMigrationOne := &SchemaMigration{}
	schemaMigrationTwo := &SchemaMigration{}
	if err = randomize.Struct(seed, schemaMigrationOne, schemaMigrationDBTypes, false, schemaMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SchemaMigration struct: %s", err)
	}
	if err = randomize.Struct(seed, schemaMigrationTwo, schemaMigrationDBTypes, false, schemaMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SchemaMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = schemaMigrationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = schemaMigrationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := SchemaMigrations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSchemaMigrationsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	schemaMigrationOne := &SchemaMigration{}
	schemaMigrationTwo := &SchemaMigration{}
	if err = randomize.Struct(seed, schemaMigrationOne, schemaMigrationDBTypes, false, schemaMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SchemaMigration struct: %s", err)
	}
	if err = randomize.Struct(seed, schemaMigrationTwo, schemaMigrationDBTypes, false, schemaMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SchemaMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = schemaMigrationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = schemaMigrationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SchemaMigrations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func schemaMigrationBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *SchemaMigration) error {
	*o = SchemaMigration{}
	return nil
}

func schemaMigrationAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *SchemaMigration) error {
	*o = SchemaMigration{}
	return nil
}

func schemaMigrationAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *SchemaMigration) error {
	*o = SchemaMigration{}
	return nil
}

func schemaMigrationBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *SchemaMigration) error {
	*o = SchemaMigration{}
	return nil
}

func schemaMigrationAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *SchemaMigration) error {
	*o = SchemaMigration{}
	return nil
}

func schemaMigrationBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *SchemaMigration) error {
	*o = SchemaMigration{}
	return nil
}

func schemaMigrationAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *SchemaMigration) error {
	*o = SchemaMigration{}
	return nil
}

func schemaMigrationBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *SchemaMigration) error {
	*o = SchemaMigration{}
	return nil
}

func schemaMigrationAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *SchemaMigration) error {
	*o = SchemaMigration{}
	return nil
}

func testSchemaMigrationsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &SchemaMigration{}
	o := &SchemaMigration{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, schemaMigrationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize SchemaMigration object: %s", err)
	}

	AddSchemaMigrationHook(boil.BeforeInsertHook, schemaMigrationBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	schemaMigrationBeforeInsertHooks = []SchemaMigrationHook{}

	AddSchemaMigrationHook(boil.AfterInsertHook, schemaMigrationAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	schemaMigrationAfterInsertHooks = []SchemaMigrationHook{}

	AddSchemaMigrationHook(boil.AfterSelectHook, schemaMigrationAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	schemaMigrationAfterSelectHooks = []SchemaMigrationHook{}

	AddSchemaMigrationHook(boil.BeforeUpdateHook, schemaMigrationBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	schemaMigrationBeforeUpdateHooks = []SchemaMigrationHook{}

	AddSchemaMigrationHook(boil.AfterUpdateHook, schemaMigrationAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	schemaMigrationAfterUpdateHooks = []SchemaMigrationHook{}

	AddSchemaMigrationHook(boil.BeforeDeleteHook, schemaMigrationBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	schemaMigrationBeforeDeleteHooks = []SchemaMigrationHook{}

	AddSchemaMigrationHook(boil.AfterDeleteHook, schemaMigrationAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	schemaMigrationAfterDeleteHooks = []SchemaMigrationHook{}

	AddSchemaMigrationHook(boil.BeforeUpsertHook, schemaMigrationBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	schemaMigrationBeforeUpsertHooks = []SchemaMigrationHook{}

	AddSchemaMigrationHook(boil.AfterUpsertHook, schemaMigrationAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	schemaMigrationAfterUpsertHooks = []SchemaMigrationHook{}
}

func testSchemaMigrationsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SchemaMigration{}
	if err = randomize.Struct(seed, o, schemaMigrationDBTypes, true, schemaMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SchemaMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SchemaMigrations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSchemaMigrationsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SchemaMigration{}
	if err = randomize.Struct(seed, o, schemaMigrationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize SchemaMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(schemaMigrationColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := SchemaMigrations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSchemaMigrationsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SchemaMigration{}
	if err = randomize.Struct(seed, o, schemaMigrationDBTypes, true, schemaMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SchemaMigration struct: %s", err)
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

func testSchemaMigrationsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SchemaMigration{}
	if err = randomize.Struct(seed, o, schemaMigrationDBTypes, true, schemaMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SchemaMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SchemaMigrationSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSchemaMigrationsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SchemaMigration{}
	if err = randomize.Struct(seed, o, schemaMigrationDBTypes, true, schemaMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SchemaMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := SchemaMigrations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	schemaMigrationDBTypes = map[string]string{`Version`: `bigint`, `Dirty`: `boolean`}
	_                      = bytes.MinRead
)

func testSchemaMigrationsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(schemaMigrationPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(schemaMigrationAllColumns) == len(schemaMigrationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &SchemaMigration{}
	if err = randomize.Struct(seed, o, schemaMigrationDBTypes, true, schemaMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SchemaMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SchemaMigrations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, schemaMigrationDBTypes, true, schemaMigrationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SchemaMigration struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testSchemaMigrationsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(schemaMigrationAllColumns) == len(schemaMigrationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &SchemaMigration{}
	if err = randomize.Struct(seed, o, schemaMigrationDBTypes, true, schemaMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SchemaMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SchemaMigrations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, schemaMigrationDBTypes, true, schemaMigrationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SchemaMigration struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(schemaMigrationAllColumns, schemaMigrationPrimaryKeyColumns) {
		fields = schemaMigrationAllColumns
	} else {
		fields = strmangle.SetComplement(
			schemaMigrationAllColumns,
			schemaMigrationPrimaryKeyColumns,
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

	slice := SchemaMigrationSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testSchemaMigrationsUpsert(t *testing.T) {
	t.Parallel()

	if len(schemaMigrationAllColumns) == len(schemaMigrationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := SchemaMigration{}
	if err = randomize.Struct(seed, &o, schemaMigrationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize SchemaMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert SchemaMigration: %s", err)
	}

	count, err := SchemaMigrations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, schemaMigrationDBTypes, false, schemaMigrationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SchemaMigration struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert SchemaMigration: %s", err)
	}

	count, err = SchemaMigrations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
