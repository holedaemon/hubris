// Code generated by SQLBoiler 4.8.3 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

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

// Pronoun is an object representing the database table.
type Pronoun struct {
	ID        int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	RoleID    string    `boil:"role_id" json:"role_id" toml:"role_id" yaml:"role_id"`
	GuildID   string    `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	Pronoun   string    `boil:"pronoun" json:"pronoun" toml:"pronoun" yaml:"pronoun"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *pronounR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L pronounL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PronounColumns = struct {
	ID        string
	RoleID    string
	GuildID   string
	Pronoun   string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	RoleID:    "role_id",
	GuildID:   "guild_id",
	Pronoun:   "pronoun",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var PronounTableColumns = struct {
	ID        string
	RoleID    string
	GuildID   string
	Pronoun   string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "pronouns.id",
	RoleID:    "pronouns.role_id",
	GuildID:   "pronouns.guild_id",
	Pronoun:   "pronouns.pronoun",
	CreatedAt: "pronouns.created_at",
	UpdatedAt: "pronouns.updated_at",
}

// Generated where

var PronounWhere = struct {
	ID        whereHelperint64
	RoleID    whereHelperstring
	GuildID   whereHelperstring
	Pronoun   whereHelperstring
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	ID:        whereHelperint64{field: "\"pronouns\".\"id\""},
	RoleID:    whereHelperstring{field: "\"pronouns\".\"role_id\""},
	GuildID:   whereHelperstring{field: "\"pronouns\".\"guild_id\""},
	Pronoun:   whereHelperstring{field: "\"pronouns\".\"pronoun\""},
	CreatedAt: whereHelpertime_Time{field: "\"pronouns\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"pronouns\".\"updated_at\""},
}

// PronounRels is where relationship names are stored.
var PronounRels = struct {
}{}

// pronounR is where relationships are stored.
type pronounR struct {
}

// NewStruct creates a new relationship struct
func (*pronounR) NewStruct() *pronounR {
	return &pronounR{}
}

// pronounL is where Load methods for each relationship are stored.
type pronounL struct{}

var (
	pronounAllColumns            = []string{"id", "role_id", "guild_id", "pronoun", "created_at", "updated_at"}
	pronounColumnsWithoutDefault = []string{"role_id", "guild_id", "pronoun"}
	pronounColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	pronounPrimaryKeyColumns     = []string{"id"}
)

type (
	// PronounSlice is an alias for a slice of pointers to Pronoun.
	// This should almost always be used instead of []Pronoun.
	PronounSlice []*Pronoun

	pronounQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	pronounType                 = reflect.TypeOf(&Pronoun{})
	pronounMapping              = queries.MakeStructMapping(pronounType)
	pronounPrimaryKeyMapping, _ = queries.BindMapping(pronounType, pronounMapping, pronounPrimaryKeyColumns)
	pronounInsertCacheMut       sync.RWMutex
	pronounInsertCache          = make(map[string]insertCache)
	pronounUpdateCacheMut       sync.RWMutex
	pronounUpdateCache          = make(map[string]updateCache)
	pronounUpsertCacheMut       sync.RWMutex
	pronounUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single pronoun record from the query.
func (q pronounQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Pronoun, error) {
	o := &Pronoun{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for pronouns")
	}

	return o, nil
}

// All returns all Pronoun records from the query.
func (q pronounQuery) All(ctx context.Context, exec boil.ContextExecutor) (PronounSlice, error) {
	var o []*Pronoun

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Pronoun slice")
	}

	return o, nil
}

// Count returns the count of all Pronoun records in the query.
func (q pronounQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count pronouns rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q pronounQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if pronouns exists")
	}

	return count > 0, nil
}

// Pronouns retrieves all the records using an executor.
func Pronouns(mods ...qm.QueryMod) pronounQuery {
	mods = append(mods, qm.From("\"pronouns\""))
	return pronounQuery{NewQuery(mods...)}
}

// FindPronoun retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPronoun(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Pronoun, error) {
	pronounObj := &Pronoun{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"pronouns\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, pronounObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from pronouns")
	}

	return pronounObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Pronoun) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no pronouns provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(pronounColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	pronounInsertCacheMut.RLock()
	cache, cached := pronounInsertCache[key]
	pronounInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			pronounAllColumns,
			pronounColumnsWithDefault,
			pronounColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(pronounType, pronounMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(pronounType, pronounMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"pronouns\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"pronouns\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into pronouns")
	}

	if !cached {
		pronounInsertCacheMut.Lock()
		pronounInsertCache[key] = cache
		pronounInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Pronoun.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Pronoun) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	pronounUpdateCacheMut.RLock()
	cache, cached := pronounUpdateCache[key]
	pronounUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			pronounAllColumns,
			pronounPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update pronouns, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"pronouns\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, pronounPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(pronounType, pronounMapping, append(wl, pronounPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	_, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update pronouns row")
	}

	if !cached {
		pronounUpdateCacheMut.Lock()
		pronounUpdateCache[key] = cache
		pronounUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAll updates all rows with the specified column values.
func (q pronounQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for pronouns")
	}

	return nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PronounSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pronounPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"pronouns\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, pronounPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in pronoun slice")
	}

	return nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Pronoun) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no pronouns provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(pronounColumnsWithDefault, o)

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

	pronounUpsertCacheMut.RLock()
	cache, cached := pronounUpsertCache[key]
	pronounUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			pronounAllColumns,
			pronounColumnsWithDefault,
			pronounColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			pronounAllColumns,
			pronounPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert pronouns, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(pronounPrimaryKeyColumns))
			copy(conflict, pronounPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"pronouns\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(pronounType, pronounMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(pronounType, pronounMapping, ret)
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
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert pronouns")
	}

	if !cached {
		pronounUpsertCacheMut.Lock()
		pronounUpsertCache[key] = cache
		pronounUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Pronoun record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Pronoun) Delete(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil {
		return errors.New("models: no Pronoun provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), pronounPrimaryKeyMapping)
	sql := "DELETE FROM \"pronouns\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from pronouns")
	}

	return nil
}

// DeleteAll deletes all matching rows.
func (q pronounQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if q.Query == nil {
		return errors.New("models: no pronounQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from pronouns")
	}

	return nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PronounSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pronounPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"pronouns\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, pronounPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from pronoun slice")
	}

	return nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Pronoun) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPronoun(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PronounSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PronounSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pronounPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"pronouns\".* FROM \"pronouns\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, pronounPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PronounSlice")
	}

	*o = slice

	return nil
}

// PronounExists checks if the Pronoun row exists.
func PronounExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"pronouns\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if pronouns exists")
	}

	return exists, nil
}
