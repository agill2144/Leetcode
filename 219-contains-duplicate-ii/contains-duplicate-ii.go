func containsNearbyDuplicate(nums []int, k int) bool {
    idx := map[int]int{}
    for i := 0; i < len(nums); i++ {
        left, ok := idx[nums[i]]
        if ok {
            if i-left <= k {return true}
        }
        idx[nums[i]] = i
    }
    return false
}
