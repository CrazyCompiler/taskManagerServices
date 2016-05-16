package routers

import (
	"github.com/gorilla/mux"
	"net/http"
	"taskManagerServices/config"
	"taskManagerServices/handlers"
	"io"
	"golang.org/x/net/http2"
	"log"
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
		io.WriteString(res, "hello, world!\n")
	}).Methods("GET")

	srv := &http.Server{
		Addr:    ":8080", // Normally ":443"
		Handler: r,
	}
	http2.ConfigureServer(srv, &http2.Server{})
	log.Fatal(srv.ListenAndServeTLS("server.cert", "server.key"))

	//http.Handle("/", r)
}