func minimumAddedInteger(nums1 []int, nums2 []int) int {
    slices.Sort(nums1)
    slices.Sort(nums2)
    
    ans := math.MaxInt
    for i := 0; i <= 2; i++ {
        interval := nums2[0] - nums1[i]
        skiped := i
        
        for j := 0; j < len(nums2) && skiped <= 2; j++ {
            n1 := nums2[j] - interval
            
            for skiped <= 2 && n1 != nums1[j + skiped] {
                skiped++
            }
        }
        
        if skiped <= 2 {
            ans = min(ans, interval)
        }
    }
    
    return ans
}