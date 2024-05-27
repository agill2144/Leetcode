func numberOfPairs(nums1 []int, nums2 []int, k int) int {
    maxN1 := math.MinInt64
    n1Freq := map[int]int{}
    // n1
    for i := 0; i < len(nums1); i++ {
        n1Freq[nums1[i]]++
        maxN1 = max(maxN1, nums1[i])
    }
    n2Freq := map[int]int{}
    // n2
    for i := 0; i < len(nums2); i++ {
        n2Freq[nums2[i]*k]++
    }
    count := 0
    // n2
    for ele, freq := range n2Freq {
        b := ele
        // o(maxNumInNums1)
        for b <= maxN1 {
            val, ok := n1Freq[b]
            if ok {
                count += (val * freq)
            }
            b += ele
        }
    }
    return count

}