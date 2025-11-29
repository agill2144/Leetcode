/*
    approach: sort then two ptrs
    tc = o(nlogn) + o(mlogn) + o(min(n,m))
    sc = o(1) if not counting sorting space otherwise its o(m) + o(n)
*/
func intersection(nums1 []int, nums2 []int) []int {
    sort.Ints(nums1)
    sort.Ints(nums2)
    n1 := 0
    n2 := 0
    out := []int{}
    for n1 < len(nums1) && n2 < len(nums2) {
        if nums1[n1] == nums2[n2] {
            if len(out) == 0 || nums1[n1] != out[len(out)-1] {
                out = append(out, nums1[n1])
            }
            n1++
            n2++
        } else if nums1[n1] < nums2[n2] {
            n1++
        }  else {
            n2++
        }
    }
    return out
}

/*
    approach: freq map
    - put smaller nums into a freq map
    - then iterate over other nums , check if exists in freq and if it does save it , decrement its counter
    m = len(nums1)
    n = len(nums2)
    tc = o(m)+o(n)
    sc = o(n)

*/
// func intersection(nums1 []int, nums2 []int) []int {
//     freq := map[int]int{}    
//     for i := 0; i < len(nums1); i++ {
//         freq[nums1[i]]++
//     }
//     out := []int{}
//     for i := 0; i < len(nums2); i++ {
//         if freq[nums2[i]] > 0 {
//             // DUPLICATES ARE NOT ALLOWED, hence updating counter to 0 once used
//             freq[nums2[i]] = 0
//             out = append(out, nums2[i])
//         }
//     }
//     return out
// }