func compress(chars []byte) int {
    //               f
    // [a,3,b,c,3,c,c]
    //            s
    slow := 0
    fast := 1
    count := 1
    
    for fast < len(chars) {
        if chars[fast] == chars[fast-1] {
            count++
            fast++
        } else {
            if count != 1 {
                strCount := strconv.Itoa(count)
                for _, digit := range strCount {
                    chars[slow+1] = byte(digit)
                    slow++
                }
                slow++
                chars[slow] = chars[fast]
                count = 1
            } else if count == 1 {
                slow++
                chars[slow] = chars[fast]   
            }
            fast++
        }
    }
    if count != 1 {
        strCount := strconv.Itoa(count)
        for _, digit := range strCount {
            chars[slow+1] = byte(digit)
            slow++
        }
    }
    return slow+1
}