func find132pattern(nums []int) bool {
    // nums[i] < nums[k] < nums[j]
    // nums[j] is supposed to be the largest element
    // nums[k] is supposed to be the 2nd largest element , its idx > idx of j
    // nums[i] is supposed to be smallest and smaller than nums[j]. its idx < idx of j
    st := []int{} // indices, acting as nums[j]
    n := len(nums)
    numsK := math.MinInt64
    for i := n-1; i >= 0; i-- {
        // if we run into a number < numsK , we have our 132 pattern
        // numsK gets set later, because numsK idx is supposed to greatest but its value is supposed to be 2nd largest
        // when numsJ is detected ( largest seen so far ), numsK gets set by numJ processing items in stack
        
        if nums[i] < numsK {return true}
        numsJ := nums[i]
        // if we find a number larger than what we have seen so far from right hand side
        // this means, curr number is largest so far
        // and 2nd largest is somewhere in the stack
        // therefore curr num is numsJ(first largest) ; its idx is also less than elements in stack
        // and elements in stack are potential candidates for numsK
        // but there could be many elements in stack, we want the highest one , therefore keep popping until we can
        // elements in stack will be auto stored in decreasing order from top
        // therefore the 2nd largest is somewhere at the bottom of the stack
        // once we pop, assign it to numsK ( 2nd largest )
        // and now all we need is a number < numsK 
        for len(st) != 0 && numsJ > nums[st[len(st)-1]] {
            numsK = nums[st[len(st)-1]]
            st = st[:len(st)-1]
        }
        st = append(st, i)
    }
    return false
}