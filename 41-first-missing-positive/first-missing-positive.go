func firstMissingPositive(nums []int) int {
    n := len(nums)
    set := map[int]struct{}{}
    smallestSoFar := math.MaxInt64
    for i := 0; i < n; i++ {
        if nums[i] >= 0 {
            smallestSoFar = min(smallestSoFar, nums[i])
            set[nums[i]] = struct{}{}
        }
    }
    start := smallestSoFar+1
    if start != 1 {start = 1}
    for true {
        if _, ok := set[start]; !ok {return start}
        start++
    }
    return -1
}