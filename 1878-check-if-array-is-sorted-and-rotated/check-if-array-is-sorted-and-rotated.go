/*
    approach: 2 pass
    - find the ending and starting nodes
    - ending = largest element
    - starting = smallest element
    - if its rotated, ending will show up right before starting
    - meaning if nums[i] > nums[i+1]
    - ith element is our ending node
    - i+1th element is our starting node
    - then take another from starting node until ending node
    - to make sure all elements are sorted

    tc = o(2n) = o(n)
    sc = o(1)
*/
func check(nums []int) bool {
    n := len(nums)
    startIdx := -1
    endIdx := -1
    for i := 0; i+1 < n; i++ {
        if nums[i] > nums[i+1] {
            startIdx = i+1
            endIdx = i
            break
        }
    }
    // already sorted, not rotated
    // examples;
    // [1,1,1,1,1,1]
    // [1,2,2,2,2,5,5,6]
    // [1]
    if startIdx == -1 {return true}

    for startIdx%n != endIdx%n {
        if nums[startIdx%n] > nums[(startIdx+1)%n] {return false}
        startIdx++
    }
    return true
}