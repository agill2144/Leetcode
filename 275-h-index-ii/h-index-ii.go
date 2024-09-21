/*
    - find the first position where h <= citations[i]
    - and h = n-idx
    - so really its, find the first position where; n-idx <= citations[idx]
*/
func hIndex(citations []int) int {
    n := len(citations)
    left := 0
    right := n-1
    ans := 0
    for left <= right {
        mid := left + (right-left)/2
        h := n-mid
        
        if h > citations[mid] {
            left = mid+1
        } else {
            ans = h
            right = mid-1
        }
    }
    return ans
}