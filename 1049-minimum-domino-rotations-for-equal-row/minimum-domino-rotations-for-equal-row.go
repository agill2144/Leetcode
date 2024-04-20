func minDominoRotations(tops []int, bottoms []int) int {
    n := len(tops)
    freq := map[int]int{}
    target := -1
    for i := 0; i < n; i++ {
        t, b := tops[i], bottoms[i]
        freq[t]++
        if freq[t] >= n { target = t; break }
        if t == b {continue}
        freq[b]++
        if freq[b] >= n { target = b; break }
    }
    // fmt.Println(target, freq)
    if target == -1 {return -1}
    if tops[0] != target && bottoms[0] != target {return -1}

    tCount, bCount := 0, 0
    for i := 0; i < n; i++ {
        t, b := tops[i], bottoms[i]
        if t != target {tCount++}
        if b != target {bCount++}
    }
    return min(tCount, bCount)
}