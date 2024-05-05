func minimumOperationsToMakeKPeriodic(word string, k int) int {
    kSizeFreq := map[string]int{}
    maxFreq := 0
    for i := 0; i+k <= len(word); i+=k {
        subStr := word[i:i+k]
        kSizeFreq[subStr]++
        maxFreq = max(maxFreq, kSizeFreq[subStr])
    }
    
    return len(word)/k - maxFreq
}