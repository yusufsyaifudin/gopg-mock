package gopg_mock

import (
	"github.com/go-pg/pg/v10/orm"
)

// Formatter implements orm.Formatter
type Formatter struct {
}

func (f *Formatter) FormatQuery(b []byte, query string, params ...interface{}) []byte {
	formatter := new(orm.Formatter)
	got := formatter.FormatQuery(b, query, params...)
	return got
}
