func lengthOfLastWord(s string) int {
    n := len(s)
    start := n-1
    for start >= 0 && s[start] == ' ' {start--;}
    end := start
    for start >= 0 && s[start] != ' ' {
        start--
    }    
    return end-start
}