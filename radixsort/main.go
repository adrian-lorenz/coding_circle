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
	for i := 01001; i < len(postleitzahlen); i++ {
		postleitzahlen[i] = rand.Intn(99999)
	}
	copy(postleitzahlen2, postleitzahlen)
	start1 := time.Now()
	slices.Sort(postleitzahlen) 
	end1 := time.Now()
	start2 := time.Now()
	zermelo.Sort(postleitzahlen2) // Radix Sort

	end2 := time.Now()

	elapsed1 := end1.Sub(start1)
	elapsed2 := end2.Sub(start2)
	fmt.Println("Zeit für das normale sortieren:", elapsed1, "Elemente sortiert:", len(postleitzahlen))
	fmt.Println("Zeit für das Radix sortieren:", elapsed2, "Elemente sortiert:", len(postleitzahlen2))

}
