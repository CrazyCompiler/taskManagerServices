package models

import (
	"taskManagerServices/config"
	"taskManagerServices/errorHandler"
)

type Task struct {
	taskDescription,priority string
}

func(task *Task) Create(context config.Context)error{
	_,err := context.Db.Exec(dbInsertQuery, task.taskDescription, task.priority)
	if err != nil {
		errorHandler.ErrorHandler(context.ErrorLogFile,err)
		return err
	}
	return nil
}
