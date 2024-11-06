func minChanges(s string) int {
    n := len(s)
    ones := 0
    zeros := 0
    total := 0
    left := 0
    for i := 0; i < n; i++ {
        if s[i] == '0' {zeros++}
        if s[i] == '1' {ones++}
        if i-left+1 == 2 {
            total += min(ones, zeros)
            ones = 0
            zeros = 0
            left = i+1
        }
    }
    return total
}

/*
    10011011
    00011111
*/
