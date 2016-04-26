package models

import (
	"taskManagerServices/config"
	"taskManagerServices/errorHandler"
)

type Task struct {
	taskDescription,priority string
}

func(task *Task) Create(configObject config.ContextObject)error{
	_,err := configObject.Db.Exec(dbInsertQuery, task.taskDescription, task.priority)
	if err != nil {
		errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
		return err
	}
	return nil
}
