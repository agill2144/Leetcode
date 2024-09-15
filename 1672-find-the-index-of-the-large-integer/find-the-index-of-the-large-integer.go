/**
 * // This is the ArrayReader's API interface.
 * // You should not implement it, or speculate about its implementation
 * type ArrayReader struct {
 * }
 * // Compares the sum of arr[l..r] with the sum of arr[x..y] 
 * // return 1 if sum(arr[l..r]) > sum(arr[x..y])
 * // return 0 if sum(arr[l..r]) == sum(arr[x..y])
 * // return -1 if sum(arr[l..r]) < sum(arr[x..y])
 * func (this *ArrayReader) compareSub(l, r, x, y int) int {}
 * 
 * // Returns the length of the array
 * func (this *ArrayReader) length() int {}
 */

func getIndex(reader *ArrayReader) int {
    left := 0
    n := reader.length()
    right := n-1
    for left < right {
        mid := left + (right-left)/2
        cmp := -2
        if (right-left+1) % 2 == 0 {
            cmp = reader.compareSub(left, mid, mid+1, right)
        } else {
            cmp = reader.compareSub(left, mid-1, mid+1, right)
        }
        
        if cmp == 1 {
            right = mid  
        } else if cmp == -1 {
            left = mid+1
        } else {
            right = mid
        }
    }
    return left
}