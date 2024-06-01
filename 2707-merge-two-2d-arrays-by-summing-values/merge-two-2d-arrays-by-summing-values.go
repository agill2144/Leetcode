func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
    n1IdToVal := map[int]int{}
    n2IdToVal := map[int]int{}
    for i := 0; i < len(nums1); i++ {n1IdToVal[nums1[i][0]] = nums1[i][1]}
    for i := 0; i < len(nums2); i++ {n2IdToVal[nums2[i][0]] = nums2[i][1]}
    out := [][]int{}
    n1 := 0
    n2 := 0
    for n1 < len(nums1) && n2 < len(nums2) {
        n1Id, n1Val := nums1[n1][0], nums1[n1][1]
        n2Id, n2Val := nums2[n2][0], nums2[n2][1]
        if n1Id == n2Id {
            out = append(out, []int{n1Id, n1Val+n2Val})
            n1++
            n2++
        } else {
            if n1Id < n2Id {
                val := n2IdToVal[n1Id]
                out = append(out, []int{n1Id, val+n1Val})
                n1++
            } else {
                val := n1IdToVal[n2Id]
                out = append(out, []int{n2Id, val+n2Val})
                n2++
            }
        }
    }

    for n1 < len(nums1) {
        out = append(out, nums1[n1])
        n1++
    }
    for n2 < len(nums2) {
        out = append(out, nums2[n2])
        n2++
    }
    return out
}