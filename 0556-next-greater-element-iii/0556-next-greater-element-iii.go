func nextGreaterElement(n int) int {
    nStr := strconv.Itoa(n)
    idx := -1
    nums := []int{}
    for i := 0; i < len(nStr); i++ {
        nums = append(nums, int(nStr[i]-48))
    }
    
    // find the first decreasing element from right
    for i := len(nums)-2; i>=0; i-- {
        curr := nums[i]
        next := nums[i+1]
        if curr < next {
            idx = i
            break
        }
    }
    
    // didnt find it, this is the largest number we can have ( i.e the everything is in desc order from left to right )
    // in next-permutation, we still processed because we have to go back to the smallest number from the largest
    // but in this problem, we have to exit and return -1 if its not possible
    if idx == -1 {return -1}
    
    
    // find the next right element of idx from right side
    right := -1
    for i := len(nums)-1; i>= 0; i-- {
        curr := nums[i]
        if curr > nums[idx] {
            right = i
            break
        }
    }
    
    // swap idx, and right
    nums[idx], nums[right] = nums[right], nums[idx]
    // reverse from idx+1:end
    nums = reverse(idx+1, nums)
    
    
    out := 0
    for i := 0; i < len(nums); i++ {
        out = out * 10 + nums[i]
    }
    // Note that the returned integer should fit in 32-bit integer, if there is a valid answer but it does not fit in 32-bit integer, return -1.
    if out > math.MaxInt32 {return -1}
    
    return out
}


func reverse(left int, nums []int) []int {
    right := len(nums)-1
    for left < right {
        nums[left],nums[right] = nums[right],nums[left]
        left++
        right--
    }
    return nums
}