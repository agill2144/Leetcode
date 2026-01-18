func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
    out := [][]int{}
    for i := 0; i < len(nums); i++ {
        if i > 0 && nums[i] == nums[i-1] {continue}
        for j := i+1; j < len(nums); j++ {
            if j != i+1 && nums[j] == nums[j-1] {continue}
            left := j+1
            right := len(nums)-1
            for left < right {
                total := nums[left] + nums[right] + nums[i] + nums[j]
                if total == target {
                    out = append(out, []int{nums[i], nums[j], nums[left], nums[right]})
                    left++
                    for left < right && nums[left] == nums[left-1] {left++}
                    right--
                    for left < right && nums[right] == nums[right+1] {right--}
                } else if total > target {
                    right--
                } else {
                    left++
                }
            }
        }
    }
    return out
}