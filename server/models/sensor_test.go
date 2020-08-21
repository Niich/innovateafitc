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

func testSensors(t *testing.T) {
	t.Parallel()

	query := Sensors()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testSensorsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sensor{}
	if err = randomize.Struct(seed, o, sensorDBTypes, true, sensorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sensor struct: %s", err)
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

	count, err := Sensors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSensorsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sensor{}
	if err = randomize.Struct(seed, o, sensorDBTypes, true, sensorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sensor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Sensors().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Sensors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSensorsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sensor{}
	if err = randomize.Struct(seed, o, sensorDBTypes, true, sensorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sensor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SensorSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Sensors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSensorsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sensor{}
	if err = randomize.Struct(seed, o, sensorDBTypes, true, sensorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sensor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := SensorExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Sensor exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SensorExists to return true, but got false.")
	}
}

func testSensorsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sensor{}
	if err = randomize.Struct(seed, o, sensorDBTypes, true, sensorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sensor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	sensorFound, err := FindSensor(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if sensorFound == nil {
		t.Error("want a record, got nil")
	}
}

func testSensorsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sensor{}
	if err = randomize.Struct(seed, o, sensorDBTypes, true, sensorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sensor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Sensors().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testSensorsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sensor{}
	if err = randomize.Struct(seed, o, sensorDBTypes, true, sensorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sensor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Sensors().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSensorsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	sensorOne := &Sensor{}
	sensorTwo := &Sensor{}
	if err = randomize.Struct(seed, sensorOne, sensorDBTypes, false, sensorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sensor struct: %s", err)
	}
	if err = randomize.Struct(seed, sensorTwo, sensorDBTypes, false, sensorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sensor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = sensorOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = sensorTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Sensors().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSensorsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	sensorOne := &Sensor{}
	sensorTwo := &Sensor{}
	if err = randomize.Struct(seed, sensorOne, sensorDBTypes, false, sensorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sensor struct: %s", err)
	}
	if err = randomize.Struct(seed, sensorTwo, sensorDBTypes, false, sensorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sensor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = sensorOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = sensorTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Sensors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func sensorBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Sensor) error {
	*o = Sensor{}
	return nil
}

func sensorAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Sensor) error {
	*o = Sensor{}
	return nil
}

func sensorAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Sensor) error {
	*o = Sensor{}
	return nil
}

func sensorBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Sensor) error {
	*o = Sensor{}
	return nil
}

func sensorAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Sensor) error {
	*o = Sensor{}
	return nil
}

func sensorBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Sensor) error {
	*o = Sensor{}
	return nil
}

func sensorAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Sensor) error {
	*o = Sensor{}
	return nil
}

func sensorBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Sensor) error {
	*o = Sensor{}
	return nil
}

func sensorAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Sensor) error {
	*o = Sensor{}
	return nil
}

func testSensorsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Sensor{}
	o := &Sensor{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, sensorDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Sensor object: %s", err)
	}

	AddSensorHook(boil.BeforeInsertHook, sensorBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	sensorBeforeInsertHooks = []SensorHook{}

	AddSensorHook(boil.AfterInsertHook, sensorAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	sensorAfterInsertHooks = []SensorHook{}

	AddSensorHook(boil.AfterSelectHook, sensorAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	sensorAfterSelectHooks = []SensorHook{}

	AddSensorHook(boil.BeforeUpdateHook, sensorBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	sensorBeforeUpdateHooks = []SensorHook{}

	AddSensorHook(boil.AfterUpdateHook, sensorAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	sensorAfterUpdateHooks = []SensorHook{}

	AddSensorHook(boil.BeforeDeleteHook, sensorBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	sensorBeforeDeleteHooks = []SensorHook{}

	AddSensorHook(boil.AfterDeleteHook, sensorAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	sensorAfterDeleteHooks = []SensorHook{}

	AddSensorHook(boil.BeforeUpsertHook, sensorBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	sensorBeforeUpsertHooks = []SensorHook{}

	AddSensorHook(boil.AfterUpsertHook, sensorAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	sensorAfterUpsertHooks = []SensorHook{}
}

func testSensorsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sensor{}
	if err = randomize.Struct(seed, o, sensorDBTypes, true, sensorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sensor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Sensors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSensorsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sensor{}
	if err = randomize.Struct(seed, o, sensorDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Sensor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(sensorColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Sensors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSensorToOneTrackedDeviceUsingTrackedDevice(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Sensor
	var foreign TrackedDevice

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, sensorDBTypes, true, sensorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sensor struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, trackedDeviceDBTypes, false, trackedDeviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TrackedDevice struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.TrackedDeviceID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.TrackedDevice().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := SensorSlice{&local}
	if err = local.L.LoadTrackedDevice(ctx, tx, false, (*[]*Sensor)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.TrackedDevice == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.TrackedDevice = nil
	if err = local.L.LoadTrackedDevice(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.TrackedDevice == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testSensorToOneSetOpTrackedDeviceUsingTrackedDevice(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Sensor
	var b, c TrackedDevice

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, sensorDBTypes, false, strmangle.SetComplement(sensorPrimaryKeyColumns, sensorColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, trackedDeviceDBTypes, false, strmangle.SetComplement(trackedDevicePrimaryKeyColumns, trackedDeviceColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, trackedDeviceDBTypes, false, strmangle.SetComplement(trackedDevicePrimaryKeyColumns, trackedDeviceColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*TrackedDevice{&b, &c} {
		err = a.SetTrackedDevice(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.TrackedDevice != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Sensors[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.TrackedDeviceID, x.ID) {
			t.Error("foreign key was wrong value", a.TrackedDeviceID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.TrackedDeviceID))
		reflect.Indirect(reflect.ValueOf(&a.TrackedDeviceID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.TrackedDeviceID, x.ID) {
			t.Error("foreign key was wrong value", a.TrackedDeviceID, x.ID)
		}
	}
}

func testSensorToOneRemoveOpTrackedDeviceUsingTrackedDevice(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Sensor
	var b TrackedDevice

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, sensorDBTypes, false, strmangle.SetComplement(sensorPrimaryKeyColumns, sensorColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, trackedDeviceDBTypes, false, strmangle.SetComplement(trackedDevicePrimaryKeyColumns, trackedDeviceColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetTrackedDevice(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveTrackedDevice(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.TrackedDevice().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.TrackedDevice != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.TrackedDeviceID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.Sensors) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testSensorsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sensor{}
	if err = randomize.Struct(seed, o, sensorDBTypes, true, sensorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sensor struct: %s", err)
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

func testSensorsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sensor{}
	if err = randomize.Struct(seed, o, sensorDBTypes, true, sensorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sensor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SensorSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSensorsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sensor{}
	if err = randomize.Struct(seed, o, sensorDBTypes, true, sensorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sensor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Sensors().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	sensorDBTypes = map[string]string{`ID`: `character varying`, `TrackedDeviceID`: `bigint`, `BatStatus`: `integer`}
	_             = bytes.MinRead
)

func testSensorsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(sensorPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(sensorAllColumns) == len(sensorPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Sensor{}
	if err = randomize.Struct(seed, o, sensorDBTypes, true, sensorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sensor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Sensors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, sensorDBTypes, true, sensorPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Sensor struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testSensorsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(sensorAllColumns) == len(sensorPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Sensor{}
	if err = randomize.Struct(seed, o, sensorDBTypes, true, sensorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sensor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Sensors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, sensorDBTypes, true, sensorPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Sensor struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(sensorAllColumns, sensorPrimaryKeyColumns) {
		fields = sensorAllColumns
	} else {
		fields = strmangle.SetComplement(
			sensorAllColumns,
			sensorPrimaryKeyColumns,
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

	slice := SensorSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testSensorsUpsert(t *testing.T) {
	t.Parallel()

	if len(sensorAllColumns) == len(sensorPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Sensor{}
	if err = randomize.Struct(seed, &o, sensorDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Sensor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Sensor: %s", err)
	}

	count, err := Sensors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, sensorDBTypes, false, sensorPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Sensor struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Sensor: %s", err)
	}

	count, err = Sensors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}