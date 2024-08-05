func numberOfSubstrings(s string) int {
    a := 0
    b := 0
    c := 0
    left := 0
    n := len(s)
    count := 0
    for i := 0; i < n; i++ {
        if s[i] == 'a' {a++}
        if s[i] == 'b' {b++}
        if s[i] == 'c' {c++}
        for a >= 1 && b >= 1 && c >= 1 {
            count += (n-i)
            if s[left] == 'a' {a--}
            if s[left] == 'b' {b--}
            if s[left] == 'c' {c--}
            left++
        }
    }
    return count
}