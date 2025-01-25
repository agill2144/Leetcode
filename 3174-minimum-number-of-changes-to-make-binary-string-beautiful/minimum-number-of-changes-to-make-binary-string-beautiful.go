/*
    0101001010

    size; 6+4
    010100 1010
    2       2
    4 changes

    size; 2
    01 01 00 10 10    
    1. 1     1.  1
    4 changes

*/
func minChanges(s string) int {
    n := len(s)
    if n % 2 != 0 {return -1}
    k := 2    
    left := 0
    zero := 0
    one := 0
    changes := 0
    for i := 0; i < n; i++ {
        if s[i] == '0' {zero++}
        if s[i] == '1' {one++}
        if i-left+1 == k {
            changes += min(zero, one)
            one = 0
            zero = 0
            left = i+1
        }
    }
    return changes
}