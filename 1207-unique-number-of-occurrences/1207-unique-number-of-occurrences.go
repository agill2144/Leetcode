func uniqueOccurrences(arr []int) bool {
    freqMap := map[int]int{}
    for i := 0; i < len(arr); i++ {
        freqMap[arr[i]]++
    }
    occ := map[int]struct{}{}
    for _, v := range freqMap {
        if _, ok := occ[v]; ok {return false}
        occ[v] = struct{}{}
    }
    return true
}