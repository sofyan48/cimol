package logging

import (
	"fmt"
	"log"
	"os"

	entity "github.com/sofyan48/cimol/src/entity/http/v1"
)

// CreateLog ...
func CreateLog(logging entity.Logging) {
	if os.Getenv("APP_LOG") == "local" {
		f, err := os.OpenFile("./log/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()
		log.SetOutput(f)
		log.Println("This is a test log entry")
	} else {
		fmt.Println("COMING SOON")
		log.Println()
	}
}
