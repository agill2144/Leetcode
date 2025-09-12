func isPerfectSquare(num int) bool {
    if num == 1 {return true}
    left := 2
    right := num
    for left <= right {
        mid := left + (right-left)/2
        sq := mid*mid
        if sq == num {return true}
        if num > sq {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return false
}