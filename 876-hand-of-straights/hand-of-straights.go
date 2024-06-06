func isNStraightHand(hand []int, groupSize int) bool {
    n := len(hand)
    if n % groupSize != 0 {return false}
    freq := map[int]int{}
    start := math.MaxInt64
    end := math.MinInt64
    for i := 0; i < len(hand); i++ {
        freq[hand[i]]++
        start = min(start, hand[i])
        end = max(end, hand[i])
    }
    rollingStartPtr := start
    ptr := rollingStartPtr
    currGrpSize := 0
    lastGrpNum := math.MinInt64
    for ptr <= end {
        currNum := ptr
        count := freq[currNum]
        if currGrpSize >= 1 && count == 0 {return false}
        if count >= 1 {
            if currGrpSize > 0 {
                // we have a grp in-progress
                if lastGrpNum+1 != currNum {return false}
                lastGrpNum = currNum
                currGrpSize++
                
            } else {
                // we do not have a grp in-progress
                // i.e this is a new grp
                lastGrpNum = currNum
                currGrpSize = 1
            }
            freq[currNum]--
            if freq[currNum] == 0 {delete(freq, currNum)}
                if count == 1 {rollingStartPtr++}

        } else {
            if currGrpSize >= 1 {return false}
        }
        // we made a grp
        if currGrpSize == groupSize {
            ptr = rollingStartPtr
            currGrpSize = 0
            lastGrpNum = math.MinInt64
            continue
        }
        ptr++
    }
    
    return len(freq) == 0
}