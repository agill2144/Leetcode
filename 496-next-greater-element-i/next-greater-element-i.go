func nextGreaterElement(nums1 []int, nums2 []int) []int {
    // find ngr ( next greater to right ) for all elements in nums2
    // keeping saving nums2 element to idx 
    n2Idx := map[int]int{}
    ngr := make([]int, len(nums2))
    st := []int{}
    for i := len(nums2)-1; i >= 0; i-- {
        for len(st) != 0 && st[len(st)-1] <= nums2[i] {
            st = st[:len(st)-1]
        }
        if len(st) != 0 {
            ngr[i] = st[len(st)-1]
        } else {
            ngr[i] = -1
        }
        st = append(st, nums2[i])
        n2Idx[nums2[i]] = i
    }
    out := []int{}
    for i := 0; i < len(nums1); i++ {
        num := nums1[i]
        idxInNums2 := n2Idx[num]
        ngrVal := ngr[idxInNums2]
        out = append(out, ngrVal)
    }
    return out
}