package errorHandler

import (
	"os"
	"time"
	"fmt"
	"log"
)

func ErrorHandler(errorFile *os.File,errorOccurred error) {
	errorBody := time.ANSIC + " : "  +errorOccurred.Error()+ "\n"
	fmt.Print(errorBody)
	log.SetOutput(errorFile)
}

