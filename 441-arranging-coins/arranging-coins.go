func arrangeCoins(n int) int {
    if n <= 1 {
        if n == 1 {return 1}
        return 0
    }
    left := 1
    right := n
    ans := 0
    for left <= right {
        mid := left + (right-left)/2
        ap := mid*(mid+1)/2
        if ap <= n {
            ans = mid
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return ans
}