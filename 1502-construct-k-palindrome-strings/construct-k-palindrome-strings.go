func canConstruct(s string, k int) bool {
    if len(s) == k {return true}
    if len(s) < k {return false}
    freq := make([]int, 26)
    for i := 0; i < len(s); i++ {
        freq[int(s[i]-'a')]++
    }
    oddCount := 0
    for i := 0; i < len(freq); i++ {
        if freq[i] % 2 != 0 {oddCount++}
    }
    return oddCount <= k
}