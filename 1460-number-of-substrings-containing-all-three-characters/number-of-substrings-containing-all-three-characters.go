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
            // we have the smallest window from left:i that contains a,b,c
            // if we take a count from left -> i ( count += (i-left+1))
            // this will count substrings within left -> i
            // which reduces the window each time and therefore making invalid windows
            // instead once we have a valid window, this is actually the smallest 
            // possible window we have , therefore how many substrings on the right
            // of this window can be created?
            // meaning, if the prefix string [left:i] is always included,
            // how many substrings can we create from right to this window?
            // count num of substring from right to i ( instead of classic left -> i)
            // i.e n-1-i+1, we can drop -1 + 1 = so it becomes n-1
            total += (n-i)
            if s[left] == 'a' {a--}
            if s[left] == 'b' {b--}
            if s[left] == 'c' {c--}
            left++
        }
    }
    return total
}