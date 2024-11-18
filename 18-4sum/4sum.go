func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
    n := len(nums)
    out := [][]int{}
    set := map[[4]int]bool{}
    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            left := j+1
            right := n-1
            for left < right {
                sum := nums[i] + nums[j] + nums[left] + nums[right]
                if sum == target {
                    tmp := [4]int{nums[i], nums[j], nums[left], nums[right]}
                    if !set[tmp] {
                        out = append(out, []int{nums[i], nums[j], nums[left], nums[right]})
                        set[tmp] = true
                    }
                    left++
                    for left < right && nums[left] == nums[left-1] {left++}
                    right--
                    for left < right && nums[right] == nums[right+1] {right--}
                } else if sum > target {
                    right--
                } else {
                    left++
                }
            }
        }
    }
    return out
}