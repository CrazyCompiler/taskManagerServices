package routers

import (
	"github.com/gorilla/mux"
	"net/http"
	"taskManagerServices/config"
	"taskManagerServices/handlers"
)


func HandleRequests(configObject config.ContextObject) {
	r := mux.NewRouter()
	r.HandleFunc("/update",handlers.UpdateTask(configObject)).Methods("POST")
	r.HandleFunc("/tasks/csv",handlers.UploadCsv(configObject)).Methods("POST")
	r.HandleFunc("/task/{id:[0-9]+}", handlers.DeleteTask(configObject)).Methods("DELETE")
	r.HandleFunc("/tasks", handlers.GetTasks(configObject)).Methods("GET")
	r.HandleFunc("/task", handlers.AddTask(configObject)).Methods("POST")
	http.Handle("/", r)
}