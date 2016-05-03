package models

import (
	"taskManagerServices/config"
	"taskManagerServices/errorHandler"
)

const (
	dbInsertQuery string = "insert into tasks(task,priority)  VALUES($1,$2) returning taskId;"
	dbDeleteQuery string = "delete from tasks where taskId=$1"
	dbUpdateQuery string = "update tasks set task=$1,priority=$2 where taskID=$3;"
)

type Task struct {
	TaskId int
	TaskDescription,Priority string
}

func(task *Task) Create(context config.Context)error{
	_,err := context.Db.Exec(dbInsertQuery, task.TaskDescription, task.Priority)
	if err != nil {
		errorHandler.ErrorHandler(context.ErrorLogFile,err)
		return err
	}
	return nil
}

func (task *Task)Delete(context config.Context) error{
	_,err := context.Db.Exec(dbDeleteQuery,task.TaskId)
	if err != nil {
		errorHandler.ErrorHandler(context.ErrorLogFile,err)
		return err
	}
	return nil
}

func (task *Task)Update(context config.Context) error {
	_,err := context.Db.Exec(dbUpdateQuery,task.TaskDescription,task.Priority,task.TaskId)
	if err != nil {
		errorHandler.ErrorHandler(context.ErrorLogFile,err)
		return err
	}
	return nil
}