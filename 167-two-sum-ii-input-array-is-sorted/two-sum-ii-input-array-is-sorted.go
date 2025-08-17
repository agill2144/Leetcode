// compliment search will obv work
// tc = o(n)
// sc = o(n)
func twoSum(nums []int, target int) []int {
    idx := map[int]int{}
    for i := 0; i < len(nums); i++ {
        diff := target - nums[i]
        val, ok := idx[diff]
        if ok {return []int{val+1, i+1}}
        idx[nums[i]] = i
    }
    return nil
}