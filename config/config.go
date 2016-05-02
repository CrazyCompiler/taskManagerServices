package config

import (
	"os"
	"database/sql"
)

type Context struct {
	ErrorLogFile *os.File
	Db *sql.DB
}
