func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
    out := [][]int{}   
    a,b := 0,0
    for a < len(firstList) && b < len(secondList) {
        // overlap between 2 intervals happens when
        // the lastStartTime is still before earliestEndingTime
        // if we picked the max start time between the 2
        // and we picked the min earliest time betwee the 2
        // there shouldn't be an overlap if the earliest ending time is before the last start time
        // the inverse of above; there is an overlap if last starting time is STILL before earliest ending time!
        aStart,aEnd := firstList[a][0],firstList[a][1]
        bStart,bEnd := secondList[b][0],secondList[b][1]
        start, end := max(aStart, bStart), min(aEnd, bEnd)
        if start <= end {
            out = append(out, []int{start, end})
        }
        if aEnd < bEnd {a++} else {b++}
    }
    return out
}
// func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
//     out := [][]int{}   
//     a,b := 0,0
//     for a < len(firstList) && b < len(secondList) {
//         aStart,aEnd := firstList[a][0],firstList[a][1]
//         bStart,bEnd := secondList[b][0],secondList[b][1]
//         // when is there no intersection?
//         // when aStarts after bEnds or 
//         // when bStarts after aEnds
//         if aStart > bEnd || bStart > aEnd {
//             // move away from the one that ends the earliest!
//             // why?
//             // because definetly there wont be overlap if the earliest ending one did not overlap with curr
//             if aEnd < bEnd {a++} else {b++}
//         } else {
//             // there is an overlap
//             // the common intersection will always
//             // be [$lastStartingTime, $earliestEndingTime]
//             // you can dry run in all 6 types of overlap
//             out = append(out, []int{max(aStart,bStart), min(aEnd,bEnd)})
//             // move away from the one that ends the earliest!
//             // why?
//             // because definetly there wont be overlap if the earliest ending one did not overlap with curr
//             if aEnd < bEnd {a++} else {b++}
//         }
//     }
//     return out
// }