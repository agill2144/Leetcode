func compress(chars []byte) int {
    if chars == nil || len(chars) == 0 {return 0}
    if len(chars) == 1 {return 1}
    
    char := chars[0]
    count := 1
    slow := 0
    
    for i := 1; i < len(chars); i++ {
        currChar := chars[i]
        prevChar := chars[i-1]
        if currChar == prevChar {
            count++
        } else {
            chars[slow] = char
            slow++
            if count != 1 {
                strCount := strconv.Itoa(count)
                for _, digit := range strCount {
                    chars[slow] = byte(digit)
                    slow++
                }
            }

            count = 1
            char = chars[i]
        }
    }
    
    chars[slow] = char
    slow++
    if count != 1 {
        strCount := strconv.Itoa(count)
        for _, digit := range strCount {
            chars[slow] = byte(digit)
            slow++
        }
    }
    return slow
}