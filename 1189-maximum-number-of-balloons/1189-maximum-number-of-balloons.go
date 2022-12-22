func maxNumberOfBalloons(text string) int {
    b := 0; a := 0; l := 0; o := 0; n := 0
    for i := 0; i < len(text); i++ {
        if text[i] == 'b' {b++}
        if text[i] == 'a' {a++}
        if text[i] == 'l' {l++}
        if text[i] == 'o' {o++}
        if text[i] == 'n' {n++}
    }
    l /= 2
    o /= 2
    return min(b, min(a, min(l, min(o, n))))
}

func min(x, y int) int {
    if x < y {return x}
    return y
}