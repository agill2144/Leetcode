func mySqrt(x int) int {
    if x == 0 {return 0}
    left := 1
    right := x
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        sq := mid*mid
        if sq <= x {
            if sq == x {return mid}
            ans = mid
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return ans
}