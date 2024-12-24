/*
    approach: sorted - something with binary search or stair case search
    - pretend we have a sorted array with distinct elements
    - and we want kth element
    - that element is at k-1 idx position
    - k = 5; our kth element is at k-1 idx; i.e 4
    - if our element is at 4th idx, how many elements exists before it? INCLUDING itself?
        - k elements
        - if our idx is 4, then number of elements to the left of this idx including itself is 5 (i.e k elements)
    
    - now we have dupes in our array
       0,1,2,3,4,5,6,7,8,9
    - [1,2,3,3,3,3,3,3,3,5]
    - we want k = 5
    - which is at idx 4, and that is 3
    - how many elements are <= 3?
        - 9 elements
    - so count num of elements <= 3 can be > k ... clearly with dupes
    - can count num of elements also be == k ?
       0,1,2,3,4,5,6
    - [1,2,3,4,5,6,7]; k = 2
    - meaning element at idx k-1; which is idx 1; which is 2
    - num of elements <= 2 are 2
    - so yes, count num of elements < kth element can be >= k
        - count be == k
        - count can > k
        - i.e count count be >= k
    - can count number of elements that are <= k-1 idx < k ?
        - never!
           0,1,2,3,4,5,6
        - [1,2,3,4,5,6,7]; k = 2
        - idx = k-1; 2-1 = 1;
        - therefore our kth element is 2
        - num of elements <= 2 are 2 ( which is exactly k )
        - could be < k ?
           0,1,2,3,4,5,6
        - [1,1,3,4,5,6,7]; k = 2
        - now our k-1 idx is 1
        - count <= 1 = 2 ( again exactly, k)
    - similarly, we can apply the above to our sorted matrix using binary search
    - smallest value is at matrix[0][0] and largest value is at [m-1][n-1]
    - calc mid and find number of elements <= mid
        - becuase the matrix is row wise + col wise sorted
        - we can use stair case search to count
        - pick a corner (0,n-1) or (m-1, 0)
        
    - mid is our potential kth element is count num of elements <= mid are >= k
    - there mid is our potential ans if count >= k
        - save this ans
        - but keep searching left as we may have overshot
    - otherwise if count < k
        - this mid can never be our answer
        - and all elements on the left of mid can never be our answer either
        - therefore search right
    
    tc = o(logmn * (m+n))
    - logmn for binarch search of m*n matrix
    - for each mid, we run a stair case search which runs at m+n time
    sc = o(1)
*/
func kthSmallest(matrix [][]int, k int) int {
    n := len(matrix)
    left := matrix[0][0]
    right := matrix[n-1][n-1]
    res := -1
    for left <= right {
        mid := left + (right-left)/2
        count := countLessThanOrEqualTo(matrix, mid)
        if count >= k {
            res = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return res
}

func countLessThanOrEqualTo(matrix [][]int, target int) int {
    n := len(matrix)
    r := n-1; c := 0
    count := 0
    for r >= 0 && c < n {
        if matrix[r][c] <= target {
            // entire col above this cell is <= target
            // num of elements above this cell including itself is r+1
            count += (r+1)
            // go to next col
            c++
        } else {
            r--
        }
    }
    return count
}