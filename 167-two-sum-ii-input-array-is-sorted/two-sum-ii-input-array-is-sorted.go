// sorted
// 1. binary search
// pick a number and search for missing number on right side of ith pos
// search = binary search
// tc = o(nlogn)
// sc = o(1)
func twoSum(nums []int, target int) []int {
    n := len(nums)
    for i := 0; i < n; i++ {
        t := target - nums[i]
        val := search(i+1, n-1, t, nums)
        if val != -1 {
            return []int{i+1, val+1}
        }
    }
    return nil
}

func search(left, right, target int, nums []int) int {
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            return mid
        }
        if target > nums[mid] {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return -1
}

// sorted 
// 1. two ptrs
// tc = o(n)
// sc = o(1)
// func twoSum(nums []int, target int) []int {
//     n := len(nums)
//     left := 0
//     right := n-1
//     for left < right {
//         sum := nums[left] + nums[right]
//         if sum == target {return []int{left+1, right+1}}
//         if sum > target {
//             right--
//         } else {
//             left++
//         }
//     }
//     return nil
// }

// compliment search will obv work
// tc = o(n)
// sc = o(n)
// func twoSum(nums []int, target int) []int {
//     idx := map[int]int{}
//     for i := 0; i < len(nums); i++ {
//         diff := target - nums[i]
//         val, ok := idx[diff]
//         if ok {return []int{val+1, i+1}}
//         idx[nums[i]] = i
//     }
//     return nil
// }