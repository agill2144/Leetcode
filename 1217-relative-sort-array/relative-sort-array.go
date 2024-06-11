func relativeSortArray(arr1 []int, arr2 []int) []int {
    bucket := make([]int, 1001)
    for i := 0; i < len(arr1); i++ {bucket[arr1[i]]++}
    out := []int{}
    for i := 0; i < len(arr2); i++ {
        val := arr2[i]
        for bucket[val] != 0 {
            out = append(out, val)
            bucket[val]--
        }
    }

    for i := 0; i < len(bucket); i++ {
        val := i
        for bucket[val] != 0 {
            out = append(out, val)
            bucket[val]--
        }
    }
    return out
}
