package database

import (
	"testing"
	"taskManager/fileReaders"
	"github.com/stretchr/testify/assert"
)

func TestCreateDbInfo(t *testing.T) {
	jsonObject := fileReaders.JsonObject{}
	jsonObject.DB_NAME = "sql"
	jsonObject.DB_PASSWORD = "something"
	jsonObject.DB_SCHEMA = "students"
	jsonObject.DB_USER = "sqlUser"

	expected := "user=sqlUser password=something dbname=sql sslmode=disable SEARCH_PATH=students "

	assert.Equal(t,expected,CreateDbInfo(jsonObject))

}
