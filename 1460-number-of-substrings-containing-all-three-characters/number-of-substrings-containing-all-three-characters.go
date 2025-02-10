func numberOfSubstrings(s string) int {
    a, b, c := 0, 0, 0
    total := 0
    left := 0
    n := len(s)
    for i := 0; i < n; i++ {
        if s[i] == 'a' {a++}
        if s[i] == 'b' {b++}
        if s[i] == 'c' {c++}
        for a >= 1 && b >= 1 && c >= 1 {
            total += (n-i)
            if s[left] == 'a' {a--}
            if s[left] == 'b' {b--}
            if s[left] == 'c' {c--}
            left++
        }
    }
    return total
}