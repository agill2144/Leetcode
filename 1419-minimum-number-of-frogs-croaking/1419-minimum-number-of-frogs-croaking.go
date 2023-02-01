func minNumberOfFrogs(croakOfFrogs string) int {
    if len(croakOfFrogs) == 0 {return 0}
    c := 0; r := 0; o := 0; a := 0; k := 0;
    // croak
    active := 0
    count := 0
    for i := 0; i < len(croakOfFrogs); i++ {
        char := croakOfFrogs[i]
        if char == 'c' {
            c++
            active++
            count = max(count, active)
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
            active--
        }
    }
    if k != c {return -1}
    return count
    
}

func max(x, y int) int {
    if x > y {return x}
    return y
}