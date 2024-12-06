// pick a number from 1 to n
// this is a range to select an answer for
// ranges are sorted
// binary search on answers
func maxCount(banned []int, n int, maxSum int) int {
    set := map[int]bool{}
    for i := 0; i < len(banned); i++ {set[banned[i]] = true}
    left := 1
    right := n
    ans := 0
    // m = len(banned)
    // tc = o(logn * m)
    // sc = o(m) for banned hashset
    for left <= right {
        mid := left + (right-left)/2
        sum := mid*(mid+1)/2
        count := mid
        for k, _ := range set {
            if k <= mid {sum -= k; count--}
        }
        if sum <= maxSum {
            ans = count
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return ans
}