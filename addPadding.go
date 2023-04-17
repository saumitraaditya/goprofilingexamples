package main

func genLoad(rounds int, input chan []byte) {
	for i := 0; i < rounds; i++ {
		myBytes := make([]byte, 10, 10)
		input <- myBytes
	}
}

func addPadding(input chan []byte, reuseAllocation bool) {

	paddingBytes := make([]byte, paddingBytesNeeded)
	for {
		select {
		case inputBytes, ok := <-input:
			if !ok {
				break
			}
			if reuseAllocation {
				_ = append(inputBytes, paddingBytes...)
			} else {
				newpaddingBytes := make([]byte, paddingBytesNeeded)
				_ = append(inputBytes, newpaddingBytes...)
			}
		}
	}
}

func demoPadding(rounds int, reuse bool) {
	inputChan := make(chan []byte)
	go addPadding(inputChan, reuse)
	genLoad(rounds, inputChan)
	close(inputChan)
}
