// time = o(log 10^9 ) * o(m log n)
// median == middle element
// half on left and half on right
// there are dupes, include median in the countOnLeft
// median must have countOnLeft > half 
// countOnLeft counts number of elements <= median
func matrixMedian(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    left := 1
    right := 1000000000
    total := m*n
    half := total/2
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        // is mid our median?
        // mid can be our median ONLY if it has more than half elements on its left side 
        // i.e number of elements <= mid are > half
        // half = totalElements/2
        // why > half ?, shouldn't this be == half ?
        // no, because we are using <= operator to count on left of mid ( INCLUDING MID ITSELF )
        // if we say that there are supposed to be half elements on left of median
        // if we INCLUDED median in the count, then it will ALWAYS BE > half
        // why are we including counting of median ?
        // because there are dupes!
        count := countOnLeft(grid, mid)
        if count > half {
            // we are looking at a potential median, save it and search left
            ans = mid
            right = mid-1
        } else {
            // this means, number of elements <= mid are not even half!
            // this cannot be median for sure, our median / middle is somewhere on the right
            // therefore search right
            left = mid+1
        }
    }
    return ans
}

// time = o(m * logn)
func countOnLeft(grid [][]int, target int) int {
    count := 0
    m := len(grid)
    n := len(grid[0])
    for i := 0; i < m; i++ {
        // because each row is sorted, we need to find the closest/last idx to target
        // from the left side of target
        //                           closestIdx
        // |----------------------------------|--|----------------------|
        //                                       Target
        // how will idx tell us the count ?
        // count = idx+1
        idx := -1
        left := 0
        right := n-1
        for left <= right {
            mid := left + (right-left)/2
            if grid[i][mid] <= target {
                idx = mid
                left = mid+1
            } else {
                right = mid-1
            }
        }
        count += (idx+1)
    }
    return count
}