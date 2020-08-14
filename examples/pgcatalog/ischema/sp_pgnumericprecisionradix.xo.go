// Package ischema contains the types for schema 'information_schema'.
package ischema

import "github.com/8pockets/xo/examples/pgcatalog/pgtypes"

// Code generated by xo. DO NOT EDIT.

// PgNumericPrecisionRadix calls the stored procedure 'information_schema._pg_numeric_precision_radix(oid, integer) integer' on db.
func PgNumericPrecisionRadix(db XODB, v0 pgtypes.Oid, v1 int) (int, error) {
	var err error

	// sql query
	const sqlstr = `SELECT information_schema._pg_numeric_precision_radix($1, $2)`

	// run query
	var ret int
	XOLog(sqlstr, v0, v1)
	err = db.QueryRow(sqlstr, v0, v1).Scan(&ret)
	if err != nil {
		return 0, err
	}

	return ret, nil
}
