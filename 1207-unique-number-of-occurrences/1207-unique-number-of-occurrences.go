func uniqueOccurrences(arr []int) bool {
    freq := map[int]int{}
    for i := 0; i < len(arr); i++ {freq[arr[i]]++}
    set := map[int]struct{}{}
    for _, v := range freq {
        if _, ok := set[v]; ok {return false}
        set[v] = struct{}{}
    }
    return true
}