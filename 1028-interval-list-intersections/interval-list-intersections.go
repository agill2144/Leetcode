/*
    approach: two ptrs
    - input is sorted, therefore 2 patterns; two ptrs and binary search
    - there are 3 cases when two intervals overlap
        - assume a and b are the two intervals we need to check
        - a is fully inside of b / vice versa
        - a is at the edge of b / vice versa
        - a is partially inside of b / vice versa
    - so when are 2 intervals not overlapping?
        - a starts after b ends / vice versa ( b starts after a ends )
    - so when there is no overlap, which ptr do we move away from?
        - we move away from the earliest ending one
        - because that interval is fully used and wont overlap since it has ended
    - how do we capture the overalapping times?
        - last starting time ( start time that starts later )
        - earliest ending time
        - i.e max(aStart, bStart) and min(aEnd, bEnd)
        - this is the overlapping interval

    - another way to determine if there is an overlap
    - if we capture the last start time and earliest end time
    - when we have no overlap
        - end will end before start will start
    - start = max(aStart, bStart)
    - end = min(aEnd, bEnd)
    - when there is no overlap, start > end (ALWAYS!)
    - when this is not the case (i.e start <= end )
    - there is an overlap, and the overlap is [start, end] itself!

    time = o(min(a,b))
    space = o(1)
*/
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
    if len(firstList) == 0 || len(secondList) == 0 {return nil}
    out := [][]int{}
    a, b := 0, 0
    for a < len(firstList) && b < len(secondList) {
        aStart, aEnd := firstList[a][0], firstList[a][1]
        bStart, bEnd := secondList[b][0], secondList[b][1]
        // ask yourself, when does this not overalp?
        start := max(aStart, bStart)
        end := min(aEnd, bEnd)
        // when one starts AFTER the the other one ends
        // meaning, there is an overlap if start <= end
        if start <= end {
            out = append(out, []int{start, end})
        }
        // move away from earliest ending interval
        if aEnd < bEnd {a++} else {b++}
        // naive detection approach
        // if aStart > bEnd || bStart > aEnd {
        //     // move away from earliest ending interval
        //     if aEnd < bEnd {a++} else {b++}
        // } else {
        //     // there is an overlap for sure!
        //     // overlap can be caputured by last start time and earliest end time
        //     out = append(out, []int{max(aStart, bStart), min(aEnd, bEnd)})
        //     // move away from earliest ending interval
        //     if aEnd < bEnd {a++} else {b++}
        // }
    }
    return out
}