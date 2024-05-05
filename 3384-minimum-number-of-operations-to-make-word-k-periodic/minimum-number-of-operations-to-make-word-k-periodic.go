func minimumOperationsToMakeKPeriodic(word string, k int) int {
    kSizeFreq := map[string]int{}
    maxFreq := 0
    for i := 0; i <= len(word) - k ; i+=k {
        subStr := word[i:i+k]
        kSizeFreq[subStr]++
        if val := kSizeFreq[subStr]; val > maxFreq {
            maxFreq = val
        }
    }
    ops := math.MaxInt64
    for evalSubStr, v := range kSizeFreq {
        if v == maxFreq {
            // evaluate this substr
            evalOps := 0
            for i := 0; i <= len(word)-k; i+=k {
                subStr := word[i:i+k]
                if subStr == evalSubStr {continue}
                evalOps++ 
            }
            ops = min(ops, evalOps)
        }
    }
    

    return ops
}