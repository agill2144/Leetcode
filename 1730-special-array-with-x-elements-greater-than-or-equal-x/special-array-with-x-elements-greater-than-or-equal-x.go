// no sort, binary search on answers
// time = o(n) + o(logMaxNum * n)
func specialArray(nums []int) int {
    left := 1
    right := math.MinInt64
    freqMap := map[int]int{}
    for i := 0; i < len(nums); i++ {
        freqMap[nums[i]]++
        right = max(right, nums[i])
    }

    for left <= right {
        mid := left + (right-left)/2
        // count >= mid
        count := 0
        for k, v := range freqMap {
            if k >= mid {
                count+=v
            }
        }
        if count == mid {return mid}
        if count < mid {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return -1
}
// binary search on ans pattern
// because we are looping over a range in our naive approach
// ranges are sorted, therefore binary search
// mid works if count == x
// count means; number of elements >= x
// so if count < x; it means there are less number >= x
// we need to increase our count, which means we need a smaller x value
// therefore right = mid-1
// time = o(nlogn) + o(log(maxNum) * n)
// func specialArray(nums []int) int {
//     sort.Ints(nums)
//     left := 1
//     right := nums[len(nums)-1]
//     for left <= right {
//         mid := left + (right-left)/2
//         j := 0
//         for j < len(nums) {
//             if nums[j] < mid {j++; continue}
//             break
//         }
//         count := len(nums)-j
//         if count == mid {return mid}
//         if count < mid {
//             right = mid-1
//         } else {
//             left = mid+1
//         }
//     }
//     return -1
// }
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