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

func Add(context config.Context, taskDescription string, priority string) error {
	task := Task{taskDescription,priority}
	return task.Create(context)
}

func Delete(context config.Context, taskId int) error {
	_,err := context.Db.Exec(dbDeleteQuery,taskId)
	if err != nil {
		return err
	}
	return nil
}

func Update(context config.Context, taskId int32, data string, priority string)error{
	_,err := context.Db.Exec(dbUpdateQuery, data, priority,taskId)
	if err != nil {
		errorHandler.ErrorHandler(context.ErrorLogFile,err)
		return err
	}
	return nil
}


func AddTaskByCsv(context config.Context,data string) error{
	separatedData,err := fileReaders.ReadTaskCsv(data)
	if err != nil {
		errorHandler.ErrorHandler(context.ErrorLogFile,err)
		return err
	}

	for _, each := range separatedData {
		err := Add(context,each.TASK ,each.PRIORITY)
		if err != nil {
			errorHandler.ErrorHandler(context.ErrorLogFile,err)
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

func GetCsv(context config.Context) ([]byte,error) {
	rows, err := context.Db.Query(dbSelectQuery)
	if err != nil {
		errorHandler.ErrorHandler(context.ErrorLogFile,err)
	}
	dbData := converters.ToArrayOfString(rows)
	wrt := &Wr{}
	w := csv.NewWriter(wrt)
	w.WriteAll(dbData)
	w.Flush()
	err = w.Error()
	return 	wrt.buf,err
}