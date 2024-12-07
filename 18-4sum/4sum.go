func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
    out := [][]int{}
    for i := 0; i < len(nums); i++ {
        if i > 0 && nums[i] == nums[i-1] {continue}
        iTarget := target-nums[i]
        for j := i+1; j < len(nums); j++ {
            if j > i+1 && nums[j] == nums[j-1] {continue}
            jTarget := iTarget-nums[j]
            left := j+1
            right := len(nums)-1
            for left < right {
                if nums[left] + nums[right] == jTarget {
                    out = append(out, []int{nums[i], nums[j], nums[left], nums[right]})
                    left++
                    for left < right && nums[left] == nums[right] {left++}
                    right--
                    for left < right && nums[right] == nums[right+1] {right--}
                } else if nums[left] + nums[right] > jTarget {
                    right--
                } else {
                    left++
                }
            }
        }
    }
    return out
}