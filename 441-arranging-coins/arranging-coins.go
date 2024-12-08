func arrangeCoins(n int) int {
    left := 1
    right := n
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        coinsNeeded := (mid*(mid+1))/2
        if coinsNeeded > n {
            right = mid-1
        } else {
            ans = mid
            left = mid+1
        }
    }
    return ans
}