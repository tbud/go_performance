package golang_benchmark

import (
	"sync"
	"testing"
)

func Benchmark08Normal(b *testing.B) {
	sum := 0
	for i := 0; i < b.N; i++ {
		sum += i
	}
	println(sum)
}

func workerCreateGotine(iCh int, wg *sync.WaitGroup) {
	// Decreasing internal counter for wait-group as soon as goroutine finishes
	defer wg.Done()

	globInt.Lock()
	globInt.sum += iCh
	globInt.Unlock()
}

func Benchmark08CreateGotine(b *testing.B) {
	globInt.sum = 0
	wg := new(sync.WaitGroup)

	// Adding routines to workgroup and running then
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go workerCreateGotine(i, wg)
	}

	// Waiting for all goroutines to finish (otherwise they die as main routine dies)
	wg.Wait()

	println(globInt.sum)
}

func workerReturnChan(iCh chan int, wg *sync.WaitGroup, ret chan int) {
	// Decreasing internal counter for wait-group as soon as goroutine finishes
	defer wg.Done()

	sum := 0

	for num := range iCh {
		sum += num
	}

	ret <- sum
}

var GoRoutineNum = 16

func Benchmark08ReturnWithChan(b *testing.B) {
	iCh := make(chan int)
	wg := new(sync.WaitGroup)
	iRet := make(chan int, GoRoutineNum)

	// Adding routines to workgroup and running then
	for i := 0; i < GoRoutineNum; i++ {
		wg.Add(1)
		go workerReturnChan(iCh, wg, iRet)
	}

	// Processing all links by spreading them to `free` goroutines
	for i := 0; i < b.N; i++ {
		iCh <- i
	}

	// Closing channel (waiting in goroutines won't continue any more)
	close(iCh)

	// Waiting for all goroutines to finish (otherwise they die as main routine dies)
	wg.Wait()

	sum := 0
	for i := 0; i < GoRoutineNum; i++ {
		sum += <-iRet
	}

	println(sum)
}

var globInt = struct {
	sync.Mutex
	sum int
}{}

func workerGlobInt(iCh chan int, wg *sync.WaitGroup) {
	// Decreasing internal counter for wait-group as soon as goroutine finishes
	defer wg.Done()

	sum := 0

	for num := range iCh {
		sum += num
	}

	globInt.Lock()
	globInt.sum += sum
	globInt.Unlock()
}

func Benchmark08ReturnWithGlob(b *testing.B) {
	globInt.sum = 0
	iCh := make(chan int)
	wg := new(sync.WaitGroup)

	// Adding routines to workgroup and running then
	for i := 0; i < GoRoutineNum; i++ {
		wg.Add(1)
		go workerGlobInt(iCh, wg)
	}

	// Processing all links by spreading them to `free` goroutines
	for i := 0; i < b.N; i++ {
		iCh <- i
	}

	// Closing channel (waiting in goroutines won't continue any more)
	close(iCh)

	// Waiting for all goroutines to finish (otherwise they die as main routine dies)
	wg.Wait()

	println(globInt.sum)
}
