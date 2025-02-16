func nextGreaterElement(nums1 []int, nums2 []int) []int {
    st := []int{} // idx
    n1 := len(nums1)
    n2 := len(nums2)
    idx := map[int]int{}
    ng := make([]int, n2)
    for i := 0; i < n2; i++ {
        ng[i] = -1
        idx[nums2[i]] = i
        for len(st) != 0 && nums2[i] > nums2[st[len(st)-1]] {
            top := st[len(st)-1]
            st = st[:len(st)-1]
            ng[top] = nums2[i]
        }
        st = append(st, i)
    }
    out := make([]int, n1)
    for i := 0; i < n1; i++ {
        val, ok := idx[nums1[i]]
        if !ok {
            out[i] = -1
        } else {
            out[i] = ng[val]
        }
    }
    return out
}