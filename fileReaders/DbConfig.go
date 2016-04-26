package fileReaders

import (
	"io/ioutil"
	"fmt"
	"os"
	"encoding/json"
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

func ReadJsonFile(fileName string)(JsonObject){
	file, e := ioutil.ReadFile(fileName)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	var jsonType JsonObject
	json.Unmarshal(file, &jsonType)
	return jsonType
}
