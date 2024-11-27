func maximumSwap(num int) int {
    numStr := fmt.Sprintf("%v", num)
    nums := []int{}
    for i := 0; i < len(numStr); i++ {nums = append(nums, int(numStr[i]-'0'))}
    bucket := make([]int, 10)
    for i := 0; i < len(bucket); i++ {bucket[i] = -1}
    for i := 0; i < len(nums); i++ {bucket[nums[i]] = i}
    swapped := false
    for i := 0; i < len(nums) && !swapped; i++ {
        for j := len(bucket)-1; j > nums[i]; j-- {
            if bucket[j] != -1 && bucket[j] > i {
                nums[i], nums[bucket[j]] = nums[bucket[j]], nums[i]
                swapped = true
                break
            }
        }
    }
    out := 0
    for i := 0; i < len(nums); i++ {
        out = out * 10 + nums[i]
    }
    return out
}

/*
    2199

    9129
    9192

*/

