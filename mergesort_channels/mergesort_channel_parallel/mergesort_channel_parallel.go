package mergesort_channel_parallel

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

