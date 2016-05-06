package models

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/DATA-DOG/go-sqlmock"
	"os"
	"taskManagerServices/config"
	"fmt"
)

func TestTask_Create(t *testing.T) {
	db,mock, err := sqlmock.New()
	assert.Nil(t,err)

	newTask := Task{
		TaskDescription:"Drinking Water",
		Priority:"High",
	}

	user_id := "123432"

	mock.ExpectExec("insert into tasks").
	WithArgs(newTask.TaskDescription,newTask.Priority,user_id).
	WillReturnResult(sqlmock.NewResult(1, 1))

	context := config.Context{}
	context.ErrorLogFile, err = os.OpenFile(errorLogFilePath, os.O_APPEND|os.O_WRONLY, 0600)
	context.Db = db

	err =newTask.Create(context,user_id)
	assert.Equal(t,nil,err,"error will be nil")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestTask_Create_for_error(t *testing.T) {
	db,mock, err := sqlmock.New()
	assert.Nil(t,err)

	newTask := Task{
		TaskDescription:"Drinking Water",
		Priority:"High",
	}

	user_id := "123432"

	mock.ExpectExec("insert into tasks").
	WithArgs(newTask.TaskDescription,newTask.Priority,"567").
	WillReturnError(fmt.Errorf("some Error"))

	mock.ExpectRollback()

	context := config.Context{}
	context.ErrorLogFile, err = os.OpenFile(errorLogFilePath, os.O_APPEND|os.O_WRONLY, 0600)
	context.Db = db

	err =newTask.Create(context,user_id)
	assert.NotEqual(t,nil,err,"error will not be nil")

}

func TestTask_Delete(t *testing.T) {
	db,mock,err := sqlmock.New()
	assert.Nil(t,err)

	taskToDelete := Task{
		TaskId:420,
	}
	user_id := "123432"

	mock.ExpectExec("delete from tasks").
	WithArgs(taskToDelete.TaskId,user_id).
	WillReturnResult(sqlmock.NewResult(1,1))

	context := config.Context{}
	context.ErrorLogFile, err = os.OpenFile(errorLogFilePath, os.O_APPEND|os.O_WRONLY, 0600)
	context.Db = db

	taskToDelete.Delete(context,user_id)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestTask_Delete_for_error(t *testing.T) {
	db,mock, err := sqlmock.New()
	assert.Nil(t,err)

	taskToDelete := Task{
		TaskId:420,
	}

	user_id := "123432"

	mock.ExpectExec("insert into tasks").
	WithArgs(taskToDelete.TaskDescription,taskToDelete.Priority,"567").
	WillReturnError(fmt.Errorf("some Error"))

	mock.ExpectRollback()

	context := config.Context{}
	context.ErrorLogFile, err = os.OpenFile(errorLogFilePath, os.O_APPEND|os.O_WRONLY, 0600)
	context.Db = db

	err =taskToDelete.Delete(context,user_id)
	assert.NotEqual(t,nil,err,"error will not be nil")

}

func TestTask_Update(t *testing.T) {
	db,mock,err := sqlmock.New()
	assert.Nil(t,err)

	taskToUpdate := Task{
		TaskId:420,
		TaskDescription:"Lets have tea",
		Priority:"Low",
	}
	user_id := "123432"

	mock.ExpectExec("update tasks set").
	WithArgs(taskToUpdate.TaskDescription,taskToUpdate.Priority,taskToUpdate.TaskId,user_id).
	WillReturnResult(sqlmock.NewResult(1,1))

	context := config.Context{}
	context.ErrorLogFile, err = os.OpenFile(errorLogFilePath, os.O_APPEND|os.O_WRONLY, 0600)
	context.Db = db

	taskToUpdate.Update(context,user_id)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled exceptions: %s",err)
	}

}

func TestTask_Update_for_error(t *testing.T) {
	db,mock, err := sqlmock.New()
	assert.Nil(t,err)

	taskToUpdate := Task{
		TaskId:420,
		TaskDescription:"Drinking Water",
		Priority:"High",
	}

	user_id := "123432"

	mock.ExpectExec("insert into tasks").
	WithArgs(taskToUpdate.TaskDescription,taskToUpdate.Priority,"567").
	WillReturnError(fmt.Errorf("some Error"))

	mock.ExpectRollback()

	context := config.Context{}
	context.ErrorLogFile, err = os.OpenFile(errorLogFilePath, os.O_APPEND|os.O_WRONLY, 0600)
	context.Db = db

	err =taskToUpdate.Update(context,user_id)
	assert.NotEqual(t,nil,err,"error will not be nil")

}
