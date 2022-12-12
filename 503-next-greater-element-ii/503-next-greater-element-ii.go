/*
    approach: brute force
    - using a loop for each ith element, we are looking for next greater element
    - If we are not able to find the element to the right of the ith element,
        - then we have to make a circle and start from the beginning in hopes to find an element > ith element
        - but this circular loop ONLY GOES UNTIL it meets ith idx again -- in which case it means, there is no greater element
        
    - nested for loop on doubled array len and use % n to get relative idx within array size
    - the nested for loop goes until it does not wrap around the array once and meets i again in a circle
    time: o(n^2) -- I think -- I am having a hard time prove this to myself.. help? I keep thinking its 2n.... 
    space: o(1) -- the output array does not count as extra space...
    
*/

// func nextGreaterElements(nums []int) []int {
//     out := make([]int, len(nums))
//     for i := 0; i < len(nums); i++ {
//         out[i] = -1
//     }
//     n := len(nums)
//     for i := 0; i < len(nums); i++ {
//         for j := i+1; j % n != i; j++ {
//             if nums[j%n] > nums[i] {
//                 out[i] = nums[j%n]
//                 break
//             }
//         }
//     }
//     return out
// }


/*
    approach: monotonic increasing stack
    time: o(n)
    space: o(n)
*/
func nextGreaterElements(nums []int) []int {
    out := make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        out[i] = -1
    }
    n := len(nums)
    stack := []int{}
    for i := 0; i < 2 * n; i++ {
        for len(stack) != 0 && nums[i%n] > nums[stack[len(stack)-1]] {
            top := stack[len(stack)-1]
            out[top] = nums[i%n]
            stack = stack[:len(stack)-1]
        }
        if i < n {
            stack = append(stack, i)
        }
    }
    return out
}