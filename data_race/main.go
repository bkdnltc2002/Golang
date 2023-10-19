package main

import (
	"fmt"
	"os"
	"time"
	"sync"
)
 var mutex sync.Mutex

type Watchdog struct{ last int64 }

func (w *Watchdog) KeepAlive() {
	mutex.Lock()
	defer mutex.Unlock()
	w.last = time.Now().UnixNano() //first conficting access
}

func (w *Watchdog) Start() {
	go func() {
		for {
			time.Sleep(time.Second)
			var curtime int64
			mutex.Lock()
			curtime = w.last
			mutex.Unlock()
			//Second conflicting access
			if curtime < time.Now().Add(-5*time.Second).UnixNano() {
				fmt.Println("No keepalives for 5 seconds. Dying.")
				os.Exit(1)
			} else {
				fmt.Println("Still Alive")
			}
		}
	}()
}

func main() {
	prog := Watchdog{}
	prog.Start()

	// Launch KeepAlive goroutines in a loop
	prog.KeepAlive()
	time.Sleep(5*time.Second)

	forever := make(chan int)
	<-forever
}
