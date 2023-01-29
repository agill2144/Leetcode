/*
    We can draw a line from each position and count the number of bricks we break doing.
    Then we can pick the line that broke the min amount of bricks
    
    A line is considered an answer if it breaks the minimum amounts of bricks
    Inorder to break minimum amount of brciks, we need to find where most of the bricks have a common gap
    A gap where a line can go thru without breaking any bricks
    We need to find such common line ( a line where majority of rows have a gap at )
    
    once we find a gap position and we know how many rows have the same gap at the same position 
    then we can say that, if we put a line thru this gap ( a gap thats common by MOST rows ), then we will break m-maxGapCount bricks
    m is the number of rows
    maxGapCount; is the count number of rows where each row share the same common gap
    
    if out of 6 rows, 4 of them share the same common gap, then putting line thru that gap will only break 2 bricks
    
    approach: hashmap to populate gapCount
    - gapCount will be a map[int]int ( freqMap kind of a thing )
    - key: gap position
    - value: number of rows have the same gap position
    - gap position is just a running sum for each cell in a row ( excluding last cell since we cannot draw a line after the wall ends )
    - for each row; we will
        - have a running sum = 0; sum += wall[i][j]
        - this sum is our gap position after adding a brick
        - update this gapPosition counter in gapCount map ( gapCount[gapPosition]++ )
        - while keeping a note of the most commonly / highest value of gapCount seen so far ( in the same loop )
            - i.e highest value seen so far in this freqMap
    - Once we have the highest gapCount ( a gapPosition thats common across majority of the rows )
    - then we can say that, if we put a line thru this gap ( a gap thats common by MOST rows ), then we will break m-maxGapCount bricks
    - m is the number of rows
    - maxGapCount; is the count number of rows where each row share the same common gap

time : o(m*n-1) - we visit each row and all cells except last one
space: ?
*/

func leastBricks(wall [][]int) int {
    gapCount := map[int]int{}
    m := len(wall)
    highestGapCount := 0
    for i := 0; i < m; i++ {
        gapTotal := 0
        for j := 0; j < len(wall[i])-1; j++ {
            gapTotal += wall[i][j]
            gapCount[gapTotal]++
            if count := gapCount[gapTotal]; count > highestGapCount {
                highestGapCount = count
            }
        }
    }
    return m-highestGapCount
}