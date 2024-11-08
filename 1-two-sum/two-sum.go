func twoSum(nums []int, target int) []int {
    idx := map[int]int{}
    for i := 0; i < len(nums); i++ {
        diff := target-nums[i]
        val, ok := idx[diff]
        if ok {return []int{val, i}}
        idx[nums[i]] = i
    }
    return nil
}