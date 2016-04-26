package converters

import (
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestConvertRowsToStructObjects(t *testing.T) {
	rows := sqlmock.NewRows([]string{"Have water","High"})
	ConvertRowsToStructObjects(rows)
}
