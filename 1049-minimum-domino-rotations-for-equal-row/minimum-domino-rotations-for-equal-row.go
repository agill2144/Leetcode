func minDominoRotations(tops []int, bottoms []int) int {
    if len(tops) != len(bottoms) {return -1}

    var minRotations func(target int) int
    minRotations = func(target int) int {
        tCount, bCount := 0, 0
        n := len(tops)
        for i := 0; i < n; i++ {
            t,b := tops[i], bottoms[i]
            // when t and b is not target, this is impossible
            if t != target && b != target {return -1}
            // when t == b == target; no rotations required on either side
            if t == b && t == target {continue}
            // when target is sitting at the bottom, rotate for tops
            if t != target && b == target {tCount++}
            // when target is sitting at the top, rotate for bottom
            if b != target && t == target {bCount++}
        }
        return min(tCount, bCount)
    }
    t1Count := minRotations(tops[0])
    if t1Count == -1 {
        return minRotations(bottoms[0])
    }
    return t1Count
}

// func minDominoRotations(tops []int, bottoms []int) int {
//     n := len(tops)
//     freq := map[int]int{}
//     target := -1
//     for i := 0; i < n; i++ {
//         t, b := tops[i], bottoms[i]
//         freq[t]++
//         if freq[t] >= n { target = t; break }
//         if t == b {continue}
//         freq[b]++
//         if freq[b] >= n { target = b; break }
//     }
//     if target == -1 {return -1}
//     if tops[0] != target && bottoms[0] != target {return -1}

//     tCount, bCount := 0, 0
//     for i := 0; i < n; i++ {
//         t, b := tops[i], bottoms[i]
//         if t != target {tCount++}
//         if b != target {bCount++}
//     }
//     return min(tCount, bCount)
// }