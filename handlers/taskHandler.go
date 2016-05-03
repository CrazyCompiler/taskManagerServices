package handlers

import (
	"io/ioutil"
	"strings"
	"strconv"
	"taskManagerServices/config"
	"net/http"
	"taskManagerServices/errorHandler"
	"github.com/golang/protobuf/proto"
	"taskManagerServices/models"
	"taskManagerServices/vendor/github.com/CrazyCompiler/taskManagerContract"
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

func AddTask(context config.Context) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		requestData, err := ioutil.ReadAll(req.Body)
		if err != nil {
			errorHandler.ErrorHandler(context.ErrorLogFile,err)
		}
		data := &contract.Task{}
		err = proto.Unmarshal(requestData,data)
		if err != nil {
			errorHandler.ErrorHandler(context.ErrorLogFile,err)
		}
		newTask := models.Task{}
		newTask.TaskDescription = *data.Task
		newTask.Priority = *data.Priority
		err = newTask.Create(context)
		if err != nil {
			resp := responseGenerator(http.StatusInternalServerError,err.Error())
			dataToBeSend,err :=  proto.Marshal(&resp)
			if err != nil {
				errorHandler.ErrorHandler(context.ErrorLogFile,err)
			}
			res.WriteHeader(http.StatusInternalServerError)
			res.Write(dataToBeSend)
			return 
		}
		res.WriteHeader(http.StatusCreated)
	}
}

func GetTasks(context config.Context) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		data,err := models.Get(context)
		if err != nil {
			resp := responseGenerator(http.StatusInternalServerError,err.Error())
			dataToBeSend,err :=  proto.Marshal(&resp)
			if err != nil {
				errorHandler.ErrorHandler(context.ErrorLogFile,err)
			}
			res.Write(dataToBeSend)
			res.WriteHeader(http.StatusInternalServerError)
		}
		response := contract.Response{}
		response.Data = []byte(data)
		dataToBeSend,err :=  proto.Marshal(&response)
		if err != nil {
			errorHandler.ErrorHandler(context.ErrorLogFile,err)
			return
		}
		res.WriteHeader(http.StatusOK)
		res.Write(dataToBeSend)
	}
}

func DeleteTask(context config.Context) http.HandlerFunc {
	return func(res http.ResponseWriter,req *http.Request) {
		req.ParseForm()
		taskId := strings.Split(req.RequestURI,"/")[2]
		id,err := strconv.Atoi(taskId)
		taskToDelete := models.Task{}
		taskToDelete.TaskId = id
		err = taskToDelete.Delete(context)
		if err != nil {
			resp := responseGenerator(http.StatusInternalServerError,err.Error())
			dataToBeSend,err :=  proto.Marshal(&resp)
			if err != nil {
				errorHandler.ErrorHandler(context.ErrorLogFile,err)
			}
			res.WriteHeader(http.StatusInternalServerError)
			res.Write(dataToBeSend)
			return
		}
		res.WriteHeader(http.StatusAccepted)
	}
}

func UpdateTask(context config.Context)http.HandlerFunc{
	return func(res http.ResponseWriter,req *http.Request) {
		requestData, err := ioutil.ReadAll(req.Body)
		if err != nil {
			errorHandler.ErrorHandler(context.ErrorLogFile,err)
		}
		data := &contract.Task{}
		err = proto.Unmarshal(requestData,data)
		if err != nil {
			errorHandler.ErrorHandler(context.ErrorLogFile,err)
		}
		taskId := strings.Split(req.RequestURI,"/")[2]
		id,err := strconv.Atoi(taskId)

		if err != nil {
			errorHandler.ErrorHandler(context.ErrorLogFile,err)
		}
		taskToUpdate := models.Task{}
		taskToUpdate.TaskId = id
		taskToUpdate.TaskDescription = *data.Task
		taskToUpdate.Priority = *data.Priority
		err = taskToUpdate.Update(context)
		if err != nil {
			resp := responseGenerator(http.StatusInternalServerError,err.Error())
			dataToBeSend,err :=  proto.Marshal(&resp)
			if err != nil {
				errorHandler.ErrorHandler(context.ErrorLogFile,err)
			}
			res.WriteHeader(http.StatusInternalServerError)
			res.Write(dataToBeSend)
			return
		}
		res.WriteHeader(http.StatusCreated)
	}
}

func UploadCsv(context config.Context) http.HandlerFunc{
	return func(res http.ResponseWriter,req *http.Request) {
		requestData, err := ioutil.ReadAll(req.Body)
		if err != nil {
			errorHandler.ErrorHandler(context.ErrorLogFile,err)
		}
		data := &contract.Upload{}
		err = proto.Unmarshal(requestData,data)
		if err != nil {
			errorHandler.ErrorHandler(context.ErrorLogFile,err)
		}
		err = models.AddTaskByCsv(context,string(data.Data))
		if err != nil {
			resp := responseGenerator(http.StatusInternalServerError,err.Error())
			dataToBeSend,err :=  proto.Marshal(&resp)
			if err != nil {
				errorHandler.ErrorHandler(context.ErrorLogFile,err)
			}
			res.Write(dataToBeSend)
			return
		}
		res.WriteHeader(http.StatusOK)
	}
}

func DownloadCsv(context config.Context) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		data,err := models.GetCsv(context)
		if err != nil {
			resp := responseGenerator(http.StatusInternalServerError,err.Error())
			dataToBeSend,err :=  proto.Marshal(&resp)
			if err != nil {
				errorHandler.ErrorHandler(context.ErrorLogFile,err)
			}
			res.WriteHeader(http.StatusInternalServerError)
			res.Write(dataToBeSend)
			return
		}

		response := contract.Response{}
		response.Data = data
		dataToBeSend,err :=  proto.Marshal(&response)
		if err != nil {
			errorHandler.ErrorHandler(context.ErrorLogFile,err)
			return
		}
		res.WriteHeader(http.StatusOK)
		res.Write(dataToBeSend)
	}
}

