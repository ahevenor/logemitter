package main

import (
	"log"
	"time"
)

func main() {
	for range time.Tick(1 * time.Second) {
		log.Println("logging a message")
	}
}
