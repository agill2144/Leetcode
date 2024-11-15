/*
    approach: 2 pass
    - because the array was sorted at one point, then rotated
    - we need to find the breaking point
    - that is, at what idx, does the asc order break from left to right
    - we need to find the first such idx
    - once we found the breaking idx, then a simple loop around the whole array
    - starting from breaking idx, should be in asc order
    - example; [3,4,5,1,2]
    -                 ^ breaking idx
    - now everything after that breaking idx ( wrapped around the whole array )
    - should be in asc / sorted order
    time = o(n)
    space = o(1)
*/
func check(nums []int) bool {
    // find the breaking idx
    idx := -1
    for i := 0; i < len(nums); i++ {
        if i == 0 {continue}
        if nums[i] < nums[i-1] {idx = i; break}
    }
    if idx == -1 {return true}
    n := len(nums)
    // run a loop wrapping around the whole array
    // and validate sorted order
    start := idx+1
    end := idx
    for start % n != end {
        if nums[(start-1)%n] <= nums[start%n] {start++} else {break}
    }
    return (start%n) == end
}