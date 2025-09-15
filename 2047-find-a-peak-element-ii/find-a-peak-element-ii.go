/*
    - random binary search on rows or cols
    - in this case, i am doing it on rows
    - then once we have selected a row (mid)
    - we find the max value in that row 
        - selecting a cell in such manner
        - ensures we have the highest value in this row
        - solves for left / right neighor comparisons
            - because we picked the highest val in that row
    - now we have a cell and a row(mid) , we need to see if this cell is > above and bottom cell
    - if it is, we return and call it done
    - if the above cell val > current cell val
        - take the search above ( right = mid-1 )
    - if the bottom cell val > current cell val
        - take the search below ( left = mid+1 )

    m = len(rows)
    n = len(cols)
    tc = o(logm * n)
    sc = o(1)
*/
func findPeakGrid(mat [][]int) []int {
    m := len(mat)
    n := len(mat[0])
    left := 0
    right := m-1
    for left <= right {
        mid := left + (right-left)/2
        // find max in this row
        maxVal := mat[mid][0]
        maxValIdx := 0
        for j := 1; j < n; j++ {
            if mat[mid][j] > maxVal {
                maxVal = mat[mid][j]
                maxValIdx = j
            }
        }

        // now compare with above and bottom cell
        above := math.MinInt64
        if mid-1 >= 0 {above = mat[mid-1][maxValIdx]}
        bottom := math.MinInt64
        if mid+1 < m {bottom = mat[mid+1][maxValIdx]}
        curr := maxVal
        if curr > above && curr > bottom {
            return []int{mid, maxValIdx}
        }
        if above > curr {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return nil
}