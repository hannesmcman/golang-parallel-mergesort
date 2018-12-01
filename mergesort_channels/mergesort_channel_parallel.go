package main

import (
    // "fmt"
    "math/rand"
    )

func Merge(ldata []int, rdata []int) (result []int) {
    result = make([]int, len(ldata) + len(rdata))
    lidx, ridx := 0, 0

    for i:=0;i<cap(result);i++ {
        switch {
            case lidx >= len(ldata):
                result[i] = rdata[ridx]
                ridx++
            case ridx >= len(rdata):
                result[i] = ldata[lidx]
                lidx++
            case ldata[lidx] < rdata[ridx]:
                result[i] = ldata[lidx]
                lidx++
            default:
                result[i] = rdata[ridx]
                ridx++
        }
    }

    return
}

func MergeSort(data []int, r chan []int) {
    if len(data) == 1 {
        r <- data
        return
    }

    leftChan := make(chan []int)
    rightChan := make(chan []int)
    middle := len(data)/2

    go MergeSort(data[:middle], leftChan)
    go MergeSort(data[middle:], rightChan)

    ldata := <-leftChan
    rdata := <-rightChan

    close(leftChan)
    close(rightChan)
    r <- Merge(ldata, rdata)
    return
}

func main() {
    // s := []int{22, 8, 3, 31, 4, 2, 42, 1, 16, 6, 11, 25, 9, 8, 10, 12, 18, 14, 7, 15}
    numberOfItems := 1000000;

    s := rand.Perm(numberOfItems)
    
    result := make(chan []int)
    go MergeSort(s, result)

    // r := <- result
    // for _,v := range r {
    //     fmt.Println(v)
    // }
    close(result)
}

