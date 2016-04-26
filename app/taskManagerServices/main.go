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
	configObject := config.ContextObject{}
	errorLogFilePath := "errorLog"
	errorFile, err := os.OpenFile(errorLogFilePath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer errorFile.Close()

	configObject.ErrorLogFile = errorFile

	dbConfigFilePath := "dbConfigFile"
	if len(os.Args) > 1 {
		dbConfigFilePath = os.Args[1]
	}
	dbConfigDataJson := fileReaders.ReadJsonFile(dbConfigFilePath)
	dbinfo := database.CreateDbInfo(dbConfigDataJson)

	configObject.Db, err = sql.Open("postgres", dbinfo)

	if err != nil {
		errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
	}

	configObject.Db.Ping()

	if err != nil {
		errorHandler.ErrorHandler(configObject.ErrorLogFile, err)
	}

	defer configObject.Db.Close()
	routers.HandleRequests(configObject)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("their was error ", err)
	}

}
