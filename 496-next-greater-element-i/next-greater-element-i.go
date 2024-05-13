/*
    approach: next-greater-element
    - find the NGR of all elements in nums2
    - and keep saving each elements from nums2 in a idx map
    - then using a second pass, loop over nums1
    - find its idx in nums2 idxMap
    - then using that idx, take the value we have computed in ngr

    time = o(n2) + o(n1)
    space = o(n2)
*/
func nextGreaterElement(nums1 []int, nums2 []int) []int {
    n1 := len(nums1)
    n2 := len(nums2)
    n2IdxMap := map[int]int{}
    ngr := make([]int, n2)
    st := []int{}
    for i := 0; i < n2; i++ {
        ngr[i] = -1
        n2IdxMap[nums2[i]] = i

        curr := nums2[i]
        for len(st) != 0 && curr > nums2[st[len(st)-1]] {
            top := st[len(st)-1]
            st = st[:len(st)-1]
            ngr[top] = curr
        }
        st = append(st, i)

    }
    out := make([]int, n1)
    for i := 0; i < n1; i++ {
        idx := n2IdxMap[nums1[i]]
        out[i] = ngr[idx]
    }
    return out

}