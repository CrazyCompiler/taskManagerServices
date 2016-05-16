package models

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/DATA-DOG/go-sqlmock"
	"os"
	"taskManagerServices/config"
	"fmt"
)

func TestTaskCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.Nil(t, err)

	newTask := Task{
		TaskDescription:"Drinking Water",
		Priority:"High",
	}

	user_id := "123432"

	mock.ExpectExec("insert into tasks").
	WithArgs(newTask.TaskDescription, newTask.Priority, user_id).
	WillReturnResult(sqlmock.NewResult(1, 1))

	context := config.Context{}
	context.ErrorLogFile, err = os.OpenFile(errorLogFilePath, os.O_APPEND | os.O_WRONLY, 0600)
	context.Db = db

	err = newTask.Create(context, user_id)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestTaskCreateForError(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	newTask := Task{
		TaskDescription:"Drinking Water",
		Priority:"High",
	}

	user_id := "123432"

	mock.ExpectExec("insert into tasks").
	WithArgs(newTask.TaskDescription, newTask.Priority, user_id).
	WillReturnError(fmt.Errorf("some Error"))

	context := config.Context{}
	context.ErrorLogFile, err = os.OpenFile(errorLogFilePath, os.O_APPEND | os.O_WRONLY, 0600)
	context.Db = db

	err = newTask.Create(context, user_id)
	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}

}

func TestTaskDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.Nil(t, err)

	taskToDelete := Task{
		TaskId:420,
	}
	user_id := "123432"

	mock.ExpectExec("delete from tasks").
	WithArgs(taskToDelete.TaskId, user_id).
	WillReturnResult(sqlmock.NewResult(1, 1))

	context := config.Context{}
	context.ErrorLogFile, err = os.OpenFile(errorLogFilePath, os.O_APPEND | os.O_WRONLY, 0600)
	context.Db = db

	err = taskToDelete.Delete(context, user_id)

	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestTaskDeleteForError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.Nil(t, err)

	taskToDelete := Task{
		TaskId:420,
	}

	user_id := "123432"

	mock.ExpectExec("delete from tasks").
	WithArgs(taskToDelete.TaskId, user_id).
	WillReturnError(fmt.Errorf("some Error"))

	context := config.Context{}
	context.ErrorLogFile, err = os.OpenFile(errorLogFilePath, os.O_APPEND | os.O_WRONLY, 0600)
	context.Db = db

	err = taskToDelete.Delete(context, user_id)
	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}

}

func TestTaskUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.Nil(t, err)

	taskToUpdate := Task{
		TaskId:420,
		TaskDescription:"Lets have tea",
		Priority:"Low",
	}
	user_id := "123432"

	mock.ExpectExec("update tasks set").
	WithArgs(taskToUpdate.TaskDescription, taskToUpdate.Priority, taskToUpdate.TaskId, user_id).
	WillReturnResult(sqlmock.NewResult(1, 1))

	context := config.Context{}
	context.ErrorLogFile, err = os.OpenFile(errorLogFilePath, os.O_APPEND | os.O_WRONLY, 0600)
	context.Db = db

	err = taskToUpdate.Update(context, user_id)

	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled exceptions: %s", err)
	}

}

func TestTaskUpdateForError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.Nil(t, err)

	taskToUpdate := Task{
		TaskId:420,
		TaskDescription:"Drinking Water",
		Priority:"High",
	}

	user_id := "123432"

	mock.ExpectExec("update tasks").
	WithArgs(taskToUpdate.TaskDescription, taskToUpdate.Priority,taskToUpdate.TaskId ,user_id).
	WillReturnError(fmt.Errorf("some Error"))

	context := config.Context{}
	context.ErrorLogFile, err = os.OpenFile(errorLogFilePath, os.O_APPEND | os.O_WRONLY, 0600)
	context.Db = db

	err = taskToUpdate.Update(context, user_id)
	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}

}