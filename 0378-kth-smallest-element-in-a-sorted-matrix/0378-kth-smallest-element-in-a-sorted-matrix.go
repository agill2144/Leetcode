/*
    approach: brute force
    - created 1d array
    - sort 1d array
    - return k-1 element
*/
func kthSmallest(matrix [][]int, k int) int {
    m := len(matrix)
    n := len(matrix[0])
    merged := []int{}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            merged = append(merged, matrix[i][j])
        }
    }
    sort.Ints(merged)
    return merged[k-1]
}