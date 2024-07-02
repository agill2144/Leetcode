func longestMountain(arr []int) int {
    // simulate it
    i := 0
    maxDist := 0
    for i < len(arr) {
        
        // walk up the mountain
        start := i
        for i+1 < len(arr) && arr[i+1] > arr[i] {
            i++
        }
        if start == i {i++; continue}

        peakIdx := i
        // i at some peak val from left hand side
        // i+1 should be going downhill now
        // which means we should start walking down
        for i+1 < len(arr) && arr[i] > arr[i+1] {
            i++
        }
        if i == peakIdx {i++; continue}

        dist := i-start+1
        maxDist = max(maxDist, dist)
    }
    return maxDist
}