func minChanges(s string) int {
    left := 0
    zero := 0
    one := 0
    total := 0
    k := 2
    for i := 0; i < len(s); i++ {
        if s[i] == '1' {one++}
        if s[i] == '0' {zero++}
        if i-left+1 == k {
            total += min(zero, one)
            left = i+1
            one = 0
            zero = 0
        }
    }
    return total
}

/*
    10011000 1010

paritionSize = numOfChnages
2 = 5
4 = 5
6 = 5
8 = 5

No matter what parition size we pick, 
the total num of changes will all be the same
therefore we can pick an parition of size 2
because 2 size will always exist ( smallest even len )
*/