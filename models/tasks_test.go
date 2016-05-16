package models

import (
	"testing"
	"os"
	"github.com/DATA-DOG/go-sqlmock"
	"taskManagerServices/config"
	"taskManagerServices/fileReaders"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/assert"
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

type MockCsvReader struct {
	mock.Mock
}

func (m *MockCsvReader)Read()([]fileReaders.TableContent,error)  {
	m.Called()
	return nil,nil
}
func TestAddTaskByCsv(t *testing.T) {
	db,_, err := sqlmock.New()
	assert.Nil(t,err)

	user_id := "1234"

	m := &MockCsvReader{}
	m.On("Read").Return([]fileReaders.TableContent{},12)

	contextObject := config.Context{}
	contextObject.ErrorLogFile, err = os.OpenFile(errorLogFilePath, os.O_APPEND|os.O_WRONLY, 0600)
	contextObject.Db = db

	err = AddTaskByCsv(contextObject,user_id,m)
	assert.NoError(t,err)

}