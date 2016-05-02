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
	r.HandleFunc("/tasks/{id:[0-9]+}",handlers.UpdateTask(context)).Methods("PATCH")
	r.HandleFunc("/tasks/csv",handlers.UploadCsv(context)).Methods("POST")
	r.HandleFunc("/task/{id:[0-9]+}", handlers.DeleteTask(context)).Methods("DELETE")
	r.HandleFunc("/tasks", handlers.GetTasks(context)).Methods("GET")
	r.HandleFunc("/tasks", handlers.AddTask(context)).Methods("POST")
	http.Handle("/", r)
}