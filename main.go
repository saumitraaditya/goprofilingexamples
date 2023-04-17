package main

import (
	"flag"

	"github.com/pkg/profile"
)

const (
	load               = 1000000
	input              = "this is to demo profiling"
	charset            = "abcdefghijklmnopqrstuvwxyz"
	inputlength        = 6
	filepath           = "./out.txt"
	paddingBytesNeeded = 2000
)

func main() {
	profilePtr := flag.String("profile", "none", "one of: cpu, mem, block, trace")
	modePtr := flag.String("mode", "seq", "mode of execution: one of seq, parallel, pool, rawWriter, bufferedWriter, padding, paddingReuse")
	roundPtr := flag.Int("rounds", 8, "number of instances of work")
	poolsizePtr := flag.Int("workers", 8, "number of goroutines in worker pool")
	wordsPtr := flag.Int("wordcount", 10000, "number of random words to generate")
	producerPtr := flag.Int("producers", 10, "number of word generators")

	flag.Parse()
	switch *profilePtr {
	case "cpu":
		defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	case "heap":
		defer profile.Start(profile.MemProfileHeap, profile.ProfilePath(".")).Stop()
	case "trace":
		defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()
	case "block":
		defer profile.Start(profile.BlockProfile, profile.ProfilePath(".")).Stop()
	default:
		// no profiling
	}
	switch *modePtr {
	case "seq", "parallel", "pool":
		doHeavyWork(*modePtr, *roundPtr, *poolsizePtr)
	case "rawWriter":
		writerDemo(*wordsPtr, *producerPtr, false)
	case "bufferedWriter":
		writerDemo(*wordsPtr, *producerPtr, true)
	case "padding":
		demoPadding(*roundPtr, false)
	case "paddingReuse":
		demoPadding(*roundPtr, true)
	}
}
