func removeDuplicates(nums []int) int {
    k := 2
    slow := 1
    n := len(nums)
    count := 1
    for i := 1; i < n; i++ {
        if nums[i] == nums[i-1] {
            count++
        } else {
            count = 1
        }
        if count <= k {
            nums[slow] = nums[i]
            slow++
        }
    }
    return slow
}