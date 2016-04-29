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

func responseGenerator(status int,errorBody string) (contract.Response){
	response := contract.Response{}
	serverStatus := int32(status)
	errorResponse := &contract.Error{}
	errorResponse.Status = &serverStatus
	errorResponse.Description = &errorBody
	response.Err = errorResponse
	return response
}

func AddTask(configObject config.ContextObject) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		requestData, err := ioutil.ReadAll(req.Body)
		if err != nil {
			errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
		}
		data := &contract.Task{}
		err = proto.Unmarshal(requestData,data)
		if err != nil {
			errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
		}
		err = models.Add(configObject, *data.Task, *data.Priority)
		if err != nil {
			resp := responseGenerator(http.StatusInternalServerError,err.Error())
			dataToBeSend,err :=  proto.Marshal(&resp)
			if err != nil {
				errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
			}
			res.Write(dataToBeSend)
			res.WriteHeader(http.StatusInternalServerError)
		}
		res.WriteHeader(http.StatusCreated)
	}
}

func GetTasks(configObject config.ContextObject) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		data,err := models.Get(configObject)
		if err != nil {
			resp := responseGenerator(http.StatusInternalServerError,err.Error())
			dataToBeSend,err :=  proto.Marshal(&resp)
			if err != nil {
				errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
			}
			res.Write(dataToBeSend)
			res.WriteHeader(http.StatusInternalServerError)
		}
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
			resp := responseGenerator(http.StatusInternalServerError,err.Error())
			dataToBeSend,err :=  proto.Marshal(&resp)
			if err != nil {
				errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
			}
			res.Write(dataToBeSend)
			res.WriteHeader(http.StatusInternalServerError)
		}
		res.WriteHeader(http.StatusAccepted)
	}
}

func UpdateTask(configObject config.ContextObject)http.HandlerFunc{
	return func(res http.ResponseWriter,req *http.Request) {
		requestData, err := ioutil.ReadAll(req.Body)
		if err != nil {
			errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
		}
		data := &contract.Task{}
		err = proto.Unmarshal(requestData,data)
		if err != nil {
			errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
		}
		err = models.Update(configObject, *data.TaskId, *data.Task,*data.Priority)
		if err != nil {
			resp := responseGenerator(http.StatusInternalServerError,err.Error())
			dataToBeSend,err :=  proto.Marshal(&resp)
			if err != nil {
				errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
			}
			res.Write(dataToBeSend)
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
			resp := responseGenerator(http.StatusInternalServerError,err.Error())
			dataToBeSend,err :=  proto.Marshal(&resp)
			if err != nil {
				errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
			}
			res.Write(dataToBeSend)
		}
		res.WriteHeader(http.StatusOK)
	}
}

