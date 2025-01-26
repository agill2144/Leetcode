func smallerNumbersThanCurrent(nums []int) []int {
    n := len(nums)
    freq := map[int]int{}
    for i := 0; i < n; i++ {freq[nums[i]]++}

    // bucket acts as prefix count
    // count num of elements < ith bucket element will always be
    // prev count ( bucket[i-1] ) + num times prev element appears ( freq[i-1] )
    bucket := make([]int, 501)
    for i := 0; i < len(bucket); i++ {
        if i == 0 {continue}
        bucket[i] = bucket[i-1] + freq[i-1]        
    }
    out := make([]int, n)
    for i := 0; i < n; i++ {
        out[i] = bucket[nums[i]]
    }
    return out
}