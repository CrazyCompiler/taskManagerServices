package main

import (
	"os"
	"database/sql"
	"fmt"
	"taskManagerServices/config"
	"taskManagerServices/fileReaders"
	"taskManagerServices/database"
	"taskManagerServices/errorHandler"
	"taskManagerServices/routers"
	"net/http"
	_ "github.com/lib/pq"
)

func main() {
	context := config.Context{}
	errorLogFilePath := "errorLog"
	errorFile, err := os.OpenFile(errorLogFilePath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer errorFile.Close()

	context.ErrorLogFile = errorFile

	dbConfigFilePath := "dbConfigFile"
	if len(os.Args) > 1 {
		dbConfigFilePath = os.Args[1]
	}
	dbConfigDataJson,err := fileReaders.ReadJsonFile(dbConfigFilePath,context)
	if err != nil {
		os.Exit(1)
	}
	dbInfo := database.CreateDbInfo(dbConfigDataJson)

	context.Db, err = sql.Open("postgres", dbInfo)

	if err != nil {
		errorHandler.ErrorHandler(context.ErrorLogFile,err)
	}

	context.Db.Ping()

	if err != nil {
		errorHandler.ErrorHandler(context.ErrorLogFile, err)
	}

	defer context.Db.Close()
	routers.HandleRequests(context)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("their was error ", err)
	}

}
