func reverse(x int) int {
    out := 0
    isNegative := false
    if x < 0 {isNegative = true; x *= -1}
    
    for x != 0 {
        last := x % 10
        x = x / 10
        out = out * 10 + last
        if out > math.MaxInt32 {return 0}
    }
    
    if isNegative {out *= -1}
    return out
}