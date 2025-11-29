/*
    approach: freq map
    - put smaller nums into a freq map
    - then iterate over other nums , check if exists in freq and if it does save it , decrement its counter
    m = len(nums1)
    n = len(nums2)
    tc = o(m)+o(n)
    sc = o(n)
*/
func intersection(nums1 []int, nums2 []int) []int {
    freq := map[int]int{}    
    for i := 0; i < len(nums1); i++ {
        freq[nums1[i]]++
    }
    out := []int{}
    for i := 0; i < len(nums2); i++ {
        if freq[nums2[i]] > 0 {
            freq[nums2[i]] = 0
            out = append(out, nums2[i])
        }
    }
    return out
}