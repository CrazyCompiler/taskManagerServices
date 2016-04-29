package errorHandler

import (
	"testing"
	"os"
	"errors"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strings"
)

func TestErrorHandler(t *testing.T) {
	errorLogFilePath := "./errorLogForTest"
	errorFile, err := os.OpenFile(errorLogFilePath, os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer errorFile.Close()
	errorToBeWritten := "file not found"
	errorToTest := errors.New(errorToBeWritten)
	ErrorHandler(errorFile,errorToTest)
	data,err := ioutil.ReadFile(errorLogFilePath)
	originalArray := strings.Split(string(data)," ")
	original := strings.Join(originalArray[2:]," ")
	expected := errorToBeWritten+"\n"
	assert.Equal(t, expected,original)
}
