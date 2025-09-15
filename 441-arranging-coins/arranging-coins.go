func arrangeCoins(n int) int {
    if n == 1 {return 1}
    left := 1
    right := n
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        ap := mid*(mid+1)/2
        if n >= ap {
            ans = mid
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return ans
}