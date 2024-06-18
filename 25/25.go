package main

import "time"

func main() {
	Sleep(10 * time.Second)
}

func Sleep(t time.Duration) {
	<-time.After(t)
}
