package main

import (
	"fmt"
	"math/rand"
	"slices"
	"time"

	"github.com/shawnsmithdev/zermelo/v2"
)

func main() {
	postleitzahlen := make([]int, 1000000)
	postleitzahlen2 := make([]int, 1000000)

	for i := 0; i < len(postleitzahlen); i++ {
		postleitzahlen[i] = rand.Intn(98999) + 1000
	}
	copy(postleitzahlen2, postleitzahlen)
	start1 := time.Now()
	slices.Sort(postleitzahlen)

	quickSort(postleitzahlen)
	end1 := time.Now()
	start2 := time.Now()
	zermelo.Sort(postleitzahlen2) // Radix Sort

	end2 := time.Now()
	start3 := time.Now()
	quickSort(postleitzahlen)

	end3 := time.Now()
	elapsed1 := end1.Sub(start1)
	elapsed2 := end2.Sub(start2)
	elapsed3 := end3.Sub(start3)
	fmt.Println("Zeit für das Slice sortieren:", elapsed1, ", Elemente sortiert:", len(postleitzahlen))
	fmt.Println("Zeit für das Radix sortieren:", elapsed2, ", Elemente sortiert:", len(postleitzahlen2))
	fmt.Println("Zeit für das Paralell sortieren:", elapsed3, ", Elemente sortiert:", len(postleitzahlen2))

}

func quickSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	left, right := 0, len(arr)-1

	// Use a random element as the pivot
	pivotIndex := rand.Int() % len(arr)
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	// Parallelize the sorting of the two partitions
	go quickSort(arr[:left])
	go quickSort(arr[left+1:])

	// Waiting for the goroutines to finish
	if left > 0 {
		<-time.After(1 * time.Nanosecond)
	}
}
