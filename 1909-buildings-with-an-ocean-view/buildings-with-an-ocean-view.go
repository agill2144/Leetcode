// keep track of tallest building from right side
// a building on left side of tallest seen so far can see ocean if
// its height > tallest seen so far
// time = o(n)
// space = o(1)
func findBuildings(heights []int) []int {
    n := len(heights)
    tallestSoFar := heights[n-1]
    out := []int{n-1}
    for i := n-2; i >= 0; i-- {
        if heights[i] > tallestSoFar {
            out = append(out, i)
        }
        tallestSoFar = max(tallestSoFar, heights[i])
    }
    for i := 0; i < len(out)/2; i++ {
        out[i], out[len(out)-1-i] = out[len(out)-1-i], out[i]
    }
    return out
}

// next greater on right ( >= on right )
// if an ith building has a ngr, then this building cannot see the ocean
// if an ith building does not have a ngr, it can see the ocean
// time = o(2n) = o(n)
// space = o(n) for st
// func findBuildings(heights []int) []int {
//     nextGreater := nextGreaterOnRight(heights)
//     out := []int{}
//     for i := 0; i < len(nextGreater); i++ {
//         if nextGreater[i] == -1 {out = append(out, i)}
//     }
//     return out
// }

// func nextGreaterOnRight(nums []int) []int {
//     // process top
//     st := []int{} //idx
//     n := len(nums)
//     out := make([]int, n)
//     for i := 0; i < n; i++ {
//         out[i] = -1
//     }
//     for i := 0; i < n; i++ {
//         for len(st) != 0 && nums[i] >= nums[st[len(st)-1]] {
//             top := st[len(st)-1]
//             st = st[:len(st)-1]
//             out[top] = i
//         }
//         st = append(st, i)
//     }
//     return out
// }