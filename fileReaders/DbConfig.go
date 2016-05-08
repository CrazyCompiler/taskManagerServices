package fileReaders

import (
	"io/ioutil"
	"encoding/json"
	"taskManagerServices/config"
	"taskManagerServices/errorHandler"
)

type JsonObject struct {
	DB_USER string
	DB_PASSWORD string
	DB_NAME string
	DB_SCHEMA string
}

func (db *JsonObject) IsInOrder() bool{
	return db.DB_NAME != "" && db.DB_USER != "" && db.DB_SCHEMA != "" && db.DB_PASSWORD != ""
}

func ReadJsonFile(fileName string,contextObject config.Context)(JsonObject,error){
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		errorHandler.ErrorHandler(contextObject.ErrorLogFile,err)
		return JsonObject{}, err
	}
	var jsonType JsonObject
	json.Unmarshal(file, &jsonType)
	return jsonType,nil
}
