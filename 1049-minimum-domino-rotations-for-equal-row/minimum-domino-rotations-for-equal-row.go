func minDominoRotations(tops []int, bottoms []int) int {    
    desired := tops[0]
    tSwaps := 0
    bSwaps := 0
    possible := true
    for i := 0; i < len(tops); i++ {
        if tops[i] != desired && bottoms[i] != desired {possible = false; break}
        if tops[i] != desired && bottoms[i] == desired {tSwaps++}
        if bottoms[i] != desired && tops[i] == desired {bSwaps++}
    }
    if possible {return min(tSwaps, bSwaps)}

    desired = bottoms[0]
    tSwaps = 0
    bSwaps = 0
    possible = true    
    for i := 0; i < len(tops); i++ {
        if tops[i] != desired && bottoms[i] != desired {possible = false; break}
        if tops[i] != desired && bottoms[i] == desired {tSwaps++}
        if bottoms[i] != desired && tops[i] == desired {bSwaps++}
    }
    if possible {return min(tSwaps, bSwaps)}
    return -1
}