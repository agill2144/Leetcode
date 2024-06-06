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

    ptr := start
    success := 0
    currGrpSize := 0
    lastGrpNum := math.MinInt64
    grp := []int{}
    for ptr <= end {
        currNum := ptr
        count := freq[currNum]
        if count >= 1 {
            if currGrpSize == 0 {
                lastGrpNum = currNum
                currGrpSize++
                grp = append(grp, currNum)
            } else {
                if lastGrpNum+1 == currNum {
                    lastGrpNum = currNum
                    currGrpSize++
                    grp = append(grp, currNum)
                } else {
                    return false
                }
            }
            freq[currNum]--
            if freq[currNum] == 0 {delete(freq, currNum)}
        } else {
            if currGrpSize >= 1 {return false}
        }
        if currGrpSize == groupSize {
            // fmt.Println("made grp: ", grp)
            grp = []int{}
            success++
            ptr = start
            currGrpSize = 0
            lastGrpNum = math.MinInt64
            continue
        }
        ptr++
    }
    
    return len(freq) == 0
}