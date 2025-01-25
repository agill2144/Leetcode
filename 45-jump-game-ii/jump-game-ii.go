func jump(nums []int) int {
    n := len(nums)
    i := 0
    right := 0
    count := 0
    farthestRight := 0
    for i < n && right < n-1 {
        farthestRight = max(farthestRight, i+nums[i])
        if i == right {
            count++
            right = farthestRight
        }
        i++
    }
    return count
}
// brute force: 
// tc = x^n
// sc = o(n)
// func jump(nums []int) int {
//     n := len(nums)
//     minJumps := math.MaxInt64
//     var dfs func(start, count int)
//     dfs  = func(start, count int) {
//         if start >= n-1 {
//             if start == n-1 {
//                 minJumps = min(minJumps, count)
//             }
//             return
//         }

//         for i := nums[start]; i >= 1; i-- {
//             dfs(start+i, count+1)
//         }
//     }
//     dfs(0, 0)
//     return minJumps
// }