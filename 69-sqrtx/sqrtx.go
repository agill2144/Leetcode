func mySqrt(x int) int {
    // right most on left side of x
    // x = 8; right most is 2 because 3*3 is 9 and thats > x
    left := 1
    right := (x/2)+1
    ans := 0
    for left <= right {
        mid := left + (right-left)/2
        sq := mid*mid
        if sq <= x {
            ans = mid
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return ans
}