func minDominoRotations(tops []int, bottoms []int) int {    
    desired := tops[0]
    tSwaps := 0
    bSwaps := 0
    possible1 := true
    res := math.MaxInt64
    for i := 0; i < len(tops); i++ {
        if tops[i] != desired && bottoms[i] != desired {possible1 = false; break}
        if tops[i] != desired && bottoms[i] == desired {tSwaps++}
        if bottoms[i] != desired && tops[i] == desired {bSwaps++}
    }
    if possible1 {res = min(tSwaps, bSwaps)}

    desired = bottoms[0]
    tSwaps = 0
    bSwaps = 0
    possible2 := true    
    for i := 0; i < len(tops); i++ {
        if tops[i] != desired && bottoms[i] != desired {possible2 = false; break}
        if tops[i] != desired && bottoms[i] == desired {tSwaps++}
        if bottoms[i] != desired && tops[i] == desired {bSwaps++}
    }
    if !possible1 && !possible2 {return -1}
    if possible2 {res = min(res, min(tSwaps, bSwaps))}
    return res
}