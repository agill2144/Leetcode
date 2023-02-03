func hIndex(citations []int) int {
    left := 0
    right := len(citations)-1
    
    hIdx := 0
    for left <= right {
        mid := left + (right-left)/2
        diff := len(citations) - mid
        if citations[mid] >= diff {
            if citations[mid] == diff {return diff}
            hIdx = diff
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return hIdx
}