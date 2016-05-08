package models

import (
	"taskManagerServices/config"
	"taskManagerServices/errorHandler"
)

const (
	dbInsertQuery string = "insert into task(task,priority,userid)  VALUES($1,$2,$3);"
	dbDeleteQuery string = "delete from task where taskid=$1 and userid = $2"
	dbUpdateQuery string = "update task set task=$1,priority=$2 where taskiD=$3 and userid = $4;"
)


type Task struct {
	TaskId int
	TaskDescription,Priority string
}

func(task *Task) Create(context config.Context, userId string)error{
	_,err := context.Db.Exec(dbInsertQuery, task.TaskDescription, task.Priority,userId)
	if err != nil {
		errorHandler.ErrorHandler(context.ErrorLogFile,err)
		return err
	}
	return nil
}

func (task *Task)Delete(context config.Context,userId string) error{
	_,err := context.Db.Exec(dbDeleteQuery,task.TaskId,userId)
	if err != nil {
		errorHandler.ErrorHandler(context.ErrorLogFile,err)
		return err
	}
	return nil
}

func (task *Task)Update(context config.Context, userId string) error {
	_,err := context.Db.Exec(dbUpdateQuery,task.TaskDescription,task.Priority,task.TaskId,userId)
	if err != nil {
		errorHandler.ErrorHandler(context.ErrorLogFile,err)
		return err
	}
	return nil
}