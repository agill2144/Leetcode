func hIndex(citations []int) int {
    left := 0
    right := len(citations)-1
    n := len(citations)
    hIdx := 0
    for left <= right {
        mid := left + (right-left)/2
        tmp := n-mid
        if citations[mid] >= tmp {
            hIdx = tmp
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return hIdx
}