// time : o(2n)
// func getConcatenation(nums []int) []int {
//     ans := make([]int, 2*len(nums))
//     ansPtr := 0
//     nPtr := 0
//     for ansPtr < len(ans) {
//         ans[ansPtr] = nums[nPtr%len(nums)]
//         ansPtr++
//         nPtr++
//     }
//     return ans
// }



/*
    approach:
    - create a output array of size 2*len(nums)
    - loop over nums and
    - for each ith element;
        - place this value at out[i] and out[i+len(nums)]
    
    time: o(n)
    space: o(1)
*/
func getConcatenation(nums []int) []int {
    ans := make([]int, 2*len(nums))
    for i := 0; i < len(nums); i++ {
        ans[i] = nums[i]
        ans[i+len(nums)] = nums[i]
    }
    return ans
}


