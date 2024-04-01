/*
    approach: binary search on answers
    - we know the smallest element is at 0,0
    - we know the largest element is at m-1,n-1
    - and the numbers within the matrix are within this min,max range ( guranteed )
    - so we have a range of numbers
    - ranges are sorted, therefore we can use binary search
    - how can mid be our answer?
        - if there are >= k elements to the left of mid number
        
*/
func kthSmallest(matrix [][]int, k int) int {
    n := len(matrix)
    left := matrix[0][0]
    right := matrix[n-1][n-1]
    for left <= right {
        mid := left + (right-left)/2
        count := count(matrix, mid)
        if count < k {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return left
}

func count(matrix [][]int, target int) int {
    n := len(matrix)
    r := n-1
    c := 0
    total := 0
    for r >= 0 && c < n {
        if matrix[r][c] <= target {
            total += r+1
            c++
        } else {
            r--
        }
    }
    return total
}