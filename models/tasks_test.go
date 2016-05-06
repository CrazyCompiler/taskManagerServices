package models

import (
)
import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
	"github.com/DATA-DOG/go-sqlmock"
	"taskManagerServices/config"
)

const errorLogFilePath string = "../errorLog"

func TestGet(t *testing.T) {
	db,mock, err := sqlmock.New()
	assert.Nil(t,err)

	user_id := "1234"
	rows := sqlmock.NewRows([]string{"hello","High"})
	mock.ExpectQuery("select taskId,task,priority from tasks").
	WithArgs(user_id).
	WillReturnRows(rows)

	contextObject := config.Context{}
	contextObject.ErrorLogFile, err = os.OpenFile(errorLogFilePath, os.O_APPEND|os.O_WRONLY, 0600)
	contextObject.Db = db

	Get(contextObject,user_id)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}