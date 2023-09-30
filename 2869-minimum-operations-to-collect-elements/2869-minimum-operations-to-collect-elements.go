func minOperations(nums []int, k int) int {
    collections := map[int]struct{}{}
    for i := len(nums)-1; i>=0; i-- {
        if nums[i] >= 1 && nums[i] <= k {
            if _, ok := collections[nums[i]]; !ok {
                collections[nums[i]] = struct{}{}
            }
        }

        if len(collections) == k {return len(nums)-i}
    } 
    return -1
}