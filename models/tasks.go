package models

import (
	"encoding/json"
	"taskManagerServices/config"
	"taskManagerServices/errorHandler"
	"taskManagerServices/converters"
	"taskManagerServices/fileReaders"
)

const (
	dbSelectQuery string = "select taskId,task,priority from tasks;"

)


func Get(context config.Context) ([]byte,error) {
	rows, err := context.Db.Query(dbSelectQuery)
	if err != nil {
		errorHandler.ErrorHandler(context.ErrorLogFile,err)
	}
	dbData := converters.ToStructObjects(rows)
	data, err := json.Marshal(dbData)
	if err != nil {
		errorHandler.ErrorHandler(context.ErrorLogFile,err)
	}
	return data,err
}

func AddTaskByCsv(context config.Context,data string) error{
	separatedData,err := fileReaders.ReadTaskCsv(data)
	if err != nil {
		errorHandler.ErrorHandler(context.ErrorLogFile,err)
		return err
	}

	for _, each := range separatedData {
		newTask := Task{}
		newTask.TaskDescription = each.TASK
		newTask.Priority=each.PRIORITY
		err := newTask.Create(context)
		if err != nil {
			errorHandler.ErrorHandler(context.ErrorLogFile,err)
			return err
		}
	}
	return  nil
}

func GetCsv(context config.Context) ([][]string,error) {
	rows, err := context.Db.Query(dbSelectQuery)
	if err != nil {
		errorHandler.ErrorHandler(context.ErrorLogFile,err)
	}
	dbData := converters.ToArrayOfString(rows)
	return 	dbData,err
}