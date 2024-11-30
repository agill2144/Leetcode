/*
    order;
    1. classic sort
        tc = o(nlogn)
        sc = o(n)
    2. heap
        tc = (nlogk)
        sc = o(n) + o(k)
    3. bucket sort ( sort by freq ; use idx as freq )
        tc = o(n)
        sc = o(n)
    4. quick select ( sort by freq, sort in asc order by freq until ns is at targetIdx )
        tc = o(n)
        sc = o(n)
*/
func topKFrequent(arr []int, k int) []int {
    freq := map[int]int{}
    for i := 0; i < len(arr); i++ {freq[arr[i]]++}
    nums := []int{}
    for k, _ := range freq {nums = append(nums, k)}
    n := len(nums)
    targetIdx := n-k
    left := 0
    right := n-1
    for left <= right {
        pivot := right
        nsf := left
        for i := left; i < pivot; i++ {
            pivotFreq := freq[nums[pivot]]
            iFreq := freq[nums[i]]
            if iFreq <= pivotFreq {
                nums[i], nums[nsf] = nums[nsf], nums[i]
                nsf++
            }
        }
        nums[pivot], nums[nsf] = nums[nsf], nums[pivot]
        if nsf == targetIdx {return nums[targetIdx:]}
        if targetIdx > nsf {
            left = nsf+1
        } else {
            right = nsf-1
        }
    }
    return nil
}
// func topKFrequent(nums []int, k int) []int {
//     freq := map[int]int{}
//     for i := 0; i < len(nums); i++ {freq[nums[i]]++}

//     bucket := make([][]int, len(nums)+1)
//     for num, count := range freq {
//         if bucket[count] == nil {bucket[count] = []int{}}
//         bucket[count] = append(bucket[count], num)
//     }
//     out := []int{}
//     for i := len(bucket)-1; i >= 0 && len(out) != k; i-- {
//         if bucket[i] == nil {continue}
//         out = append(out, bucket[i]...)
//     }
//     return out
// }
