package main

import (
	"fmt"
	"time"
)

type ZenChannel interface {
	Open() chan ZenMessage
	Close()
}

type ZenMessage interface {
	Content() string
	Reply(string)
}

// type EmailChannel


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
	var stops = []func(){}

	stops = append(stops, doOnTimer(2, func(){
		fmt.Println("sfasdf")
	}))

	time.Sleep(10 * time.Second)
	for _,stop := range stops {
		stop()
	}
}
