// Package sqlite3 defines and registers usql's SQLite3 driver. Requires CGO.
//
// See: https://github.com/mattn/go-sqlite3
// Group: base
package sqlite3

import (
	"context"
	"strconv"

	// "github.com/mattn/go-sqlite3" // DRIVER
	"github.com/glebarez/go-sqlite" // DRIVER
	"github.com/xo/usql/drivers"
	"github.com/xo/usql/drivers/sqlite3/sqshared"
)

func init() {
	drivers.Register("sqlite", drivers.Driver{
		AllowMultilineComments: true,
		ForceParams: drivers.ForceQueryParameters([]string{
			"loc", "auto",
		}),
		Version: func(ctx context.Context, db drivers.DB) (string, error) {
			var ver string
			err := db.QueryRowContext(ctx, `SELECT sqlite_version()`).Scan(&ver)
			if err != nil {
				return "", err
			}
			return "SQLite3 " + ver, nil
		},
		Err: func(err error) (string, string) {
			if e, ok := err.(*sqlite.Error); ok {
				return strconv.Itoa(int(e.Code())), e.Error()
			}
			code, msg := "", err.Error()
			// if e, ok := err.(*sqlite.ErrNo); ok {
			// 	code = strconv.Itoa(int(e))
			// }
			return code, msg
		},
		ConvertBytes:      sqshared.ConvertBytes,
		NewMetadataReader: sqshared.NewMetadataReader,
		Copy:              drivers.CopyWithInsert(func(int) string { return "?" }),
	})
}
