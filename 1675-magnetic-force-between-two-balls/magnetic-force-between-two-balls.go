func maxDistance(position []int, m int) int {
    // min ans is maximized ; "minimum magnetic force between any two balls is maximum"
    sort.Ints(position)
    left := 1
    right := position[len(position)-1]
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        
        // mid is our at-min val
        // how many balls can be placed if at-min the force is mid
        count := 1
        lastPlaced := 0
        for i := 1; i < len(position); i++ {
            currPos := position[i]
            lastPos := position[lastPlaced]
            if currPos - lastPos >= mid {
                count++
                lastPlaced = i
            }
        }
        
        if count >= m {
            ans = mid
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return ans
}