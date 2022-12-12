// time: o(m) * o(logn) = o(mlogn)
func findPeakGrid(mat [][]int) []int {
    m := len(mat)
    n := len(mat[0])
    dirs := [][]int{
        {-1,0},
        {1,0},
        {0,-1},
        {0,1},
    }
    // for each row : o(m)
    for i := 0; i < m; i++ {
        // find out if there is a peak: o(logn)
        left := 0
        right := n-1
        for left <= right {
            mid := left + (right-left)/2
            validDirs := 0
            peakCount := 0
            for _, dir := range dirs {
                r := i+dir[0]
                c := mid+dir[1]
                if r >= 0 && r < m && c >= 0 && c < n {
                    validDirs++
                    if mat[i][mid] > mat[r][c] {
                        peakCount++
                    }
                }
            }
            if peakCount == validDirs {
                return []int{i, mid}
            }
            if mid != 0 && mid != n-1 {
                if mat[i][mid-1] > mat[i][mid+1] {
                    right = mid-1
                } else {
                    left = mid+1
                }
                continue
            }
            if mid == 0 {
                left = mid+1
            } else {
                right = mid-1
            }
            
        }
        
    }
    return nil   
}