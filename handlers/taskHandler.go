package handlers

import (
	"strings"
	"strconv"
	"net/http"
	"taskManagerServices/config"
	"taskManagerServices/models"
	"taskManagerServices/errorHandler"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"github.com/CrazyCompiler/taskManagerContract"
)

func AddTask(configObject config.ContextObject) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		requestData, err := ioutil.ReadAll(req.Body)
		if err != nil {
			errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
		}
		data := &contract.AddTask{}
		err = proto.Unmarshal(requestData,data)
		if err != nil {
			errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
		}
		err = models.Add(configObject, *data.Task, *data.Priority)
		if err != nil {
			errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
			res.WriteHeader(http.StatusInternalServerError)
		}
		res.WriteHeader(http.StatusCreated)
	}
}

func GetTasks(configObject config.ContextObject) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		data := models.Get(configObject)
		res.WriteHeader(http.StatusOK)
		res.Write(data)
	}
}

func DeleteTask(configObject config.ContextObject) http.HandlerFunc {
	return func(res http.ResponseWriter,req *http.Request) {
		req.ParseForm()
		taskId := strings.Split(req.RequestURI,"/")[2]
		task,err := strconv.Atoi(taskId)
		err = models.Delete(configObject,task)
		if err != nil {
			errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
			res.WriteHeader(http.StatusInternalServerError)
		}
		res.WriteHeader(http.StatusAccepted)
	}
}

func UpdateTaskPriority(configObject config.ContextObject)http.HandlerFunc{
	return func(res http.ResponseWriter,req *http.Request) {
		requestData, err := ioutil.ReadAll(req.Body)
		if err != nil {
			errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
		}
		data := &contract.UpdatePriority{}
		err = proto.Unmarshal(requestData,data)
		if err != nil {
			errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
		}
		err = models.UpdatePriority(configObject, *data.TaskId, *data.Priority)
		if err != nil {
			errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
			res.WriteHeader(http.StatusInternalServerError)
		}
		res.WriteHeader(http.StatusCreated)
	}
}

func UpdateTaskDescription(configObject config.ContextObject)http.HandlerFunc{
	return func(res http.ResponseWriter,req *http.Request) {
		requestData, err := ioutil.ReadAll(req.Body)
		if err != nil {
			errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
		}
		data := &contract.UpdateTaskDescription{}
		err = proto.Unmarshal(requestData,data)
		if err != nil {
			errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
		}
		err = models.UpdateTaskDescription(configObject, *data.TaskId, *data.Data)
		if err != nil {
			errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
			res.WriteHeader(http.StatusInternalServerError)
		}
		res.WriteHeader(http.StatusCreated)
	}
}

func UploadCsv(configObject config.ContextObject) http.HandlerFunc{
	return func(res http.ResponseWriter,req *http.Request) {
		requestData, err := ioutil.ReadAll(req.Body)
		if err != nil {
			errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
		}
		data := &contract.UploadCsvData{}
		err = proto.Unmarshal(requestData,data)
		if err != nil {
			errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
		}
		err = models.AddTaskByCsv(configObject,string(*data.CsvData))
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
		}
		res.WriteHeader(http.StatusOK)
	}
}

