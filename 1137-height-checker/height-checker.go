func heightChecker(heights []int) int {
    expected := make([]int, len(heights))
    copy(expected, heights)
    sort.Ints(expected)
    count := 0
    for i := 0; i < len(heights); i++ {
        if heights[i] != expected[i] {count++}
    }
    return count
}