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
			fmt.Println("go shutdown")
			return
		case t := <-ticker.C:
			fmt.Println("go:", t)
		}
	}
}

// StopService stops the go service
func StopService() {
	stopChan <- true
}
