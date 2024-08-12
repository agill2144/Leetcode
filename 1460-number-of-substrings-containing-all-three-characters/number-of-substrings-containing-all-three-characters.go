func numberOfSubstrings(s string) int {
    a := 0
    b := 0
    c := 0
    left := 0
    count := 0
    n := len(s)
    for i := 0; i < n; i++ {
        if s[i] == 'a' {a++}
        if s[i] == 'b' {b++}
        if s[i] == 'c' {c++}
        for a > 0 && b > 0 && c > 0 {
            count += ((n-1)-i+1)
            leftChar := s[left]
            if leftChar == 'a' {a--}
            if leftChar == 'b' {b--}
            if leftChar == 'c' {c--}
            left++            
        }
    }
    return count
}