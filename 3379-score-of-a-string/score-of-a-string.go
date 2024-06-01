func scoreOfString(s string) int {
    total := 0
    for i := 0; i < len(s)-1; i++ {
        curr := int(s[i])
        next := int(s[i+1])
        total += abs(curr-next)
    }
    return total
}

func abs(x int) int {
    if x < 0 {return x * -1}
    return x
}