func minDominoRotations(tops []int, bottoms []int) int {
    result := check(tops, bottoms, tops[0])
    if result != -1 { return result }
    return check(tops, bottoms, bottoms[0])
}

func check(tops, bottoms []int, target int) int {
    aRot := 0
    bRot := 0
    
    for i := 0; i < len(tops); i++ {
        topVal := tops[i]
        botVal := bottoms[i]
        if topVal != target && botVal != target {
            return -1
        }
        
        if topVal != target && botVal == target {
            aRot++
        }
        if botVal != target && topVal == target {
            bRot++
        }
    }
    return int(math.Min(float64(aRot), float64(bRot)))
}