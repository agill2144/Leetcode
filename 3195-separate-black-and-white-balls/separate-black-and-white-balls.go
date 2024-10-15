func minimumSteps(s string) int64 {
    var total int64
    left := 0
    for i := 0; i < len(s); i++ {
        if s[i] == '0' {
            total += (int64(i)-int64(left))
            left++
        }
    }
    return total
}