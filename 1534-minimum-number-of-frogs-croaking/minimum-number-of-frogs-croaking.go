func minNumberOfFrogs(croakOfFrogs string) int {
    active := 0
    count := 0
    c := 0; r := 0; o := 0; a := 0; k := 0;
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
    if c != k {return -1}
    return count
}