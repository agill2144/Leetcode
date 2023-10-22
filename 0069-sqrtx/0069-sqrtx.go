/*
    approach: binary search on answers!
    - we know the sqrt is between 1 and x
    - THIS IS A SORTED RANGE
    - which means we can easily binary search on these answers
    
    *BINARY SEARCH ON ANSWERS IS A COMMON PATTER*
    - Identify if there is a range we can use
    - ITS a range, ranges are sorted
    - then binary search on that range

    time = o(logx)
    space = o(1)
*/
func mySqrt(x int) int {
    ans := 0
    left := 1
    right := x
    
    for left <= right {
        mid := left + (right-left)/2
        if mid * mid <= x {
            if mid*mid == x {return mid}
            ans = mid
            // we may have undershot, look for a higher number ( we want the left most )
            left = mid+1
        } else {
            // our sq is too high, pick a smaller number
            right = mid-1
        }
    }
    return ans
}


/*
    approach: linear scan
    - square root of a number ( square! , i.e a^2 )
    - so we need to find such $a whose square == x
    - we can start from 1 and go all the upto x itself
    - and x is not a perfect square, we want the left most possible ans
    
    eg1; x = 25; return 5; because 5^2 = 25
    eg2; x = 28; return 5; because 5^2 = 25, 6^2 = 26, which is greater than x, and we want the leftmost possible ans
    
    time = o(x)
    space = o(1)
*/
// func mySqrt(x int) int {
//     ans := 0
//     for i := 1; i <= x; i++ {
//         if i*i <= x {
//             ans = i
//         } else {
//             break
//         }
//     }
//     return ans
// }