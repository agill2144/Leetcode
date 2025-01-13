func minimumLength(s string) int {
    freq := make([]int, 26)
    for i := 0; i < len(s); i++ {freq[int(s[i]-'a')]++}
    n := len(s)
    for i := 0; i < len(freq); i++ {
        if freq[i] <= 2 {continue}
        if freq[i] % 2 == 1 {
            n -= freq[i]-1
        } else {
            n -= (freq[i]-2)
        }
    }
    return n
}
