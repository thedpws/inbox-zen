package main

import (
	"fmt"
	"time"
)

func doOnTimer(freqSeconds int, f func()) (stopFunc func()) {

	t := time.Duration(freqSeconds) * time.Second;
    ticker := time.NewTicker(t)
    stopChan := make(chan bool)

    go func () {
		for {
			select {
			case <-ticker.C:
				f()
			case <-stopChan:
				return
			}
		}
    }()

    return func() {
        stopChan <- true
    }
}	


func main() {
	stop := doOnTimer(2, func(){
		fmt.Println("sfasdf")
	})

	time.Sleep(10 * time.Second)
	stop()
}
