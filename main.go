package main

import (
	"./mergesort"
	"math/rand"
	"time"
	"log"
	"runtime"
	"fmt"
)


func main() {

	var nCPU int
  fmt.Scanf("%d", &nCPU)
	log.Printf("Number of cpus: %d", nCPU)
	runtime.GOMAXPROCS(nCPU)

	numberOfItems := 1000000;
	threshold := 10000;

	items := rand.Perm(numberOfItems)

	start := time.Now()
	mergesort.MergeSort(items, threshold)

	log.Printf("Took %s to sort %d items (with threshold %d).", time.Since(start), numberOfItems, threshold)

	start = time.Now()
	mergesort.MergeSort(items, -1)

	log.Printf("Took %s to sort %d items (synchronous).", time.Since(start), numberOfItems)
}
