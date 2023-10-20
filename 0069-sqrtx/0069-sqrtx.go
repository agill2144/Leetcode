// how to find sqrt of a number x ?
// in a linear manner; check if 1*1 == x?
// then check 2*2
// then check 3*3
// time = o(ologn)
// space = o(1)
// binary searching on ANSWERS!!!
// We find the list of possible answers
// and that list is usually a range with a max(right) value
// binary search for the left most value whose square comes as close as possible to x ( from the left side )
// what is left most ? if below is a number line from 1 to x
// 1----------------------------------x------
//                                   ^ this is the left most ( less than x but greatest possible )
// we may also find exactly x , we can return that
func mySqrt(x int) int {
    if x < 2 {return x}
    
    // use binary search to find the right window
    // or we could just use right := x/2
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