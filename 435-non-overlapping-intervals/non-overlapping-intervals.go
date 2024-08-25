func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int)bool{
        return intervals[i][0] < intervals[j][0]
    })
    /*
        [
            [1,12]
            [1,3]
            [2,3]
            [3,4]
        ]
    */
    removedCount := 0
    prev := 0 // idx of prev/last interval we decided to keep
    for i := 1; i < len(intervals); i++ {
        start, end := intervals[i][0], intervals[i][1]
        pEnd := intervals[prev][1]
        if start < pEnd {
            // we need to either remove curr or prev interval
            // therefore increment our counter thats tracking num of intervals removed
            removedCount++

            // remove the one that ends the later ( end time is further way)
            // [1,12] , [1,3] = in this one, we want to keep [1,3]
            // therefore prevIdx = i , 
            // otherwise we remove current interval, and than prev stays where it is 
            if end < pEnd {
                // current one ends before prev one
                // therefore remove prev
                prev = i
            }
        } else {
            // no overlap, prev becomes current
            prev = i
        }
    }
    return removedCount
}