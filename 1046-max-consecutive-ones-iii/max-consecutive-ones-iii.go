/*
    sliding window nested loop iterations:
    - can we use a non-shrinking window?
    - if window is invalid, oh well
        - if left ptr is at a 0, reduce zeros counter
        - move left 
    - if our window is valid, take its size and save it as needed
    time = o(n)
    space = o(1)
*/
func longestOnes(nums []int, k int) int {
    n := len(nums)
    maxWindow := 0
    zeros := 0
    left := 0
    for i := 0; i < n; i++ {
        if nums[i] == 0 {
            zeros++
        }
        if zeros > k {
            // invalid window
            if nums[left] == 0 {zeros--}
            left++
        } else {
            // valid window
            maxWindow = max(maxWindow, i-left+1)
        }
    }
    return maxWindow
}
/*
    sliding window nested loop iterations:
    - can we make left ptr take bigger jumps
    - yes, we would have maintain of a list of zeros idx positions
    - we cannot use a idx map because once we escape a earliest seen zero, what about its next position?
    - hence we need a list of idx positions 

    time = o(n)
    space = o(n) ; we are using a queue to store idx positions of zeros
*/
// func longestOnes(nums []int, k int) int {
//     zerosQ := []int{}
//     maxWindow := 0
//     left := 0
//     for i := 0; i < len(nums); i++ {
//         if nums[i] == 0  {
//             zerosQ = append(zerosQ, i)
//         }
//         if len(zerosQ) > k {
//             // invalid window, escape the earliest 0 we saw
//             // i.e first zero we saw ( hence using queue, FIFO )
//             earliest := zerosQ[0]
//             if left <= earliest {
//                 left = earliest + 1
//             }
//             zerosQ = zerosQ[1:]
//         }

//         // now we have a valid window
//         maxWindow = max(maxWindow, i-left+1)
//     }
//     return maxWindow
// }

/*
    brute force:
    - form all possible subarrays
    - while allowing at most k zeros
    - nested iterations
    time = o(n^2)
    space = o(1)

    sliding window:
    - subarrays / consecutive == sliding window
    - maintain a counter for zeros
    - if our window is invalid (zeros > allowed zeros )
        - keep shrinking our window 1 by 1 from left
        - if left int is a 0, decrement zeros counter
    - once we have a valid window, take its size, save it as needed

    time = o(2n) = o(n)
    space = o(1)
*/
// func longestOnes(nums []int, k int) int {
//     n := len(nums)
//     maxWindow := 0
//     left := 0
//     zeros := 0
//     for i := 0; i < n; i++ {
//         if nums[i] == 0 {
//             zeros++
//         }
//         for zeros > k {
//             if nums[left] == 0 {
//                 zeros--
//             }
//             left++
//         }

//         maxWindow = max(maxWindow, i-left+1)
//     }
//     return maxWindow
// }