package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/sofyan48/rll-daemon-new/src/config"

	"github.com/joho/godotenv"
	"github.com/sofyan48/rll-daemon-new/src/routes"
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
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	ConfigEnvironment(*environment)
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
