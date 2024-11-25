func findKthLargest(nums []int, k int) int {
    freq := map[int]int{}
    start := math.MaxInt64
    end := math.MinInt64
    for i := 0; i < len(nums); i++ {
        freq[nums[i]]++
        start = min(start,nums[i])
        end = max(end,nums[i])
    }
    /*
        ti = 9-4 = 5
        {3:2, 4:1, 5:2, 6:1 }
        i = 3
        idx = 5

    */
    idx := 0
    targetIdx := len(nums)-k
    for i := start; i <= end; i++ {
        idx += freq[i]
        if idx > targetIdx {return i}
    }    
    return -1
}
