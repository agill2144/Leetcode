func removeDigit(number string, digit byte) string {
    max := ""
    for i := 0; i < len(number); i++ {
        if number[i] == digit {
            tmp := number[:i]+number[i+1:]
            if max == "" || strings.Compare(tmp, max) == 1 {
                max = tmp
            } 
        }
    }
    return max
}
    