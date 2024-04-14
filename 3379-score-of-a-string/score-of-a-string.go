func scoreOfString(s string) int {
    total := 0
    for i := 1; i < len(s); i++ {
        curr := int(s[i])
        prev := int(s[i-1])
        diff := abs(curr-prev)
        total += diff
    }
    return total
}

func abs(x int) int {
    if x < 0 {return x*-1}
    return x
}