func minNumberOfFrogs(croakOfFrogs string) int {
    c := 0
    r := 0
    o := 0
    a := 0
    k := 0
    potentialFrogs := 0
    out := -1
    
    for i := 0; i < len(croakOfFrogs); i++ {
        char := croakOfFrogs[i]
        if char == 'c' {
            c++
            potentialFrogs++
            if potentialFrogs > out {
                out = potentialFrogs
            }
        } else if char == 'r' {
            r++
            if r > c {return -1}
        } else if char == 'o' {
            o++
            if o > r {return -1}
        } else if char == 'a' {
            a++
            if a > o {return -1}
        } else if char == 'k' {
            k++
            if k > a {return -1}
            potentialFrogs--
        }
    }
    if potentialFrogs != 0 {return -1}
    return out
}