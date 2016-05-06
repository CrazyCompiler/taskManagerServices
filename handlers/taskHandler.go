package handlers

import (
	"io/ioutil"
	"strings"
	"strconv"
	"taskManagerServices/config"
	"net/http"
	"taskManagerServices/errorHandler"
	"github.com/golang/protobuf/proto"
	"github.com/CrazyCompiler/taskManagerContract"
	"encoding/csv"
	"taskManagerServices/models"
	"taskManagerServices/fileReaders"
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
		req.ParseForm()
		userId := strings.Split(req.RequestURI,"/")[2]

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
		err = newTask.Create(context,userId)
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
		userId := strings.Split(req.RequestURI,"/")[2]
		data,err := models.Get(context,userId)
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
		taskId := strings.Split(req.RequestURI,"/")[3]
		id,err := strconv.Atoi(taskId)
		userIdProvided := strings.Split(req.RequestURI,"/")[1]

		taskToDelete := models.Task{}
		taskToDelete.TaskId = id

		err = taskToDelete.Delete(context,userIdProvided)
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
		taskId := strings.Split(req.RequestURI,"/")[3]
		id,err := strconv.Atoi(taskId)

		if err != nil {
			errorHandler.ErrorHandler(context.ErrorLogFile,err)
		}
		taskToUpdate := models.Task{}
		taskToUpdate.TaskId = id
		taskToUpdate.TaskDescription = *data.Task
		taskToUpdate.Priority = *data.Priority

		userIdProvided := strings.Split(req.RequestURI,"/")[1]

		err = taskToUpdate.Update(context,userIdProvided)
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

		userIdProvided := strings.Split(req.RequestURI,"/")[3]
		csvReader := fileReaders.CsvFileReader{
			FileData:string(data.Data),
		}
		err = models.AddTaskByCsv(context,userIdProvided,csvReader)
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
		userIdProvided := strings.Split(req.RequestURI,"/")[3]
		dbData,err := models.GetCsv(context,userIdProvided)
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

