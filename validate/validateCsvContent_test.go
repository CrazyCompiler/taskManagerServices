package validate

import "testing"

func TestIsValidPriority(t *testing.T) {
	permission := IsValidPriority("High")
	if permission != true{
		t.Errorf("Permission should be true but it is false")
	}

	permission = IsValidPriority("India")
	if permission != false{
		t.Errorf("Permission should be false but it is true")
	}
}
