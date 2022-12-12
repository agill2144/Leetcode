func sortedSquares(nums []int) []int {
    
    /*
    approach 1
    square them all into a new array
    and sort the new arry
    time: o(nlogn)
    space: o(1) if we do not consider output array as part of space
    
    approach 2 : using two pointers
    Have a left and right pointer ( 0, len(nums)-1)
    Cretae an output array of same size
    Have a pointer at len(output)-1
    
    We will sqaure each left and right pointer from nums
    and whoever is bigger will be placed in the back of the output array.
    Whoever was bigger, we will move that pointer
    and we will also move the output ptr back 1
    
    
    we will do this until leftPtr <= rightPtr
    
    Potential follow up:
    - Write the results back into the same array? - then we will have to sort at the end or else we loose previous element value using two ptrs
    */
    
    if nums == nil || len(nums) == 0 {
        return nil
    }
    
    out := make([]int, len(nums))
    left := 0
    right := len(nums)-1
    outPtr := len(out)-1
    
    for left <= right {
        leftSq := nums[left] * nums[left]
        rightSq := nums[right] * nums[right]
        
        if leftSq > rightSq {
            out[outPtr] = leftSq
            left++
        } else {
            out[outPtr] = rightSq
            right--
        }
        outPtr--
    }
    return out
}