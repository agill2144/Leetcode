// Make the numbers searchable by tossing them in a set
// this also de-dupes elements
// therefore instead of looping thru input nums, now we can loop thru the deduped set
// and use the set to find a start number and then keep
// forming a chain until next num exists
// while keeping track of current chain
func longestConsecutive(nums []int) int {
	set := map[int]bool{}
    for i := 0; i < len(nums); i++ {
        set[nums[i]] = true
    }
    maxWin := 0
    for k, _ := range set {
        if set[k-1] {continue}
        start := k
        count := 0
        for set[start] {
            count++
            maxWin = max(maxWin, count)
            start++
        }
    }
    return maxWin
}