/*
    approach: overalapping intervals pattern
    - sort ballons by their start x position
    - and maintain a prevEnd position
        - initially this is at points[0][1]
    - group as many ballons as we can in a single group
    - inorder for next ballon to be shot with the same arrow
    - the next ballon start position MUST be <= prevEnd position
        - i.e points[i][0] <= prevEnd
        - if yes, we can include this ballon in the same group
    - HOWEVER!
        - its possible that the ballon that just got included
        - is smaller width compared to prevEnd ballon
        - therefore!
        - when including a new ballon into the same group
        - we must take into consideration that if this new ballon width is smaller
        - then we must update our prevEnd position
        - we cannot keep the longer one because the group width just got smaller!
    - if a ballon is not overlapping with prev end
        - increment arrow counter
        - update prev end position to current ballon's end position
        - i.e start a new group of ballons

*/
func findMinArrowShots(points [][]int) int {
    sort.Slice(points, func(i, j int)bool{
        return points[i][0] < points[j][0]
    })
    count := 1
    prevEnd := points[0][1]
    for i := 1; i < len(points); i++ {
        start, end := points[i][0], points[i][1]
        if start <= prevEnd {
            prevEnd = min(end,prevEnd)
        } else {
            count++
            prevEnd = points[i][1]
        }
    }
    return count
}

/*
    [10,16],[2,8],[1,6],[7,12]]

    [1,6],[2,8],[7,12],[10,16]


*/