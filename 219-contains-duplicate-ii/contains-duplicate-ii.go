func containsNearbyDuplicate(nums []int, k int) bool {
    idx := map[int]int{}
    for i := 0; i < len(nums); i++ {
        seenAt, ok := idx[nums[i]]
        if ok && i-seenAt <= k {
            return true
        }
        idx[nums[i]] = i
    }
    return false
}
