package main

import (
	"./mergesort_channel_parallel"
	"./mergesort_sequential"
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

	items := rand.Perm(numberOfItems)

	start := time.Now()
	mergesort_sequential.MergeSort(items)

	log.Printf("Took %s to sort %d items.", time.Since(start), numberOfItems)

	start = time.Now()
	result := make(chan []int)
	go mergesort_channel_parallel.MergeSort(items, result)

	log.Printf("Took %s to sort %d items (parallel).", time.Since(start), numberOfItems)
}

