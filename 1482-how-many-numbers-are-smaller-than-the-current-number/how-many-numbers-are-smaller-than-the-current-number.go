func smallerNumbersThanCurrent(nums []int) []int {
    n := len(nums)
    pairs := [][]int{}
    for i := 0; i < n; i++ {
        pairs = append(pairs, []int{nums[i],i})
    }
    sort.Slice(pairs, func(i, j int)bool {
        return pairs[i][0] < pairs[j][0]
    })
    out := make([]int, n)
    for i := 1; i < n; i++ {
        if pairs[i][0] == pairs[i-1][0] {
            out[pairs[i][1]] = out[pairs[i-1][1]]
        } else {
            out[pairs[i][1]] = i
        }
    }
    return out
}