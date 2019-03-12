package main

import (
	"log"
	"math/rand"
	"sync"
)
var(
	total = 0
	mutex = &sync.Mutex{}
	wg sync.WaitGroup
)
func problem1() {

	log.Printf("problem1: started --------------------------------------------")

	//
	// Todo:
	//
	// Quit all go routines after
	// a total of exactly 100 random
	// numbers have been printed.
	//
	// Do not change the 25 in loop!
	//

	wg.Add(1)

	for inx := 0; inx < 10; inx++ {
		go printRandom1(inx)
	}

	//
	// Todo:
	//
	// Remove this quick and dirty sleep
	// against a synchronized wait until all
	// go routines are finished.
	//

	// block until all goroutines have finished.
	wg.Wait()

	log.Printf("problem1: finised --------------------------------------------")
}

func printRandom1(slot int) {

	//
	// Do not change 25 into 10!
	//
	wg.Add(1)
	wg.Done()
	for inx := 0; inx < 25; inx++ {
		mutex.Lock()
		total += 1
		log.Printf("problem1: slot=%03d count=%05d rand=%f", slot, inx, rand.Float32())
		if total==100 {
			wg.Done()
			break
		}
		mutex.Unlock()
	}

}
