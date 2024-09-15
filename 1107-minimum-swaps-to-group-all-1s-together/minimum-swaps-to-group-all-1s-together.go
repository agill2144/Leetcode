func minSwaps(data []int) int {
    ones := 0
    for i := 0; i < len(data); i++ {
        if data[i] == 1 {ones++}
    }
    if ones == 0 {return 0}
    swaps := len(data)+1
    left := 0
    zeros := 0
    for i := 0; i < len(data); i++ {
        if data[i] == 0 {zeros++}
        if i-left+1 == ones {
            swaps = min(swaps, zeros)
            if data[left] == 0 {zeros--}
            left++
        }
    }
    if swaps == len(data)+1 {return 0}
    return swaps
}