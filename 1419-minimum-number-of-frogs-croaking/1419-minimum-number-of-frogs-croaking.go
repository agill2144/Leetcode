func minNumberOfFrogs(croakOfFrogs string) int {
    c := 0
    r := 0
    o := 0
    a := 0
    k := 0
    activeFrogs := 0
    ans := -1
    
    for i := 0; i < len(croakOfFrogs); i++ {
        char := croakOfFrogs[i]
        if char == 'c' {
            c++
            activeFrogs++
            ans = max(ans, activeFrogs)

        } else if char == 'r' {
            r++
        } else if char == 'o' {
            o++
        } else if char == 'a' {
            a++
        } else if char == 'k' {
            k++
            activeFrogs--
        }
        if r > c || o > r || a > o || k > a {return -1}
    }
    if activeFrogs != 0 {return -1}
    return ans
}

func max(x, y int) int {
    if x > y {return x}
    return y
}