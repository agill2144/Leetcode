// because we are looping over a range in our naive approach
// ranges are sorted, therefore binary search
// mid works if count == x
// count means; number of elements >= x
// so if count < x; it means there are less number >= x
// we need to increase our count, which means we need a smaller x value
// therefore right = mid-1
func specialArray(nums []int) int {
    sort.Ints(nums)
    left := 1
    right := nums[len(nums)-1]
    for left <= right {
        mid := left + (right-left)/2
        j := 0
        for j < len(nums)  {
            if nums[j] < mid {j++; continue}
            break
        }
        count := len(nums)-j
        if count == mid {return mid}
        if count < mid {
            right = mid-1
        } else {
            left = mid+1
        }

    }
    return -1

}
// ans lies in a range; 1 to max number
// eval each ans
// time = o(nlogn) + o(maxNum * n)
// func specialArray(nums []int) int {
//     sort.Ints(nums)
//     for i := 1; i <= nums[len(nums)-1]; i++ {
//         x := i

//         j := 0
//         for j < len(nums)  {
//             if nums[j] < x {j++; continue}
//             break
//         }
//         count := len(nums)-j
//         if count == x {return x}
//     }
//     return -1
// }