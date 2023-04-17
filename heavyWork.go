package main

import (
	"crypto/sha256"
	"sync"
)

func heavyLift(input string) bool {
	for i := 0; i < load; i++ {
		_ = sha256.Sum256([]byte(input))
	}
	return true
}

func setuppool(inputchan chan string, resultchan chan bool, poolsize int) {
	for i := 0; i < poolsize; i++ {
		go func() {
			for {
				select {
				case myinput := <-inputchan:
					done := heavyLift(myinput)
					resultchan <- done
				}
			}
		}()
	}
}

func doHeavyWork(mode string, rounds int, poolsize int) {

	var wg sync.WaitGroup
	switch mode {
	case "seq":
		for i := 0; i < rounds; i++ {
			heavyLift(input)
		}
	case "parallel":
		for i := 0; i < rounds; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				heavyLift(input)
			}()
		}
		wg.Wait()
	case "pool":
		inputchan := make(chan string)
		resultchan := make(chan bool)
		setuppool(inputchan, resultchan, poolsize)
		//setup reader
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < rounds; i++ {
				_ = <-resultchan
			}
		}()
		//start sending work
		for i := 0; i < rounds; i++ {
			inputchan <- input
		}
		wg.Wait()
	}

}
