func minNumberOfFrogs(croakOfFrogs string) int {
    c,r,o,a,k := 0,0,0,0,0
    maxCount := 0
    for i := 0; i < len(croakOfFrogs); i++ {
        char := croakOfFrogs[i]
        if char == 'c' {
            c++
            if c > maxCount {maxCount = c}
        } else if char == 'r' {
            r++
            // c-r-o-a-k
            // how can there be more r's than c's ?
            // its impossible and out-of-order...
            if r > c {return -1}
        } else if char == 'o' {
            o++
            if o > r {return -1}
        } else if char == 'a' {
            a++
            if a > o {return -1 }
        } else if char == 'k' {
            k++
            if k > a {return -1}
            k--; a--; o--; r--; c--;
        } else {
            // we ran into a char thats not part of "croak"
            return -1
        }
    }
    // if the string was perfectly balanced
    // then all chars in "croak" should have been finished
    // nothing should be > 0 or < 0 - for a perfect valid input
    if c != 0 || r != 0 || o != 0 || a != 0 || k != 0 {return -1}
    return maxCount
}