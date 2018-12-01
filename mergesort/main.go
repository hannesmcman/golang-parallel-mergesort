package main

import (
	"./lib"
	"math/rand"
	"time"
	"runtime"
	"fmt"
)


func main() {

	var nCPU int
	fmt.Print("How many CPUs? ")
  fmt.Scanf("%d", &nCPU)
	fmt.Printf("Number of cpus: %d \n", nCPU)
	runtime.GOMAXPROCS(nCPU)

	numberOfItems := 1000000;
	threshold := 10000;

	items := rand.Perm(numberOfItems)

	start := time.Now()
	mergesort.MergeSort(items, threshold)

	fmt.Printf("Took %s to sort %d items (with threshold %d).\n", time.Since(start), numberOfItems, threshold)

	start = time.Now()
	mergesort.MergeSort(items, -1)

	fmt.Printf("Took %s to sort %d items (synchronous).\n", time.Since(start), numberOfItems)
}
