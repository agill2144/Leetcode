func totalFruit(fruits []int) int {
    freq := map[int]int{}
    left := 0
    maxWin := 0
    for i := 0; i < len(fruits); i++ {
        freq[fruits[i]]++
        for len(freq) > 2 {
            freq[fruits[left]]--
            if freq[fruits[left]] == 0 {delete(freq, fruits[left])}
            left++
        }
        maxWin = max(maxWin, i-left+1)
    }
    return maxWin
}