package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/sofyan48/cimol/src/config"
	"github.com/sofyan48/cimol/src/transmiter"
	"github.com/sofyan48/cimol/src/util/helper/storage"

	"github.com/joho/godotenv"
	"github.com/sofyan48/cimol/src/routes"
)

// ConfigEnvironment |
func ConfigEnvironment(env string) {
	if env == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}

func main() {
	// starting metric storage
	metric := storage.StorageHandler()
	metric.CreateMetricFolder()
	// end metric storage

	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	ConfigEnvironment(*environment)
	wg := &sync.WaitGroup{}
	transmiters := transmiter.GetTransmiter()
	wg.Add(1)
	go transmiters.ConsumerTrans(wg)

	startApp()
}

func startApp() {
	router := config.SetupRouter()
	routes.LoadRouter(router)
	serverHost := os.Getenv("SERVER_ADDRESS")
	serverPort := os.Getenv("SERVER_PORT")
	serverString := fmt.Sprintf("%s:%s", serverHost, serverPort)
	router.Run(serverString)

}
