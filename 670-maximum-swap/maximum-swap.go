func maximumSwap(num int) int {
    if num <= 9 {return num}
    numStr := fmt.Sprintf("%v", num)
    nums := []int{}
    for i := 0; i < len(numStr); i++ {
        nums = append(nums, int(numStr[i]-'0'))
    }
    idxs := make([]int, 10)
    for i := 0; i < len(nums); i++ {
        idxs[nums[i]] = -1
        idxs[nums[i]] = i
    }
    swapped := false
    for i := 0; i < len(nums) && !swapped; i++ {
        for j := 9; j > nums[i]; j-- {
            if idxs[j] != -1 && idxs[j] > i {
                nums[i], nums[idxs[j]] = nums[idxs[j]], nums[i]
                swapped = true
                break
            }
        }
    }
    res := 0
    for i := 0; i < len(nums); i++ {
        res = res * 10 + nums[i]
    }
    return res
}