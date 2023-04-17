package main

import (
	"bufio"
	"io"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

func randomString() []byte {
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, inputlength)
	for i := range b {
		b[i] = charset[seed.Intn(len(charset))]
	}
	return b
}

func setupConsumer(file *os.File, inputchan chan []byte) {
	go func(writer io.Writer) {
		for {
			select {
			case input, ok := <-inputchan:
				if !ok {
					break
				}
				_, err := writer.Write(input)
				if err != nil {
					log.Fatal(err)
				}
			}
		}

	}(file)
}

func setupBufferedConsumer(file *os.File, inputchan chan []byte) {
	bw := bufio.NewWriterSize(file, 4096)
	go func(writer io.Writer) {
		for {
			select {
			case input, ok := <-inputchan:
				if !ok {
					break
				}
				_, err := writer.Write(input)
				if err != nil {
					log.Fatal(err)
				}
			}
		}

	}(bw)
}

func writerDemo(wordcount int, producers int, buffered bool) {
	inputchan := make(chan []byte, 1)
	wordsPerProducer := wordcount / producers
	f, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if buffered {
		setupBufferedConsumer(f, inputchan)
	} else {
		setupConsumer(f, inputchan)
	}
	var wg sync.WaitGroup
	wg.Add(producers)
	for i := 0; i < producers; i++ {
		go func() {
			defer wg.Done()
			for w := 0; w < wordsPerProducer; w++ {
				genstring := randomString()
				inputchan <- genstring
			}
		}()
	}
	wg.Wait()
	close(inputchan)

}
