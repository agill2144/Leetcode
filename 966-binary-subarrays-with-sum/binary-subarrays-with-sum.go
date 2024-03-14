func numSubarraysWithSum(nums []int, goal int) int {
    count := 0
    for i := 0; i < len(nums); i++ {
        sum := 0
        for j := i; j < len(nums); j++ {
            sum += nums[j]
            if sum == goal {count++}
        }
    }
    return count
}