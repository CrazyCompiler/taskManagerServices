package models

import (
	"encoding/json"
	"taskManagerServices/config"
	"taskManagerServices/errorHandler"
	"taskManagerServices/converters"
	"taskManagerServices/fileReaders"
)

const (
	dbSelectQuery string = "select taskId,task,priority from tasks where user_id=$1;"
)

func Get(context config.Context, userId string) ([]byte, error) {
	rows, err := context.Db.Query(dbSelectQuery, userId)
	if err != nil {
		errorHandler.ErrorHandler(context.ErrorLogFile, err)
	}
	dbData := converters.ToStructObjects(rows)
	data, err := json.Marshal(dbData)
	if err != nil {
		errorHandler.ErrorHandler(context.ErrorLogFile, err)
	}
	return data, err
}

func AddTaskByCsv(context config.Context, userId string, reader fileReaders.Reader) error {
	separatedData, err := reader.Read()
	if err != nil {
		errorHandler.ErrorHandler(context.ErrorLogFile, err)
		return err
	}

	for _, each := range separatedData {
		newTask := Task{
			TaskDescription:each.TASK,
			Priority:each.PRIORITY,
		}
		err := newTask.Create(context, userId)
		if err != nil {
			errorHandler.ErrorHandler(context.ErrorLogFile, err)
			return err
		}
	}
	return nil
}

func GetCsv(context config.Context, userId string) ([][]string, error) {
	rows, err := context.Db.Query(dbSelectQuery, userId)
	if err != nil {
		errorHandler.ErrorHandler(context.ErrorLogFile, err)
	}
	dbData := converters.ToArrayOfString(rows)
	return dbData, err
}