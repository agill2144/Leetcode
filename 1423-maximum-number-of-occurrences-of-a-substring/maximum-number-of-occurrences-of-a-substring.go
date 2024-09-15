func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
    freq := map[byte]int{}
    left := 0
    wordsFreq := map[string]int{}
    maxWordsFreq := 0
    for i := 0; i < len(s); i++ {
        freq[s[i]]++
        for len(freq) > maxLetters {
            freq[s[left]]--
            if freq[s[left]] == 0 {delete(freq, s[left])}
            left++
        }
        size := i-left+1
        for size >= minSize {
            subStr := s[left:i+1]
            wordsFreq[subStr]++
            if wordsFreq[subStr] > maxWordsFreq {
                maxWordsFreq = wordsFreq[subStr]
            }
            freq[s[left]]--
            if freq[s[left]] == 0 {delete(freq, s[left])}
            left++
            size = i-left+1
        }
    }
    return maxWordsFreq
}