func validWordAbbreviation(word string, abbr string) bool {
    a, w := 0, 0
    n := 0
    for a < len(abbr) && w < len(word) {
        for a < len(abbr) && abbr[a] >= '0' && abbr[a] <= '9' {
            if abbr[a] == '0' && n == 0 {return false}
            n = n * 10 + int(abbr[a]-'0')
            a++
        }
        // whatever n we have, move w that many times
        w += n
        // since we have used all of n, it can be reset to 0
        n = 0

        // if either of them go out of bounds,
        // break and re-eval n 
        if a == len(abbr) || w == len(word) {break}

        // happy path, when both ptrs are inbound, compare and return false if they dont match
        if word[w] != abbr[a] {return false}
        a++
        w++
    }
    // move w ptr forward n times ( incase n was at the last position )
    // w could out of bounds, and thats fine!
    // we want n and w ptr to balanced
    // meaning in a true case, w should never go out of bound regardless of what the n value is
    // so blindly move w forward n times
    w += n
    n = 0
    // now check, w and a should be at the lastIdx+1 position, 
    // w could out of bounds if n was greater than remaining chars left
    // meaning the n in abbr was too big and was hence invalid!
    return w == len(word) && a == len(abbr)
}