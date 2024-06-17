func find132pattern(nums []int) bool {
    st := []int{} // indices, viable numsK candidates
    n := len(nums)
    numsK := math.MinInt64
    for i := n-1; i >= 0; i-- {
        // once we have numsK, numsJ becomes irrelevant
        // because numsJ was the largest value once discovered and processed/setted numsK val
        // therefore nums[k] < nums[j] is already satisfied is numsK is set
        // if this number is < numsK, it means we have found nums[i] < nums[k]
        // return true
        if nums[i] < numsK {return true}

        // the iterator ptr is acting as our numsJ val
        /// numsJ is supposed to the largest and its supposed to set numsK value
        // as soon as we run into a ele > top of stack(potential numsK vals), we have found nums[k] < nums[j] (and j < k)
        // process the stack until nums[i] > top of stack, and each popped element is a better numsK val
        // therefore each popped element becomes numsK ( overriden each time )
        for len(st) != 0 && nums[i] > nums[st[len(st)-1]] {
            numsK = nums[st[len(st)-1]]
            st = st[:len(st)-1]
        }

        // push this value as potential numsK val
        st = append(st, i)
    }
    return false
}