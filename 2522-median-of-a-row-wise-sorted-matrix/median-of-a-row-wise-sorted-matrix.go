func matrixMedian(grid [][]int) int {
    left := 1
    right := 1000000
    ans := -1
    m := len(grid)
    n := len(grid[0])
    total := m*n
    midIdx := (total-1)/2
    for left <= right {
        mid := left + (right-left)/2
        count := countLessOrEqualTo(grid, mid)
        if count >= midIdx+1 {
            ans = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return ans
}

func countLessOrEqualTo(grid [][]int, target int) int {
    count := 0
    m := len(grid)
    n := len(grid[0])
    for i := 0; i < m; i++ {
        left := 0
        right := n-1
        idx := -1
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
/*
             
    [1,2,3,4,5,6,7,8,9] ; 
    - including median (mid idx; idx 4; val 5)
    - we have idx+1 count <= median val

    [5,5,5,5,5,5,7,8,9]
    - in array with dupes
    - median is at idx 4
    - but count may be >= idx+1

    can a median have less idx+1 count on left side?
    - no
    - [1,2,3,5,5,5,5,5]    
    - when eval'ing 3, count <= 3 ? = 3
    - median rule without dupes; count less than or equal to MUST be == midIdx+1
    - median rule with dupes; count less than or equal to MUST be  >= midIdx+1
    - midIdx = 4
    - count less than or equal to for a median must be >= midIdx+1 ( 4+1 = 5)
    - count num of elements <= 3 is 3
    - therefore 3 is not median
*/