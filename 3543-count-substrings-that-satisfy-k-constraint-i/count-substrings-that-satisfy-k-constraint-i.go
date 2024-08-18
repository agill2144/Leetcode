func countKConstraintSubstrings(s string, k int) int {
    count := 0
    ones := 0
    zeros := 0
    left := 0
    for i := 0; i < len(s); i++ {
        if s[i] == '1' {ones++}
        if s[i] == '0' {zeros++}
        for ones > k && zeros > k {
            if s[left] == '1' {ones--}
            if s[left] == '0' {zeros--}
            left++
        }
        count += (i-left+1)
    }
    return count
}