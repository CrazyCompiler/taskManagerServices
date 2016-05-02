package models

import (
	"encoding/json"
	"taskManagerServices/config"
	"taskManagerServices/errorHandler"
	"taskManagerServices/converters"
	"taskManagerServices/fileReaders"
	"encoding/csv"
)

const (
	dbSelectQuery string = "select taskId,task,priority from tasks;"
	dbInsertQuery string = "insert into tasks(task,priority)  VALUES($1,$2) returning taskId;"
	dbDeleteQuery string = "delete from tasks where taskId=$1"
	dbUpdateQuery string = "update tasks set task=$1,priority=$2 where taskID=$3;"
)


func Get(configObject config.ContextObject) ([]byte,error) {
	rows, err := configObject.Db.Query(dbSelectQuery)
	if err != nil {
		errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
	}
	dbData := converters.ToStructObjects(rows)
	data, err := json.Marshal(dbData)
	if err != nil {
		errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
	}
	return data,err
}

func Add(configObject config.ContextObject, taskDescription string, priority string) error {
	task := Task{taskDescription,priority}
	return task.Create(configObject)
}

func Delete(configObject config.ContextObject, taskId int) error {
	_,err := configObject.Db.Exec(dbDeleteQuery,taskId)
	if err != nil {
		return err
	}
	return nil
}

func Update(configObject config.ContextObject, taskId int32, data string, priority string)error{
	_,err := configObject.Db.Exec(dbUpdateQuery, data, priority,taskId)
	if err != nil {
		errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
		return err
	}
	return nil
}


func AddTaskByCsv(configObject config.ContextObject,data string) error{
	separatedData,err := fileReaders.ReadTaskCsv(data)
	if err != nil {
		errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
		return err
	}

	for _, each := range separatedData {
		err := Add(configObject,each.TASK ,each.PRIORITY)
		if err != nil {
			errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
			return err
		}
	}
	return  nil
}

type Wr struct {
	buf []byte
}

func (w *Wr) Write(b []byte) (int, error) {
	w.buf = append(w.buf, b...)
	return len(w.buf), nil
}

func GetCsv(configObject config.ContextObject) ([]byte,error) {
	rows, err := configObject.Db.Query(dbSelectQuery)
	if err != nil {
		errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
	}
	dbData := converters.ToArrayOfString(rows)
	wrt := &Wr{}
	w := csv.NewWriter(wrt)
	w.WriteAll(dbData)
	w.Flush()
	err = w.Error()
	return 	wrt.buf,err
}