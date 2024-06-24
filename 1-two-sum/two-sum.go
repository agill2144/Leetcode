
func twoSum(nums []int, target int) []int {
    idxMap := map[int]int{}
    for i := 0; i < len(nums); i++ {
        diff := target-nums[i]
        idx, ok := idxMap[diff]
        if ok {
            return []int{idx, i}
        }
        idxMap[nums[i]] = i
    }
    return nil
}