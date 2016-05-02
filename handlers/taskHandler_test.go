package handlers

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/DATA-DOG/go-sqlmock"
	"taskManagerServices/config"
	"os"
)

const errorFileName string = "dummyErrorLogForTaskHandlerTest"

func TestGetTasks(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	configObject := config.Context{}
	errorFile, err := os.OpenFile(errorFileName, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer errorFile.Close()

	configObject.ErrorLogFile = errorFile
	configObject.Db = db
	getHandler := GetTasks(configObject)
	req, _ := http.NewRequest("GET", "/getAllTasks", nil)
	w := httptest.NewRecorder()
	getHandler.ServeHTTP(w, req)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
	if w.Code != http.StatusOK {
		t.Errorf("Home page didn't return %v", http.StatusOK)
	}
}








