// bring the largest value in smallest unit to the largest unit
// time = o(n)
// space = o(n) for nums list + o(9) for idxs array = o(n)
func maximumSwap(num int) int {
    if num <= 9 {return num}
    numStr := fmt.Sprintf("%v", num)
    nums := []int{}
    for i := 0; i < len(numStr); i++ {nums = append(nums, int(numStr[i]-'0'))}
    // last seen idxs; 
    // map each value to its idx ( val : idx ), and note down its actual idx position from input
    // if a value shows up again, thats fine, we can override its idx again ( we want the last idx , right most idx )
    // ex; 989
    // idx of 9 will be 2 at the end
    idxs := make([]int, 10) 
    // pre-fill idxs with -1 ( as if they are not yet found )
    for i := 0; i < len(idxs); i++ {idxs[i] = -1}
    for i := 0; i < len(nums); i++ {idxs[nums[i]] = i}
    
    swapped := false
    for i := 0; i < len(nums) && !swapped; i++ {
        // for each digit on the higest unit ( from left to right )
        // check for possible swaps
        // if we value 2, 3,4, we want highest possible number to swap with
        // highest possible num is 9
        // therefore start the idx loop from the back of the array
        // each idx is a value and the value in the idx array represents the last idx we've seen this value at in the input
        for j := len(idxs)-1; j > nums[i]; j-- {
            // if an idx does exist
            // and its after the ith idx,
            // then we can swap
            if idxs[j] != -1 && idxs[j] > i {
                nums[idxs[j]], nums[i] = nums[i], nums[idxs[j]]
                swapped = true
                break
            }
        }
    }
    // re-construct the number back
    res := 0
    for i := 0; i < len(nums); i++ {
        res = (res * 10) + nums[i]
    }
    return res
}
/*
    bring the largest value in smallest unit to the largest unit
    example
    2139
    9 is at the smallest unit (tens)
    and now we need a largest unit whose value is smaller than largest value (from left -> right)
    2 is smaller than 9 and is also in the largest unit ( thousands unit )
    therefore swap 2 and 9 = 9132
*/
// func maximumSwap(num int) int {
//     if num <= 9 {return num}
//     nums := []int{}
//     numStr := fmt.Sprintf("%v", num)
//     for i := 0; i < len(numStr); i++ {nums = append(nums, int(numStr[i]-'0'))}
//     n := len(nums)
//     suffix := make([]int, len(nums)) 
//     suffix[n-1] = n-1
//     for i := n-2; i >= 0; i-- {
//         maxIdx := suffix[i+1]
//         maxSoFar := nums[maxIdx]
//         rightIdx := i+1
//         rightVal := nums[rightIdx]
//         if rightVal > maxSoFar {
//             suffix[i] = rightIdx
//         } else {
//             suffix[i] = maxIdx
//         }
//     }
//     for i := 0; i < n; i++ {
//         maxOnRight := nums[suffix[i]]
//         if maxOnRight > nums[i] {
//             nums[i], nums[suffix[i]] = nums[suffix[i]], nums[i]
//             break
//         }
//     }
//     res := 0
//     for i := 0; i < len(nums); i++ {
//         res = (res * 10) + nums[i]
//     }
//     return res
// }