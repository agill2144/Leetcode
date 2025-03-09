func numberOfAlternatingGroups(colors []int, k int) int {
    total := 0
    n := len(colors)
    // tracks size of a window with alternating colors
    count := 1
    for i := 1; i < n+k-1; i++ {
        currIdx := i%n
        prevIdx := (i-1) % n
        if colors[currIdx] == colors[prevIdx] {
            // when color is same as previous
            // reset counter
            count = 1
        } else {
            // otherwise increase counter
            count++
        }

        // once we have k size window with alternative colors
        if count == k {
            // update our total number of windows found
            total++
            // one element is leaving our window, hence decrement the counter
            // (this is like left ptr moving forward, except we do not have a left ptr)
            count--
        }
    }
    return total
}