func minChanges(s string) int {
    n := len(s)
    total := 0
    left := 0
    for i := 0; i < n; i++ {
        // create a partition of size 2 
        // and if zeros != ones, we need to make a change
        // therefore increment our total counter
        if i-left+1 == 2 {
            if s[i] != s[i-1] {total++}
            left = i+1
        }
    }
    return total
}

/*
    10011011
    00011111
*/
