package fileReaders

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCsvFileReader_Read(t *testing.T) {
	reader := CsvFileReader{}
	reader.FileData = "Buy bottles,High"
	data,_ := reader.Read()
	for _, each := range data {
		assert.Equal(t,"High",each.PRIORITY)
		assert.Equal(t,"Buy bottles",each.TASK)
	}
}

func TestCsvFileReaderForError(t *testing.T) {
	reader := CsvFileReader{}
	reader.FileData = "Buy bottles,igh"
	_,err := reader.Read()
	if err == nil {
		t.Failed()
	}
}
