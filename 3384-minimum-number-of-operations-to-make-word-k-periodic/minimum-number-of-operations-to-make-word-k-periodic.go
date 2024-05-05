func minimumOperationsToMakeKPeriodic(word string, k int) int {
    kSizeFreq := map[string]int{}
    maxFreq := 0
    for i := 0; i+k <= len(word); i+=k {
        subStr := word[i:i+k]
        kSizeFreq[subStr]++
        if val := kSizeFreq[subStr]; val > maxFreq {
            maxFreq = val
        }
    }
    
    return len(word)/k - maxFreq
}