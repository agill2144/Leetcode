// greedy
func minOperations(nums []int) int {
    freq := map[int]int{}
    for i := 0; i < len(nums); i++ {freq[nums[i]]++}
    total := 0
    for _, v := range freq {
        // we need at min 2 or 3 dupes, if its not, return -1
        if v == 1 {return -1}

        // be greedy
        // and always divide by 3 
        // even if its not completely divisible by 3
        total += (v/3)
        // then update the value with whatever is left over
        v = v % 3        
        // if there is anything left, perform 1 more operation; i.e increment operation count
        // ( divide by 2 ; even if its not perfectly divisible by 2 ; eg v=1)
        // any number divided by 3 ( which is not perfectly divisible ), will ALWAYS leave a remainder of either 1 or 2
        // therefore we can imply and blindly increment operation count by 1
        if v != 0 {total++}
    }
    return total
}