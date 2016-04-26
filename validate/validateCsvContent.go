package validate

import (
	"errors"
	"strconv"
)

func ValidateAllEntry(allEntry [][]string)  error{
	errorLineNumbers := []int{}
	count := 1;
	for _, each := range allEntry{
		if isValidNoOfColumn(each)==false|| isNotEmptyTheTaskDescription(each[0])==false || isValidPriority(each[1])==false {
			errorLineNumbers = append(errorLineNumbers,count)
		}
		count++
	}
	if len(errorLineNumbers)>1{
		return errors.New("Errors in the following lines"+linesInString(errorLineNumbers))
	}
	return nil
}

func linesInString(linesNumber []int) string  {
	lines := " "+strconv.Itoa(linesNumber[0])
	for  i := 1;i< len(linesNumber)-1;i++{
		lines = lines +"," +strconv.Itoa(linesNumber[i])
	}
	return lines
}

func isNotEmptyTheTaskDescription(task string) bool {
	return task != ""
}

func isValidNoOfColumn(eachEntry []string) bool {
	return len(eachEntry)==2
}
func isValidPriority(priority string) bool{
	return priority=="High" || priority == "Medium" || priority == "Low"
}

