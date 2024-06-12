// ptrs on each end responsible for collecting their respective values 
// single pass, o(n) time, o(1) space
func sortColors(nums []int)  {
    n := len(nums)
    z, t, i := 0, n-1, 0
    for i <= t {
        if nums[i] == 2 {
            nums[i], nums[t] = nums[t], nums[i]
            t--
        } else if nums[i] == 0 {
            nums[i], nums[z] = nums[z], nums[i]
            z++
            if i < z {i++}
        } else {
            i++
        }
    }
}

// count and sort; 2 pass
// func sortColors(nums []int)  {
//     zeros, ones, twos := 0, 0, 0
//     for i := 0; i < len(nums); i++ {
//         if nums[i] == 0 {zeros++}
//         if nums[i] == 1 {ones++}
//         if nums[i] == 2 {twos++}
//     }
//     for i := 0; i < len(nums); i++ {
//         if zeros != 0 {
//             nums[i] = 0
//             zeros--
//         } else if ones != 0 {
//             nums[i] = 1
//             ones--
//         } else {
//             nums[i] = 2 
//             twos--
//         }
//     }
    
// }