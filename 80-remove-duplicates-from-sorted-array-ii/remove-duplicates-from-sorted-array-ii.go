func removeDuplicates(nums []int) int {
    n := len(nums)
    k := 2
    slow := 1
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