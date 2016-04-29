package errorHandler

import (
	"os"
	"fmt"
	"log"
)

func ErrorHandler(errorFile *os.File,errorOccurred error) {
	fmt.Print(errorOccurred.Error())
	log.SetOutput(errorFile)
	log.Println(errorOccurred.Error())
}

