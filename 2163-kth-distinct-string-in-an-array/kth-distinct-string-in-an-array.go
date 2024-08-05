func kthDistinct(arr []string, k int) string {
    freq := map[string]int{}
    for i := 0; i < len(arr); i++ {
        freq[arr[i]]++
    }
    for i := 0; i < len(arr); i++ {
        char := arr[i]
        count := freq[char]
        if count > 1 {continue}
        k--
        if k == 0 {return char}
    }
    return ""
}