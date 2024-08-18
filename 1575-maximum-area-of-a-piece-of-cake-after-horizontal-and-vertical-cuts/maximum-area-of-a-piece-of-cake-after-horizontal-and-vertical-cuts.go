func maxArea(h int, w int, hCuts []int, vCuts []int) int {
    sort.Ints(hCuts)
    sort.Ints(vCuts)
    maxHDiff := max(h-hCuts[len(hCuts)-1], hCuts[0])
    for i := len(hCuts)-2; i >= 0; i-- {
        maxHDiff = max(maxHDiff, hCuts[i+1]-hCuts[i])
    }
    maxVDiff := max(w-vCuts[len(vCuts)-1], vCuts[0])
    for i := len(vCuts)-2; i >= 0; i-- {
        maxVDiff = max(maxVDiff, vCuts[i+1]-vCuts[i])
    }
    return (maxHDiff * maxVDiff) % 1000000007
}