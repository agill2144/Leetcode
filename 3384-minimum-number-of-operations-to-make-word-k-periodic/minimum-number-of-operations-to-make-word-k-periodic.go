func minimumOperationsToMakeKPeriodic(word string, k int) int {
    kSizeFreq := map[string]int{}
    maxFreq := 0
    for i := 0; i+k <= len(word); i+=k {
        subStr := word[i:i+k]
        kSizeFreq[subStr]++
        maxFreq = max(maxFreq, kSizeFreq[subStr])
    }
    // the word is divided into k size chunks
    // for example if word size is 10, and k = 2
    // then the word is divided into 5 chunks
    // and out of those 5 chunks, we already have maxFreq number of chunks in our desired state 
    // then whatever number of chunks are left, will need to replaced, i.e thats the number of operations needed
    // for example; word size = 10, k = 2
    // how many chunks do we have ? 10/2 = 5
    // A substr of size k that is repeating the most often is 2
    // that is, out of 5 chunks, there are 2 identical substr chunks
    // that wont need to be replaced
    // so now out of 5, we already have 2 chunks good to go
    // we are left with 3, each of the chunks need to be replaced
    // thats the number of operations required
    return len(word)/k - maxFreq
}