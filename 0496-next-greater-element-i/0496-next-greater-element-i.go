func nextGreaterElement(nums1 []int, nums2 []int) []int {
    nums2Idx := map[int]int{}
    for i := 0; i < len(nums2); i++ {
        nums2Idx[nums2[i]] = i
    }
    ans2 := make([]int, len(nums2))
    ans := make([]int, len(nums1))
    
     // monotonic increasing ( asc order from top , sorted array in asc order from left to right )
    // [val, idx]
    st := [][]int{}
    for i := 0; i < len(nums2); i++ {
        for len(st) != 0 && nums2[i] > st[len(st)-1][0] {
            top := st[len(st)-1]; st = st[:len(st)-1]
            ans2[top[1]] = nums2[i]
        }
        st = append(st, []int{nums2[i],i})
    }
    
    for i := 0; i < len(nums1); i++ {
        idx := nums2Idx[nums1[i]]
        ans[i] = ans2[idx]
        if ans[i] == 0 {ans[i] = -1}
    }
    
    return ans
    
}