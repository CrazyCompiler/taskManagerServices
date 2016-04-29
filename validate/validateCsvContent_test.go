package validate

import "testing"

func TestValidateAllEntry(t *testing.T) {
	sampleEntries := [][]string{{"drinking","High"},
		{"coding","Medium"},
		{""}}
	err := ValidateAllEntry(sampleEntries)

	if err != nil {
		t.Errorf("expected err will be nil but error is %s",err.Error())
	}

	sampleEntries = [][]string{{"drinking","High"},
		{"","Medium"},
		{"sleeping","Hello"},
		{"coocking","Low"},
		{""}}

	err = ValidateAllEntry(sampleEntries)

	if err.Error() != "Errors in the following lines 2,3"{
		t.Errorf("err suppose to be Errors in the following lines 2,3 but it is %s",err.Error())
	}

}