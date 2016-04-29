package converters

import (
	"testing"
	"database/sql"
)

func TestConvertRowsToStructObjects(t *testing.T) {

	rows := &sql.Rows{}
	ConvertRowsToStructObjects(rows)
}
