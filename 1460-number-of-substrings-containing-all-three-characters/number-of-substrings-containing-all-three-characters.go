func numberOfSubstrings(s string) int {
    idx := map[byte]int{'a':-1,'b':-1,'c':-1}
    count := 0
    n := len(s)
    for i := 0; i < n; i++ {
        idx[s[i]] = i
        if idx['a'] != -1 && idx['b'] != -1 && idx['c'] != -1 {
            start := min(idx['a'], min(idx['b'], idx['c']))
            count += (start+1)
        }
    }
    return count
}