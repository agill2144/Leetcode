/*
    binary search on values ( smallest value at 0,0 and largest value at m-1,n-1 )
    each mid is being evaluated, are you my kth value?
    for mid to be the kth value, our matrix must contain atleast k or more elements less than or equal to mid value
    if count of lessThanOrEqual to mid >= k, this is a potential kth value, however we want the smallest one, therefore search left ( we may have overshot )
    if count of lessThanOrEqual to mid < k, search right
    for mid to be a kth value, there must be atleast k elements that are <= mid

    time: o(log$maxNumber) for binary search on values * o(m+n) for stair case search
    space: o(1)

*/
func kthSmallest(matrix [][]int, k int) int {
    m := len(matrix)
    n := len(matrix[0])
    
    left := matrix[0][0]
    right := matrix[m-1][n-1]
    kth := -1    
    for left <= right {
        mid := left+(right-left)/2
        // for mid to be our kth smallest from the starting range
        // there must be atleast k elements to the left of this number
        // there must be atleast k elements less than or equal to this number
        numElementsLessThanMid := countLessThanOrEqualTo(matrix, mid)
        
        // number of elements less than or equal to mid can be k (no dupes) or more than k (because of dupes)
        // as long as the number of elemnts to the left of this mid number is >= k , than this means, we have 
        // ATLEAST k numbers to the left of this mid number, save this as a potential answer, and search left
        // we want the smallest possible such mid
        if numElementsLessThanMid >= k {
            kth = mid
            right = mid-1
        } else {
            // when the number of elements to the left of mid is less than k
            // that means this mid number on our number line is not at the kth position
            // therefore search right
            left = mid+1
        }
    }
    return kth
}

// use stair case search to find count <= target
func countLessThanOrEqualTo(matrix [][]int,target int) int {
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