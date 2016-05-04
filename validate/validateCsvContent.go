package validate

import (
	"errors"
	"strconv"
)

type Validator interface  {
	IsValid(EachRow []string) bool
}

type ValidNoOfColumn struct  {
}

func (c ValidNoOfColumn)IsValid(EachRow []string) bool {
	return len(EachRow)==2
}

type TaskDescriptionChecker struct  {
}

func (c TaskDescriptionChecker)IsValid(EachRow []string)bool  {
	return EachRow[0] != ""
}

type PriorityChecker struct  {
}

func (c PriorityChecker)IsValid(EachRow []string) bool {
	return EachRow[1]=="High" || EachRow[1] == "Medium" || EachRow[1] == "Low"
}

func ValidateAllEntry(allEntry [][]string, allValidators []Validator) error {
	s := ""
	count := 1
	for _,eachEntry := range allEntry {
		for _,eachValidator :=range allValidators{
			if eachValidator.IsValid(eachEntry)==false {
				s =s + strconv.Itoa(count)+" "
			}
		}
		count++
	}
	if s!="" {
		return errors.New("errors in the following line : "+s)
	}
	return nil
}
