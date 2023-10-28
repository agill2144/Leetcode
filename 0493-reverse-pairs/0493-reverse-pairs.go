/*
    approach: divide and conquer ( merge sort )
*/
func reversePairs(nums []int) int {
    count := 0
    var dfs func(left, right int)
    dfs = func(left, right int) {
        // base
        if left >=  right {return}
        
        // logic
        mid := left + (right-left)/2
        dfs(left, mid)
        dfs(mid+1, right)
        
        // left and right halves are sorted
        // check if left (ith) element > right (jth) element
        j := mid+1
        for i := left; i <= mid; i++ {
            for j <= right {
                // if ith val > jth val * 2 
                // then all the next ith val (from left side, which are sorted )
                // WILL also be greater, therefore count += rest of the elements from left side
                if nums[i] > nums[j]*2{
                    count += (mid-i+1)
                    j++
                // if left val is not > right val, break, because we need a bigger ith val
                // leave j where it is
                } else {
                    break
                }
            }
        }
        
        // merge sorted array
        aux := []int{}
        l := left
        r := mid+1
        for l <= mid && r <= right {
            if nums[l] < nums[r] {
                aux = append(aux, nums[l])
                l++
            } else {
                aux = append(aux, nums[r])
                r++
            }
        }
        for l <= mid {aux = append(aux, nums[l]); l++}
        for r <= right {aux = append(aux, nums[r]); r++}
        auxPtr := 0
        for i := left ; i <= right; i++ {
            nums[i] = aux[auxPtr]
            auxPtr++
        }
        
    }
    dfs(0, len(nums)-1)
    return count
}

// approach: brute force
// check all possibe pairs
// time = o(n^2)
// space = o(1)
// func reversePairs(nums []int) int {
//     count := 0
//     for i := 0; i < len(nums); i++ {
//         for j := i+1; j < len(nums); j++ {
//             if nums[i] > nums[j]*2 {
//                 count++
//             }
//         }
//     }
//     return count
// }