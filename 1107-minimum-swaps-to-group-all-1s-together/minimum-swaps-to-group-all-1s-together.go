func minSwaps(data []int) int {
    ones := 0
    for i := 0; i < len(data); i++ {
        if data[i] == 1 {ones++}
    }
    if ones <= 1 {return 0}
    swaps := len(data)+1
    oneCount := 0
    left := 0
    for i := 0; i < len(data); i++ {
        if data[i] == 1 {oneCount++}
        if i-left+1 == ones {
            zeros := (i-left+1)-oneCount
            swaps = min(swaps, zeros)
            if data[left] == 1 {oneCount--}
            left++
        }
    }
    return swaps
}