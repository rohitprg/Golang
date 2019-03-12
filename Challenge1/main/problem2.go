package main

import (
	"log"
	"time"
	"math/rand"
	"sync"
)
//Create a new instance of a sync.WaitGroup
var waitgrp sync.WaitGroup
var rateLimit = time.Tick(1 * time.Second)

func problem2() {

	log.Printf("problem2: started --------------------------------------------")

	//
	// Todo:
	//
	// Throttle all go subroutines in a way,
	// that every one second one random number
	// is printed.
	//
	<-rateLimit
	rand.Seed(time.Now().UTC().UnixNano())

	for inx := 0; inx < 10; inx++ {
		//set the number of goroutines to wait for.
		waitgrp.Add(1)
		go printRandom2(inx)
	}

	//
	// Todo:
	//
	// Remove this quick and dirty sleep
	// against a synchronized wait until all
	// go routines are finished.
	//
	// Same as problem1...
	//

	// block until all goroutines have finished.
	waitgrp.Wait()

	log.Printf("problem2: finished -------------------------------------------")
}

func printRandom2(slot int) {

	//each of the printrandom2 goroutines runs and calls Done when finished.
	defer waitgrp.Done()
	for inx := 0; inx < 10; inx++ {
		<-rateLimit
		log.Printf("problem2: slot=%03d count=%05d rand=%f", slot, inx, rand.Float32())
	}
}