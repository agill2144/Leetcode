/*
    approach: binary search on answers
    - we have to find 2 such sq whose sum == c
    - sounds like 2Sum
    -
*/
func judgeSquareSum(c int) bool {
    for i := 0; i*i <= c; i++ {
        
        // complement searh
        b := c - (i*i)

        // find $a whose sq == b
        left := 0
        right := b
        for left <= right {
            mid := left + (right-left)/2
            if mid * mid == b {
                return true
            } else if mid * mid > b {
                right = mid-1
            } else {
                left = mid+1
            }
        }
    }
    return false
}