func hIndex(citations []int) int {
    // we are looking for an index where len(arr)-idx >= citations[idx]
    // we are looking for an h index where h of the papers have >= h citations
    // we are looking for the first trip because because beyond that trip , all papers are going to have > h citations
    // because this is sorted array
    n := len(citations)
    left := 0
    right := n-1
    hIdx := 0
    for left <= right {
        mid := left + (right-left)/2
        diff := n-mid
        if citations[mid] >= diff {
            hIdx = diff
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return hIdx
}