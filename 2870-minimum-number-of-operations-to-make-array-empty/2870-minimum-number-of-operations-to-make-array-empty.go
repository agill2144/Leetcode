func minOperations(nums []int) int {
    freq := map[int]int{}
    for i := 0; i < len(nums); i++ {freq[nums[i]]++}
    total := 0
    for _, v := range freq {
        // we need at min 2 or 3 dupes, if its not, return -1
        if v == 1 {return -1}

        // always divide by 3 
        // even if its not completely divisible, 14/3 = 4
        total += (v/3)
        // then update the value with whatever is left over
        left := v % 3        
        // if there is anything left, perform 1 more operation ( divide by 2)
        if left != 0 {total++}
    }
    return total
}