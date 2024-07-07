// brute force: nested iterations
// find ngr of nums2 first, while creating idx map
// then loop over nums1, look for its idx in map and find the ngr
func nextGreaterElement(nums1 []int, nums2 []int) []int {
    n1 := len(nums1)
    n2 := len(nums2)
    st := []int{} // holds values
    idxMap := map[int]int{}
    ngr := make([]int, n2)
    for i := n2-1; i >= 0; i-- {
        curr := nums2[i]
        ngr[i] = -1
        idxMap[curr] = i
        for len(st) != 0 && st[len(st)-1] < curr {
            st = st[:len(st)-1]
        }
        if len(st) != 0 {
            ngr[i] = st[len(st)-1]
        }
        st = append(st, curr)
    }
    out := make([]int, n1)
    for i := 0; i < n1; i++ {
        out[i] = ngr[idxMap[nums1[i]]]
    }    
    return out
}