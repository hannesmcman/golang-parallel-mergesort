# Parallelizing Merge Sort

## Getting Started

 1. Clone the repo
 2. If on thing2, add the following to the end of `~/.profile`:
        export PATH=$PATH:/usr/local/go/bin
 3. `cd golang-parallel-mergesort`
 4. `go run ./mergesort/main.go` or `go run ./mergesort_channels/main.go`

## Two Versions

 1. In `./mergesort` directory, we have a version that uses `sync.WaitGroup`.
 2. In `./mergesort_channels` directory, there is a version that uses golang channels.
