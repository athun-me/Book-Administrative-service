package main

import (
	"log"

	"githum.com/athunlal/bookNowAdmin-svc/pkg/config"
	"githum.com/athunlal/bookNowAdmin-svc/pkg/di"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed to load the Congiguration File: ", err)
		return
	}
	server, err := di.InitApi(cfg)
	if err != nil {
		log.Fatalln("Error in initializing the api", err)
	}
	server.Start()
}
