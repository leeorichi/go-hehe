// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package dbmodel

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// SchemaMigration is an object representing the database table.
type SchemaMigration struct {
	Version int64 `boil:"version" json:"version" toml:"version" yaml:"version"`
	Dirty   bool  `boil:"dirty" json:"dirty" toml:"dirty" yaml:"dirty"`

	R *schemaMigrationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L schemaMigrationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SchemaMigrationColumns = struct {
	Version string
	Dirty   string
}{
	Version: "version",
	Dirty:   "dirty",
}

var SchemaMigrationTableColumns = struct {
	Version string
	Dirty   string
}{
	Version: "schema_migrations.version",
	Dirty:   "schema_migrations.dirty",
}

// Generated where

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var SchemaMigrationWhere = struct {
	Version whereHelperint64
	Dirty   whereHelperbool
}{
	Version: whereHelperint64{field: "\"schema_migrations\".\"version\""},
	Dirty:   whereHelperbool{field: "\"schema_migrations\".\"dirty\""},
}

// SchemaMigrationRels is where relationship names are stored.
var SchemaMigrationRels = struct {
}{}

// schemaMigrationR is where relationships are stored.
type schemaMigrationR struct {
}

// NewStruct creates a new relationship struct
func (*schemaMigrationR) NewStruct() *schemaMigrationR {
	return &schemaMigrationR{}
}

// schemaMigrationL is where Load methods for each relationship are stored.
type schemaMigrationL struct{}

var (
	schemaMigrationAllColumns            = []string{"version", "dirty"}
	schemaMigrationColumnsWithoutDefault = []string{"version", "dirty"}
	schemaMigrationColumnsWithDefault    = []string{}
	schemaMigrationPrimaryKeyColumns     = []string{"version"}
	schemaMigrationGeneratedColumns      = []string{}
)

type (
	// SchemaMigrationSlice is an alias for a slice of pointers to SchemaMigration.
	// This should almost always be used instead of []SchemaMigration.
	SchemaMigrationSlice []*SchemaMigration

	schemaMigrationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	schemaMigrationType                 = reflect.TypeOf(&SchemaMigration{})
	schemaMigrationMapping              = queries.MakeStructMapping(schemaMigrationType)
	schemaMigrationPrimaryKeyMapping, _ = queries.BindMapping(schemaMigrationType, schemaMigrationMapping, schemaMigrationPrimaryKeyColumns)
	schemaMigrationInsertCacheMut       sync.RWMutex
	schemaMigrationInsertCache          = make(map[string]insertCache)
	schemaMigrationUpdateCacheMut       sync.RWMutex
	schemaMigrationUpdateCache          = make(map[string]updateCache)
	schemaMigrationUpsertCacheMut       sync.RWMutex
	schemaMigrationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single schemaMigration record from the query.
func (q schemaMigrationQuery) One(ctx context.Context, exec boil.ContextExecutor) (*SchemaMigration, error) {
	o := &SchemaMigration{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodel: failed to execute a one query for schema_migrations")
	}

	return o, nil
}

// All returns all SchemaMigration records from the query.
func (q schemaMigrationQuery) All(ctx context.Context, exec boil.ContextExecutor) (SchemaMigrationSlice, error) {
	var o []*SchemaMigration

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "dbmodel: failed to assign all query results to SchemaMigration slice")
	}

	return o, nil
}

// Count returns the count of all SchemaMigration records in the query.
func (q schemaMigrationQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: failed to count schema_migrations rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q schemaMigrationQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "dbmodel: failed to check if schema_migrations exists")
	}

	return count > 0, nil
}

// SchemaMigrations retrieves all the records using an executor.
func SchemaMigrations(mods ...qm.QueryMod) schemaMigrationQuery {
	mods = append(mods, qm.From("\"schema_migrations\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"schema_migrations\".*"})
	}

	return schemaMigrationQuery{q}
}

