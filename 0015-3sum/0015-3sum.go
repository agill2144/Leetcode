func threeSum(nums []int) [][]int {
    out := [][]int{}
    sort.Ints(nums)
    for i := 0; i < len(nums); i++ {
        if i != 0 && nums[i] == nums[i-1] {continue}
        target := 0-nums[i]
        l := i+1
        r := len(nums)-1
        for l < r {
            sum := nums[l] + nums[r]
            if sum == target {
                out = append(out, []int{nums[i], nums[l], nums[r]})
                l++
                for l < r && nums[l] == nums[l-1] {l++}
                r--
                for r > l && nums[r] == nums[r+1] {r--}
                
            } else if sum > target {
                r--
            } else {
                l++
            }
        }
    }
    return out
    
}