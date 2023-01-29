func permute(nums []int) [][]int {
    out := [][]int{}
    var dfs func(start int)
    dfs = func(start int) {
        // base
        if start == len(nums) {
            newL := make([]int, len(nums))
            copy(newL, nums)
            out = append(out, newL)
            return
        }
        
        // logic
        for i := start; i < len(nums); i++ {
            nums[i],nums[start] = nums[start],nums[i]
            dfs(start+1)
            nums[i],nums[start] = nums[start],nums[i]
        }
    }
    dfs(0)
    return out
}