// FindSchemaMigration retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSchemaMigration(ctx context.Context, exec boil.ContextExecutor, version int64, selectCols ...string) (*SchemaMigration, error) {
	schemaMigrationObj := &SchemaMigration{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"schema_migrations\" where \"version\"=$1", sel,
	)

	q := queries.Raw(query, version)

	err := q.Bind(ctx, exec, schemaMigrationObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodel: unable to select from schema_migrations")
	}

	return schemaMigrationObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *SchemaMigration) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("dbmodel: no schema_migrations provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(schemaMigrationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	schemaMigrationInsertCacheMut.RLock()
	cache, cached := schemaMigrationInsertCache[key]
	schemaMigrationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			schemaMigrationAllColumns,
			schemaMigrationColumnsWithDefault,
			schemaMigrationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(schemaMigrationType, schemaMigrationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(schemaMigrationType, schemaMigrationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"schema_migrations\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"schema_migrations\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "dbmodel: unable to insert into schema_migrations")
	}

	if !cached {
		schemaMigrationInsertCacheMut.Lock()
		schemaMigrationInsertCache[key] = cache
		schemaMigrationInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the SchemaMigration.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *SchemaMigration) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	schemaMigrationUpdateCacheMut.RLock()
	cache, cached := schemaMigrationUpdateCache[key]
	schemaMigrationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			schemaMigrationAllColumns,
			schemaMigrationPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("dbmodel: unable to update schema_migrations, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"schema_migrations\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, schemaMigrationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(schemaMigrationType, schemaMigrationMapping, append(wl, schemaMigrationPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to update schema_migrations row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: failed to get rows affected by update for schema_migrations")
	}

	if !cached {
		schemaMigrationUpdateCacheMut.Lock()
		schemaMigrationUpdateCache[key] = cache
		schemaMigrationUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q schemaMigrationQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to update all for schema_migrations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to retrieve rows affected for schema_migrations")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SchemaMigrationSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("dbmodel: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), schemaMigrationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"schema_migrations\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, schemaMigrationPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to update all in schemaMigration slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to retrieve rows affected all in update all schemaMigration")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *SchemaMigration) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("dbmodel: no schema_migrations provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(schemaMigrationColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	schemaMigrationUpsertCacheMut.RLock()
	cache, cached := schemaMigrationUpsertCache[key]
	schemaMigrationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			schemaMigrationAllColumns,
			schemaMigrationColumnsWithDefault,
			schemaMigrationColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			schemaMigrationAllColumns,
			schemaMigrationPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("dbmodel: unable to upsert schema_migrations, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(schemaMigrationPrimaryKeyColumns))
			copy(conflict, schemaMigrationPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"schema_migrations\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(schemaMigrationType, schemaMigrationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(schemaMigrationType, schemaMigrationMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "dbmodel: unable to upsert schema_migrations")
	}

	if !cached {
		schemaMigrationUpsertCacheMut.Lock()
		schemaMigrationUpsertCache[key] = cache
		schemaMigrationUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single SchemaMigration record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *SchemaMigration) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("dbmodel: no SchemaMigration provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), schemaMigrationPrimaryKeyMapping)
	sql := "DELETE FROM \"schema_migrations\" WHERE \"version\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to delete from schema_migrations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: failed to get rows affected by delete for schema_migrations")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q schemaMigrationQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("dbmodel: no schemaMigrationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to delete all from schema_migrations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: failed to get rows affected by deleteall for schema_migrations")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SchemaMigrationSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), schemaMigrationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"schema_migrations\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, schemaMigrationPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: unable to delete all from schemaMigration slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodel: failed to get rows affected by deleteall for schema_migrations")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *SchemaMigration) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSchemaMigration(ctx, exec, o.Version)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SchemaMigrationSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SchemaMigrationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), schemaMigrationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"schema_migrations\".* FROM \"schema_migrations\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, schemaMigrationPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "dbmodel: unable to reload all in SchemaMigrationSlice")
	}

	*o = slice

	return nil
}

// SchemaMigrationExists checks if the SchemaMigration row exists.
func SchemaMigrationExists(ctx context.Context, exec boil.ContextExecutor, version int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"schema_migrations\" where \"version\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, version)
	}
	row := exec.QueryRowContext(ctx, sql, version)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "dbmodel: unable to check if schema_migrations exists")
	}

	return exists, nil
}

// Exists checks if the SchemaMigration row exists.
func (o *SchemaMigration) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return SchemaMigrationExists(ctx, exec, o.Version)
}
