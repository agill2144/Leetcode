func minimumAddedInteger(nums1 []int, nums2 []int) int {
    n1 := len(nums1)
    n2 := len(nums2)
    if n2 > n1 {return math.MaxInt64}
    
    freq := map[int]int{}
    for i := 0; i < len(nums2); i++ { freq[nums2[i]]++  }
        
    f := len(freq)
    for i := -1000; i <= 1000; i++ {
        missing := 0
        count := 0
        local := map[int]int{}
        for j := 0; j < n1; j++ {
            val := nums1[j]+i
            desiredCount, ok := freq[val]
            currentCount := local[val]
            if !ok || currentCount == desiredCount {
                missing++
                if missing > 2 {break}
                continue
            }

            local[val]++
            if local[val] == desiredCount {
                count++
                if count == f {
                    return i
                }
            }
            
        }
    }
    return math.MaxInt64
}


func abs(x int) int {
    return x
}
