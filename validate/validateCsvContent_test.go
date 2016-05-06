package validate

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestValidateAllEntryForGivenWrongEntry(t *testing.T) {
	wrongEntry := [][]string{{"drinking water", "High"},
		{"coding", "Medium", "Accha"},
		{"Understanding", "Samjhe"},
		{"", "High"},
	}
	allValidator := []Validator{NoOfColumnValidator{}, TaskDescriptionValidator{}, PriorityValidator{}}
	err := ValidateAllEntry(wrongEntry,allValidator)
	expectedErrorMessage := "errors in the following line :\nLine No 2: expected 2 columns, but got 3 columns\nLine No 3,Column No 2: expected priority is High or Medium or Low, but got Samjhe\nLine No 4,Column No 1: Task DescripTion Can't be empty\n"
	assert.Equal(t,expectedErrorMessage,err.Error(),"There will be error")
}

func TestValidateAllEntryForGivenRightEntry(t *testing.T) {
	rightEntry := [][]string{{"drinking water", "High"},{"coding", "Medium"},{"passing test","Low"}}
	allValidator := []Validator{NoOfColumnValidator{}, TaskDescriptionValidator{}, PriorityValidator{}}
	err := ValidateAllEntry(rightEntry,allValidator)
	assert.Equal(t,nil,err,"There will be no error")
}