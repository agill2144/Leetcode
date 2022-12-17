func lengthOfLastWord(s string) int {
    wordStarted := false
    right := -1
    left := 0
    for i := len(s)-1; i >= 0; i-- {
        if s[i] == 32 {
            if wordStarted {
                left = i+1
                break
            }
            continue
        }
        if !wordStarted {
            wordStarted = true
            right = i
        }
        
        
    }
    return right-left+1
}