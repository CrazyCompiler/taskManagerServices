package validate

import (
	"errors"
	"strconv"
)

type Validator interface {
	hasError(eachRow []string) bool
	GenerateErrorMessage(lineNo int, eachRow []string) string
}


type NoOfColumnValidator struct  {
}

func (c NoOfColumnValidator) hasError(eachRow []string) bool {
	return len(eachRow)!=2
}

func (c NoOfColumnValidator) GenerateErrorMessage(lineNo int,eachRow []string)string  {
	return "Line No "+strconv.Itoa(lineNo)+": expected 2 columns, but got "+strconv.Itoa(len(eachRow))+" columns\n"
}


type TaskDescriptionValidator struct  {
}

func (c TaskDescriptionValidator) hasError(eachRow []string)bool  {
	return eachRow[0] == ""
}

func (c TaskDescriptionValidator) GenerateErrorMessage(lineNo int,eachRow []string)string  {
	return "Line No "+strconv.Itoa(lineNo)+",Column No 1:"+" Task DescripTion Can't be empty\n"
}



type PriorityValidator struct  {
}

func (c PriorityValidator) hasError(eachRow []string) bool {
	return eachRow[1] != "High" && eachRow[1] != "Medium" && eachRow[1] != "Low"
}

func (c PriorityValidator) GenerateErrorMessage(lineNo int,eachRow []string)string  {
	return "Line No "+strconv.Itoa(lineNo)+",Column No 2:"+" expected priority is High or Medium or Low, but got "+eachRow[1]+"\n"
}


func ValidateAllEntry(allEntry [][]string, allValidators []Validator) error {
	errMsg := ""
	count := 1
	for _, eachEntry := range allEntry {
		for _, eachValidator := range allValidators {
			if eachValidator.hasError(eachEntry) {
				errMsg += eachValidator.GenerateErrorMessage(count, eachEntry)
			}
		}
		count++
	}
	if errMsg!="" {
		return errors.New("errors in the following line :\n"+errMsg)
	}
	return nil
}
