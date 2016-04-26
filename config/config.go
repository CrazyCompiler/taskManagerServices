package config

import (
	"os"
	"database/sql"
)

type ContextObject struct {
	ErrorLogFile *os.File
	Db *sql.DB
}
