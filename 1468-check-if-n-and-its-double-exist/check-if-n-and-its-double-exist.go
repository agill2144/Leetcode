func checkIfExist(arr []int) bool {
    sort.Ints(arr)
    for i := 0; i < len(arr); i++ {
        ok, idx := bs(arr, arr[i]*2)
        if ok && idx != i {return true}
    }
    return false
}

func bs(nums []int, targetJ int) (bool, int) {
    left := 0
    right := len(nums)-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == targetJ {return true, mid}
        if targetJ > nums[mid] {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return false, -1
}

// hash hash - brain turning into a potato
// func checkIfExist(arr []int) bool {
//     set := map[int]bool{}
//     for i := 0; i < len(arr); i++ {
//         if arr[i] % 2 == 0 && set[arr[i]/2] {return true}
//         if set[arr[i]*2] {return true}
//         set[arr[i]] = true
//     }
//     return false
// }