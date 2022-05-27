package main

import (
	"embed"
	"fmt"
	"github.com/genleel/go-sync/server"
	"os"
	"os/signal"
	"sync"
)

//go:embed server/frontend/dist/*
var FS embed.FS

func recoverFromError() {
	if r := recover(); r != nil {
		fmt.Println("Recovering from panic:", r)
	}
}

func main() {
	var endWriter sync.WaitGroup
	endWriter.Add(1)

	go server.Run()

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	select {
	case <-signalChannel:
		endWriter.Done()
	}
	endWriter.Wait()

}
