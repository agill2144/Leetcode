func arrangeCoins(n int) int {
    left := 1
    right := n
    for left <= right {
        mid := left + (right-left)/2
        totalCoinsUsed := mid*(mid+1)/2
        if totalCoinsUsed > n {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    if right*(right-1)/2 <= n {return right}
    return left
}