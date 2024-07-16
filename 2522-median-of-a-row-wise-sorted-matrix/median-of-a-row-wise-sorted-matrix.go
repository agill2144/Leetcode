/*
    important underlying pattern:
    - median is always the middle element
    - given our matrix will have odd number of elements in total
    - if our matrix was flattened into a 1d array, median would be the middle idx
    - we did median with 2 sorted arrays;
        - we used binary search to place a partition such that
        - all elements on left <= all elements on right
        - once such partition was discovered, we were lookig at our middle elements (l1,l2,r1,r2)
    - another median solution is that since we only care about 2 middle elements (2 because in even sized array)
        - we could also use a maxHeap to keep track of left side of partition
        - i.e use maxHeap to always tell us the max number on right side of partition
        - and use minHeap to always tell us the min number on left side of partition
        - we would have to ensure that maxHeap and minHeap are kind of balanced
        - and that top of minHeap (right side of sorted arr) >= top of maxHeap(left side of sorted array)
    
    - this is yet another solution; binary searching on answers
    - we know our min answer is at matrix[0][0] and max answer is at matrix[m-1][n-1]
    - then we have a range of sorted integers to evaluate

    IMPORTANT
    - median in a odd size array means we would have half the number of elements to the left of median
    - left as in, number of elements < median will be half of the array
    - and thats a fact
    - if we had duplicates, then maybe number of elements <= median element are more than half

    - HOWEVER, if we say, count number of elements that are <= median value ( including median in the count )
    - WE WILL ALWAYS HAVE A COUNT > half of the size of the array
    - IT WILL BE > half when we are looking at a potential median
    - arr = [1,1,3,3,4]
        - array size = 5
        - half = 5/2 2
        - so median maybe 2... 

        - median lies within 1 and 4
        - mid = 2
        - if mid is our answer, how many elements exists are <= 2
        - number of elements <= 2 = 2
        - and we said, for a answer to median its count number of elements <= median itself SHOULD ALWAYS BE > half
        - is 2 > half (2) ? no
        - therefore search right
        - mid = 3
        - if mid is our answer, how many elements exists are <= 3
        - number of elements <= 3 = 4
        - and we said, for a answer to median its count number of elements <= median itself SHOULD ALWAYS BE > half
        - is 4 > half ? yes
        - save this potential ans and continue searching left
        - essentially we are looking for the very first answer/midValue where num of elements <= answer/midValue is > half

*/
func matrixMedian(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    ans := -1
    left := 1
    right := 1000000000
    for left <= right {
        mid := left + (right-left)/2
        count := countLessThanOrEqualTo(grid,mid)
        // for mid to be median
        // count num of elements <= mid(including mid) should always be > half
        if count > (m*n)/2 {
            ans = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return ans
}

func countLessThanOrEqualTo(grid [][]int, target int) int {
    m := len(grid)
    n := len(grid[0])
    count := 0
    for i := 0; i < m; i++ {
        left := 0
        right := n-1
        // find the right most idx on left side of target ( target may exists, find the right most position)
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
        count += idx+1
    }
    return count
}