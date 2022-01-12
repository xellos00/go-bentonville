package main

import (
	"log"
	"time"

	plugin_grpc "heimdall-plugin/grpc"
	"heimdall-plugin/manager"
)

const (
	servName = "Node manager"
)

func main() {
	log.Println("Start ", servName)

	manager.RunManager()
	err := manager.RunManager().Start()
	if err != nil {
		log.Panic(err)
	}

	plugin_grpc.StartServer()

	for {
		// TODO: Handle SIGTERM, Shutdown gracefully.
		time.Sleep(time.Second * 10)
	}
}
