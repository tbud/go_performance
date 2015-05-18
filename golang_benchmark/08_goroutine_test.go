package golang_benchmark

import (
	"sync"
	"testing"
)

func worker(linkChan chan string, wg *sync.WaitGroup) {
	// Decreasing internal counter for wait-group as soon as goroutine finishes
	defer wg.Done()

	for url := range linkChan {
		// Analyze value and do the job here
	}
}

func Benchmark08GoRoutine(b *testing.B) {
	for i := 0; i < b.N; i++ {

		lCh := make(chan string)
		wg := new(sync.WaitGroup)

		// Adding routines to workgroup and running then
		for i := 0; i < 250; i++ {
			wg.Add(1)
			go worker(lCh, wg)
		}

		// Processing all links by spreading them to `free` goroutines
		for _, link := range yourLinksSlice {
			lCh <- link
		}

		// Closing channel (waiting in goroutines won't continue any more)
		close(lCh)

		// Waiting for all goroutines to finish (otherwise they die as main routine dies)
		wg.Wait()
	}
}
