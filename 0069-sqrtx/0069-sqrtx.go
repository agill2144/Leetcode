func mySqrt(x int) int {
    if x < 2 {return x}
    
    right := 1
    for right * right < x {
        right *= 2
    }
    
    ans := -1
    left := 1
    
    for left <= right {
        mid := left + (right-left)/2
        if mid * mid <= x {
            if mid * mid == x {return mid}
            ans = mid
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return ans
}