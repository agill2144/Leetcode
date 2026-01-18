func sortColors(nums []int)  {
    z,t := 0, len(nums)-1
    i := 0
    for i <= t {
        if nums[i] == 0 {
            nums[z], nums[i] = nums[i], nums[z]
            z++
            i++
        } else if nums[i] == 2 {
            nums[t],nums[i] = nums[i], nums[t]
            t--
        } else {
            i++
        }
    }
}
// lol
// func sortColors(nums []int)  {
//     sort.Ints(nums)
// }

// because there are only 3 values possible, we can use counting sort
// where we count the occurence of each val in its own counter var
// then write it back in the asc order
// tc = o(2n) = o(n)
// sc = o(1)
// func sortColors(nums []int)  {
// 	z,o,t := 0,0,0
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] == 0 {z++}
// 		if nums[i] == 1 {o++}
// 		if nums[i] == 2 {t++}
// 	}
// 	for i := 0; i < len(nums); i++ {
// 		if z > 0 {nums[i] = 0; z--; continue}
// 		if o > 0 {nums[i] = 1; o--; continue}
// 		if t > 0 {nums[i] = 2; t--}
// 	}

// }