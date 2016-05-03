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