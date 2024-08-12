func numberOfSubstrings(s string) int {
    a := -1
    b := -1
    c := -1
    count := 0
    n := len(s)
    for i := 0; i < n; i++ {
        if s[i] == 'a' {a=i}
        if s[i] == 'b' {b=i}
        if s[i] == 'c' {c=i}
        if a != -1 && b != -1 && c != -1 {
            start := min(a, min(b, c))
            count += (start+1)
        }
    }
    return count
}