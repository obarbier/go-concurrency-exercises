package counting

import (
	"math"
	"math/rand"
	"runtime"
	"time"
	"sync"
	"sync/atomic"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenerateNumbers - random number generation
func GenerateNumbers(max int) []int {
	rand.Seed(time.Now().UnixNano())
	numbers := make([]int, max)
	for i := 0; i < max; i++ {
		numbers[i] = rand.Intn(10)
	}
	return numbers
}

// Add - sequential code to add numbers
func Add(numbers []int) int64 {
	var sum int64
	for _, n := range numbers {
		sum += int64(n)
	}
	return sum
}

//TODO: complete the concurrent version of add function.

// AddConcurrent - concurrent code to add numbers
func AddConcurrent(numbers []int) int64 {
	var sum int64
	// Utilize all mcores on machine
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)
	// Divide the input into parts
	dividedGroup := splitSlice(numbers, numCPU)
	var wg sync.WaitGroup
	wg.Add(numCPU)
	// Run computation for each part in seperate goroutine.
	for _, group := range dividedGroup {
		go func( nums []int){
			var partSum int64
			defer wg.Done()
			for _, n := range numbers {
				partSum += int64(n)
			}
			atomic.AddInt64(&sum, partSum)
		}(group)
	}
	// Add part sum to cummulative sum
	wg.Wait()
	return sum
}

func splitSlice(s []int, chunks int) [][]int {
	var newSlice = [][]int{}
	var length int
	if len(s)%2 == 0 {
		length = int(math.Ceil(float64(len(s)) / float64(chunks)))
	} else {
		length = int(math.Floor(float64(len(s)) / float64(chunks)))
	}
	var end int
	for i := 0; i < len(s) -1; i += length {
		end = i + length
		if end >= len(s) -1 {
			newSlice = append(newSlice, s[i:])

		} else {
			newSlice = append(newSlice, s[i:end])
		}

	}
	return newSlice
}
