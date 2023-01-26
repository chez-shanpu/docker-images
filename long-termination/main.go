package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log.Println("start")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	s := <-sigs
	switch s {
	case syscall.SIGINT:
		fallthrough
	case syscall.SIGTERM:
		log.Println("A signal is catched")
		time.Sleep(5 * time.Minute)
		log.Println("os exit")
		os.Exit(0)
	}

	for {
		log.Println("for loop")
		time.Sleep(5 * time.Second)
	}
}
