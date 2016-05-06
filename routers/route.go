package routers

import (
	"github.com/gorilla/mux"
	"net/http"
	"taskManagerServices/config"
	"taskManagerServices/handlers"
)


func HandleRequests(context config.Context) {
	r := mux.NewRouter()
	r.HandleFunc("/tasks/csv/{id:[0-9]+}",handlers.DownloadCsv(context)).Methods("GET")
	r.HandleFunc("/{id:[0-9]+}/tasks/{id:[0-9]+}",handlers.UpdateTask(context)).Methods("PATCH")
	r.HandleFunc("/tasks/csv/{id:[0-9]+}",handlers.UploadCsv(context)).Methods("POST")
	r.HandleFunc("/{id:[0-9]+}/tasks/{id:[0-9]+}", handlers.DeleteTask(context)).Methods("DELETE")
	r.HandleFunc("/tasks/{id:[0-9]+}", handlers.GetTasks(context)).Methods("GET")
	r.HandleFunc("/tasks/{id:[0-9]+}", handlers.AddTask(context)).Methods("POST")
	r.HandleFunc("/hello", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("hello"))
	}).Methods("GET")

	http.Handle("/", r)
}