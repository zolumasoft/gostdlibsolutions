package main

import "log"

func main() {
	// Panic is recovered to stop program from crashing
	defer func() {
		if r := recover(); r != nil {
			log.Printf("boom() failed: %v\n", r)
		}
	}()
	log.Println("Hello, Gopher!")
	boom()
}

func boom() {
	panic("BOOM!")
}
