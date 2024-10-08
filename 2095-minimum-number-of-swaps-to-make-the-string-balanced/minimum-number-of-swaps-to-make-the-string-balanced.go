func minSwaps(s string) int {
    count := 0
    for i := 0; i < len(s); i++ {
        if s[i] == '[' {
            count++
        } else {
            count--
            if count < 0 {count=0}
        }
    }
    return (count+1)/2
}