package fileReaders

import (
	"testing"
	"taskManagerServices/config"
	"os"
	"github.com/stretchr/testify/assert"
)

func generateContextObject()(config.Context){
	context := config.Context{}
	errorLogFilePath := "../errorLog"
	errorFile, err := os.OpenFile(errorLogFilePath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer errorFile.Close()
	context.ErrorLogFile = errorFile
	return context
}

func TestReadJsonFileGivesErrorForWrontFile(t *testing.T) {
	context := generateContextObject()
	wrongFile := "hello"
	_,err := ReadJsonFile(wrongFile,context)
	if err == nil {
		t.Failed()
	}

}

func TestReadJsonFileGivesJsonData(t *testing.T) {
	context := generateContextObject()
	wrongFile := "./dbConfigForTests"
	jsonData,_ := ReadJsonFile(wrongFile,context)

	assert.Equal(t,"postgres",jsonData.DB_NAME)
	assert.Equal(t,"postgres",jsonData.DB_PASSWORD)
	assert.Equal(t,"postgres",jsonData.DB_USER)
	assert.Equal(t,"todoMaker",jsonData.DB_SCHEMA)
}

func TestJsonObject_IsInOrder(t *testing.T) {
	context := generateContextObject()
	wrongFile := "./dbConfigForTests"
	jsonData, _ := ReadJsonFile(wrongFile, context)
	assert.True(t,jsonData.IsInOrder())
}
