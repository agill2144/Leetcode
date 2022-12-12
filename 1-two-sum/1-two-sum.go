func twoSum(nums []int, target int) []int {
    m := map[int]int{}
    for i := 0; i < len(nums); i++ {
        diff := target-nums[i]
        idx, exists := m[diff] 
        if exists{
            return []int{i, idx}
        }
        m[nums[i]] = i

    }
    return nil
}