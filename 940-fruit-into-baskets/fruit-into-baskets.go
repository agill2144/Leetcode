func totalFruit(fruits []int) int {
    freq := map[int]int{}
    left := 0
    maxWin := 0
    for i := 0; i < len(fruits); i++ {
        freq[fruits[i]]++
        if len(freq) <= 2 {
            maxWin = max(maxWin, i-left+1)
        } else {
            freq[fruits[left]]--
            if freq[fruits[left]] == 0 {delete(freq, fruits[left])}
            left++
        }
    }
    return maxWin
}