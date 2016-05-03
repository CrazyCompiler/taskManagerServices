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
	"encoding/csv"
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
		taskId := strings.Split(req.RequestURI,"/")[2]
		task,err := strconv.Atoi(taskId)

		if err != nil {
			errorHandler.ErrorHandler(context.ErrorLogFile,err)
		}
		err = models.Update(context, task, *data.Task,*data.Priority)
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

type Wr struct {
	buf []byte
}

func (w *Wr) Write(b []byte) (int, error) {
	w.buf = append(w.buf, b...)
	return len(w.buf), nil
}

func DownloadCsv(context config.Context) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		dbData,err := models.GetCsv(context)
		wrt := &Wr{}
		w := csv.NewWriter(wrt)
		w.WriteAll(dbData)
		w.Flush()
		err = w.Error()

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
		response.Data = wrt.buf
		dataToBeSend,err :=  proto.Marshal(&response)
		if err != nil {
			errorHandler.ErrorHandler(context.ErrorLogFile,err)
			return
		}
		res.WriteHeader(http.StatusOK)
		res.Write(dataToBeSend)
	}
}

