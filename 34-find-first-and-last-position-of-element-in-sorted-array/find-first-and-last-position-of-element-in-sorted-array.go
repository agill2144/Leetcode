/*
    approach: binary search
    - use binary search to find the first position of our target ( i.e left most idx )
    - use another binary to find the end position of our target ( use left idx from above )
        - if start idx was not found, it means target does not exist
        - which means exit early and we dont have to run 2nd binary search

    time = o(logn) + o(logn)
    space = o(1)
*/
func searchRange(nums []int, target int) []int {
    start, end := -1, -1

    // find start idx
    left := 0
    right := len(nums)-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] >= target {
            if nums[mid] == target {
                start = mid
            }
            right = mid-1
        } else {
            left = mid+1
        }
    }
    if start == -1 {return []int{start, end}}

    // find end idx
    left = start
    right = len(nums)-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            end = mid
            left = mid+1
        } else if target < nums[mid] {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return []int{start, end}
}


/*
    approach: linear search
    - as soon as we run into target
    - check if start idx is set
        - if not, save this idx as start idx
    - and keep saving this same idx also as end idx
        - its possible that there is only 1 occurence of this element
        - so save this idx everytime
    - keep traversing next set of elements until element != target
    - return the start and end indicies
    time = o(n)
    space = o(1)
*/
// func searchRange(nums []int, target int) []int {
//     start, end := -1, -1
//     for i := 0; i < len(nums); i++ {
//         num := nums[i]
//         if num == target {
//             if start == -1 {
//                 start = i
//             }
//             end = i
//         }
//     }
//     return []int{start, end}
// }