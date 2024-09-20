func isPerfectSquare(num int) bool {
    left := 1
    right := num
    for left <= right {
        mid := left + (right-left)/2
        sq := mid*mid
        if sq == num {return true}
        if sq > num {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return false
}