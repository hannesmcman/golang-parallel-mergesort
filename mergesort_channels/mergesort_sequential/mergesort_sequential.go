package mergesort_sequential

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

func MergeSort(data []int) (r []int){
    if len(data) == 1 {
        r = data
        return
    }

    middle := len(data)/2

    ldata :=  MergeSort(data[:middle])
    rdata :=  MergeSort(data[middle:])

    r = Merge(ldata, rdata)
    return 
}
