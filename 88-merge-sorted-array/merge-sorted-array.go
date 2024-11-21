func merge(nums1 []int, m int, nums2 []int, n int)  {
    ptr := len(nums1)-1
    n1 := m-1
    n2 := n-1
    for n1 >= 0 && n2 >= 0 {
        if nums1[n1] > nums2[n2] {
            nums1[ptr] = nums1[n1]
            n1--
        } else {
            nums1[ptr] = nums2[n2]
            n2--
        }
        ptr--
    }
    for n1 >= 0 {nums1[ptr] = nums1[n1]; n1--; ptr--}
    for n2 >= 0 {nums1[ptr] = nums2[n2]; n2--; ptr--}
    
}