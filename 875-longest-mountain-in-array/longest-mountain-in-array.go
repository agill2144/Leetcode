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
        // next value is smaller, therefore being a new start point over again
        // this means, start ptr we assumed as a good start point is not a good start point
        if i == start {i++; continue}

        peakIdx := i
        // i at some peak val from left hand side
        // i+1 should be going downhill now
        // which means we should start walking down
        for i+1 < len(arr) && arr[i] > arr[i+1] {
            i++
        }
        
        // if peakIdx a valid peakIdx ?
        // peakIdx valid if i ended up > peakIdx
        // or i != peakIdx
        // but if i == peakIdx, it means value next to peakIdx is not smaller, 
        // therefore this is not a good peak idx
        // start all over again
        if i == peakIdx {
            i++
            continue
        }
        // now we have a valid mountain, take it size and save it if better than last size
        dist := i-start+1
        maxDist = max(maxDist, dist)
    }
    return maxDist
}