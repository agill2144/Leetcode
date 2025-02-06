func findKthLargest(nums []int, k int) int {
    n := len(nums)
    minval := slices.Min(nums)
    maxval := slices.Max(nums)
    freq := map[int]int{}
    for i := 0; i < n; i++ {freq[nums[i]]++}
    idx := 0
    for i := minval; i <= maxval; i++ {
        if freq[i] == 0 {continue}
        idx += freq[i]
        if idx > n-k {return i}
    }
    return -1
}


