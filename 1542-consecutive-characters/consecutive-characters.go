func maxPower(s string) int {
    maxPower := 1
    runningPower := 1
    for i := 1; i < len(s); i++ {
        if s[i] == s[i-1] {runningPower++} else { runningPower = 1}
        maxPower = max(maxPower, runningPower)        
    }
    return maxPower
}