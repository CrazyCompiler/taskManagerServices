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
	"github.com/taskManagerContract"
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
		err = models.Add(context, *data.Task, *data.Priority)
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
		response := contract.GetTasks{}
		response.Bytedata = data
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
		task,err := strconv.Atoi(taskId)
		err = models.Delete(context,task)
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
		err = models.Update(context, *data.TaskId, *data.Task,*data.Priority)
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
		data := &contract.UploadCsvData{}
		err = proto.Unmarshal(requestData,data)
		if err != nil {
			errorHandler.ErrorHandler(context.ErrorLogFile,err)
		}
		err = models.AddTaskByCsv(context,string(*data.CsvData))
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

		response := contract.GetTasks{}
		response.Bytedata = data
		dataToBeSend,err :=  proto.Marshal(&response)
		if err != nil {
			errorHandler.ErrorHandler(context.ErrorLogFile,err)
			return
		}
		res.WriteHeader(http.StatusOK)
		res.Write(dataToBeSend)
	}
}

