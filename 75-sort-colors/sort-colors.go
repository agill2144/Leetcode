func sortColors(nums []int)  {
    zeros, twos := 0, len(nums)-1
    i := 0
    for i <= twos {
        if nums[i] == 2 {
            nums[i], nums[twos] = nums[twos], nums[i]
            twos--
        } else if nums[i] == 0 {
            nums[i], nums[zeros] = nums[zeros], nums[i]
            zeros++
            if i < zeros {i++}
        } else {
            i++
        }
    }
}

// func sortColors(nums []int)  {
//     zeros := 0
//     ones := 0
//     twos := 0
//     for i := 0; i < len(nums); i++ {
//         if nums[i] == 0 {
//             zeros++
//         } else if nums[i] == 1 {
//             ones++
//         } else {
//             twos++
//         }
//     }
//     for i := 0; i < len(nums); i++ {
//         if zeros > 0 {
//             nums[i] = 0
//             zeros--
//         } else if ones > 0 {
//             nums[i] = 1
//             ones--
//         } else {
//             nums[i] = 2
//             twos--
//         }
//     }
    
// }