func nextGreaterElement(nums1 []int, nums2 []int) []int {
    idx := map[int]int{}
    st := []int{} // idx
    ng := make([]int, len(nums2))
    for i := 0; i < len(nums2); i++ {
        ng[i] = -1
        idx[nums2[i]] = i
        for len(st) > 0 && nums2[i] > nums2[st[len(st)-1]] {
            top := st[len(st)-1]
            st = st[:len(st)-1]
            ng[top] = nums2[i]        
        }
        st = append(st, i)
    }
    res := make([]int, len(nums1))
    for i := 0; i < len(nums1); i++ {
        pos, ok := idx[nums1[i]]
        if !ok {continue}
        res[i] = ng[pos]
    }
    return res
}