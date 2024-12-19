/*
    - numbers are from 0 to n-1
    - if the array was sorted, each number would map to its own idx position
        - 0 -> 0, 1 -> 1, 2 -> 2, etc        
    - a subarray in range [i,j] should contain the values [i,j]
    - thats when a subarray could be sorted
    - if we are nums[i], for this ith num to be in its correct pos
    - the correct idx would be at nums[i] (i.e the value itself)
    - for a number to be in its correct position, we need to go to that idx position
    - thats when we can create a chunk!
    - but what if our example looks like this

     0 1 2 3 4 5
    [0,4,5,2,3,1]

    - 0 is at the correct position, create a chunk
    - 4 is supposed to be at idx 4, so then we jump to idx 4 and create a chunk
        - this chunk included [4,5,2,3]
        - we can place 4 correctly
        - but now 5 is in the wrong place
    - so we cannot skip and jump forward
    - we need to respect each element's position
    - we can keep a max number seen so far
    - if i == max number, we know this max number belongs at this idx, create a chunk

     0 1 2 3 4 5
    [0,4,5,2,3,1]

    - lets say max val is now 4
    - then we go to 5, max val is now 5
    - does i == maxVal ? ( 2 == 5 ) nope
    - i goes all the till idx 5, and max is still 5
    - does i == maxVal ? ( 5 == 5 ), yes, create a chunk
    - for the maxVal to be at its correct position, it needs to be at $maxVal idx
    - this means all the numbers in this partition can be sorted
        - [4,5,2,3,1] -> [1,2,3,4,5]

    tc = o(n)
    sc = o(1)
*/
func maxChunksToSorted(arr []int) int {
    maxVal := -1
    count := 0
    for i := 0; i < len(arr); i++ {
        maxVal = max(maxVal, arr[i])
        if maxVal == i {count++}
    }
    return count
}