package fileReaders

import (
	"encoding/csv"
	"strings"
	"taskManagerServices/validate"
)

type TableContent struct {
	TASK     string
	PRIORITY string
}

type Reader interface {
	Read() ([]TableContent,error)
}

type CsvFileReader struct {
	FileData string
}

func (r CsvFileReader) Read() ([]TableContent,error) {
	dataArray := []TableContent{}
	reader := csv.NewReader(strings.NewReader(r.FileData))

	reader.FieldsPerRecord = -1
	rawCsvData, err := reader.ReadAll()

	allValidators := []validate.Validator{validate.NoOfColumnValidator{},
		validate.TaskDescriptionValidator{},
		validate.PriorityValidator{}}

	err = validate.ValidateAllEntry(rawCsvData,allValidators)
	if err != nil {
		return dataArray,err
	}

	for _, each := range rawCsvData {
		entry := TableContent{}
		if(len(each) == 2) {
			entry.TASK = each[0]
			entry.PRIORITY = each[1]
			dataArray = append(dataArray, entry)
		}
	}
	return dataArray,err

}

