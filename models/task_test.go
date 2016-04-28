package models

import (
	"testing"
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"taskManagerServices/config"
	"os"
)

const errorLogFilePath string = "../errorLog"

func TestGet(t *testing.T) {
	db,mock, err := sqlmock.New()
	assert.Nil(t,err)
	rows := sqlmock.NewRows([]string{"hello","High"})
	mock.ExpectQuery("select taskId,task,priority from tasks;").WillReturnRows(rows)

	contextObject := config.ContextObject{}
	contextObject.ErrorLogFile, err = os.OpenFile(errorLogFilePath, os.O_APPEND|os.O_WRONLY, 0600)
	contextObject.Db = db

	Get(contextObject)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestAdd(t *testing.T) {
	db,mock, err := sqlmock.New()
	assert.Nil(t,err)

	task := "dring water"
	priority := "High"

	mock.ExpectExec("insert into tasks").
	WithArgs(task,priority).
	WillReturnResult(sqlmock.NewResult(1, 1))

	contextObject := config.ContextObject{}
	contextObject.ErrorLogFile, err = os.OpenFile(errorLogFilePath, os.O_APPEND|os.O_WRONLY, 0600)
	contextObject.Db = db

	Add(contextObject,task,priority)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestDelete(t *testing.T) {
	db,mock, err := sqlmock.New()
	assert.Nil(t,err)

	taskId := 1

	mock.ExpectExec("delete from tasks").
	WithArgs(taskId).
	WillReturnResult(sqlmock.NewResult(1, 1))

	contextObject := config.ContextObject{}
	contextObject.ErrorLogFile, err = os.OpenFile(errorLogFilePath, os.O_APPEND|os.O_WRONLY, 0600)
	contextObject.Db = db

	Delete(contextObject,taskId)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestUpdatePriority(t *testing.T) {
	db,mock, err := sqlmock.New()
	assert.Nil(t,err)

	taskId := 1
	priority := "High"

	mock.ExpectExec("update tasks set ").
	WithArgs(priority,int32(taskId)).
	WillReturnResult(sqlmock.NewResult(0, 1))

	contextObject := config.ContextObject{}
	contextObject.ErrorLogFile, err = os.OpenFile(errorLogFilePath, os.O_APPEND|os.O_WRONLY, 0600)
	contextObject.Db = db

	UpdatePriority(contextObject,int32(taskId),priority)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestUpdateTaskDescription(t *testing.T) {
	db,mock, err := sqlmock.New()
	assert.Nil(t,err)

	taskId := 1
	data := "Drink water"

	mock.ExpectExec("update tasks set ").
	WithArgs(data,int32(taskId)).
	WillReturnResult(sqlmock.NewResult(0, 1))

	contextObject := config.ContextObject{}
	contextObject.ErrorLogFile, err = os.OpenFile(errorLogFilePath, os.O_APPEND|os.O_WRONLY, 0600)
	contextObject.Db = db

	UpdateTaskDescription(contextObject,int32(taskId),data)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}