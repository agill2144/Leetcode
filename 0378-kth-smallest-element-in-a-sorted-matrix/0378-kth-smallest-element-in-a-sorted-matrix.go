/*
time: o(log$maxNumber) for binary search on values * o(m+n) for stair case search
*/
func kthSmallest(matrix [][]int, k int) int {
    // binary search on values ( smallest valie at 0,0 and largest value at m-1,n-1 )
    // each mid is being evaluated, are you my kth value?
    // for mid to be the kth value, our matrix must contain atleast k or more elements less than or equal to mid value
    // if count of lessThanOrEqual to mid >= k, this is a potential kth value, however we want the smallest one, therefore search left ( we may have overshot )
    // if count of lessThanOrEqual to mid < k, search right
    // for mid to be a kth value, there must be atleast k elements that are <= mid
    
    m := len(matrix)
    n := len(matrix[0])
    left := matrix[0][0]
    right := matrix[m-1][n-1]
    kthSmallestVal := -1
    
    for left <= right {
        mid := left+(right-left)/2
        count := countLessThanOrEqualTo(mid, matrix)
        if count >= k {
            kthSmallestVal = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return kthSmallestVal
}

// use stair case search to find count <= target
func countLessThanOrEqualTo(target int, matrix [][]int) int {
    m := len(matrix)
    n := len(matrix[0])
    count := 0
    r := m-1
    c := 0
    for r >= 0 && c < n {
        if matrix[r][c] <= target {
            // all elements above this cell are less than or equal to target
            // therefore add all elements in this col
            count += r+1
            // and eval the next col
            c++
        } else {
            r--
        }
    }
    return count
}