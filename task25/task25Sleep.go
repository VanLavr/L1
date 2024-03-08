package main

import (
	"fmt"
	"syscall"
	"time"
)

func Sleep(duration time.Duration) {
	<-time.NewTimer(duration * time.Second).C // NewTimer creates a new Timer that will send the current time on its channel after at least duration d.
} // channel is C

func AnotherSleep(duration time.Duration) {
	ts := syscall.NsecToTimespec(int64(duration)) // converting time.Duration to syscall.Timespec struct
	syscall.Nanosleep(&ts, nil)                   // first arg is for sleeping time, second is for remainig time if any secs remained (they will be stored there)
}

func main() {
	fmt.Println("hello world")
	Sleep(2)
	fmt.Println("hello world")
	AnotherSleep(time.Second * 3)
	fmt.Println("hello world")
}
