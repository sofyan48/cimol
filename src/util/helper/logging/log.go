package logging

import (
	"encoding/json"
	"log"
	"os"
	"sync"
	"time"

	entity "github.com/sofyan48/cimol/src/entity/http/v1"
)

// Log ...
type Log struct {
}

// LogHandler ..
func LogHandler() *Log {
	return &Log{}
}

// LogInterface ...
type LogInterface interface {
	Write(name string, desc interface{})
}

func envLog() bool {
	env := os.Getenv("APP_LOG")
	return env == "local"
}

// CreateLog ...
func logStorage() *log.Logger {
	logFile, err := os.OpenFile("./log/cimol.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	// defer logFile.Close()
	logger := log.New(logFile, "", log.LstdFlags)
	return logger
}

func logsLocal(name string, desc interface{}, wg *sync.WaitGroup) {
	logs := logStorage()
	logging := &entity.Logging{}
	logging.Name = name
	logging.Description = desc
	logging.TimeAt = time.Now()
	data, err := json.Marshal(logging)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	logs.Println(string(data))
	wg.Done()
}

// s3logs ...
func s3logs(name string, desc interface{}) {
	logging := &entity.Logging{}
	logging.Name = name
	logging.Description = desc
	logging.TimeAt = time.Now()
	data, err := json.Marshal(logging)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.Println(string(data))
}

// Write ...
func (lg *Log) Write(name string, desc interface{}) {
	wg := &sync.WaitGroup{}
	if envLog() {
		wg.Add(1)
		logsLocal(name, desc, wg)
		return
	}
	wg.Add(1)
	go s3logs(name, desc)
}
