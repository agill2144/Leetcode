/*
    approach: greedy
    - consider current block ( block of elements that we can jump from )
    - find the next block by finding the farthestIdx that we can jump to
    
    [x, x,x,x, x,x,x,x]
        |____|

    - assume current block
    - we have 3 elements that we can use to find our next block
    - where would the next block start at?
        - it would start at next position after the ending of current block
    - where would the next block end at?
        - THIS IS WHAT WE WANT TO MAXIMIZE!
        - loop thru current elements in current block
        - and check where can we jump the farthest from
        - thats the end of our next block 
    - eg:
    [x,x,x,x, x,x,x,x]
              |_____|

    - each block == 1 jump
    time = o(n)
    space = o(1)

*/
func jump(nums []int) int {
    n := len(nums)
    if n <= 1 {return 0}

    // curr block starts at left ptr
    left := 0
    // curr block ends at right ptr
    right := 0

    // counter to track number of jumps taken so far
    jumps := 0

    // we keep going until right ptr has reached or crossed last idx
    for right < n-1 {
        // now we need to find the next block
        // the start of next block is the next position where current block ends at ( i.e right + 1)
        nextStart := right+1
        nextEnd := -1 // we want to find the farthest idx we can reach

        // our possible elements we can jump from are the elements in our current block
        for i := left; i <= right; i++ {
            nextEnd = max(nextEnd, i+nums[i])
        }
        // update our block ptrs to next block
        left = nextStart
        right = nextEnd
        jumps++
    }
    return jumps
}