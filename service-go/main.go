package servicego

import (
	"fmt"
	"time"
)

var stopChan chan bool

// StartService starts the go service. This process never exit
func StartService() {

	stopChan = make(chan bool, 1)
	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-stopChan:
			fmt.Println("Go Service Shutdown")
		case t := <-ticker.C:
			fmt.Println("Go Service: Tick at", t)
		}
	}
}

// StopService stops the go service
func StopService() {
	stopChan <- true
}
