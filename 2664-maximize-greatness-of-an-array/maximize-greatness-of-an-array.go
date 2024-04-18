func maximizeGreatness(nums []int) int {
    n := len(nums)
    sort.Ints(nums)

    slow, fast, count := 0, 1, 0
    for fast < n {
        if nums[fast] > nums[slow] {
            count++
            slow++
            fast++
        } else {
            fast++
        }
    }
    return count
}

/*
    [1,1,1,2,3,3,5]
           LR
    count = 3
*/

// time = o(n) + o(nlogn) + o(n) + o(n*n)
// func maximizeGreatness(nums []int) int {
//     n := len(nums)
//     dupe := make([]int, n)
//     copy(dupe, nums)
//     sort.Ints(dupe)
//     lastIdx := map[int]int{}
//     for i := 0; i < n; i++ {
//         lastIdx[dupe[i]] = i
//     }
//     used := make([]bool, n)
//     count := 0
//     for i := 0; i < n; i++ {
//         num := nums[i]
//         idx := lastIdx[num]+1
//         for idx < n && used[idx] {
//             idx++
//         }
//         if idx < n {used[idx] = true; count++}
//     }
//     return count
// }

/*

    
    [1,3,5,2,1,3,1]
     2 5 0 3 3 0 i
    [1,1,1,2,3,3,5]
                 i
    [F,F,F,T,T,T,T]


    count = 3
    [1,2,3,4]
     2 3 4
    [1,2,3,4]
             i        
    [F,T,T,T]
*/