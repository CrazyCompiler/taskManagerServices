package routers

import (
	"github.com/gorilla/mux"
	"net/http"
	"taskManagerServices/config"
	"taskManagerServices/handlers"
)


func HandleRequests(context config.Context) {
	r := mux.NewRouter()
	r.HandleFunc("/tasks/download/csv",handlers.DownloadCsv(context)).Methods("GET")
	r.HandleFunc("/update",handlers.UpdateTask(context)).Methods("POST")
	r.HandleFunc("/tasks/csv",handlers.UploadCsv(context)).Methods("POST")
	r.HandleFunc("/task/{id:[0-9]+}", handlers.DeleteTask(context)).Methods("DELETE")
	r.HandleFunc("/tasks", handlers.GetTasks(context)).Methods("GET")
	r.HandleFunc("/task", handlers.AddTask(context)).Methods("POST")
	http.Handle("/", r)
}