func twoSum(nums []int, target int) []int {
    idx := map[int]int{}
    for i := 0; i < len(nums); i++ {
        diff := target - nums[i] 
        if _, ok := idx[diff]; ok {
            return []int{i, idx[diff]}
        }
        idx[nums[i]] = i
    }
    return nil
}