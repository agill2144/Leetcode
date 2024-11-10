func removeDuplicates(nums []int) int {
    count := 0
    slow := 0
    fast := 0
    k := 2
    for fast < len(nums) {
        curr := nums[fast]
        prev := curr
        if fast - 1 >= 0 {prev = nums[fast-1]}
        if curr == prev {
            count++
        } else {
            count = 1
        }
        if count <= k {
            nums[slow] = nums[fast]
            slow++
            fast++
        } else {
            fast++
        }
    }
    return slow
}