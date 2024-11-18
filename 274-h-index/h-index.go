func hIndex(citations []int) int {
    sort.Ints(citations)
    n := len(citations)
    left := 0
    right := n-1
    ans := 0
    for left <= right {
        mid := left + (right-left)/2
        h := n-mid
        if citations[mid] >= h {
            ans = h
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return ans

}