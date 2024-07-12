func maxDistance(position []int, m int) int {
    // min ans is max
    // atMin = mid
    // and look for max such mid value

    sort.Ints(position)
    left := 1
    right := position[len(position)-1]  - position[0]
    ans := 0
    for left <= right {
        mid := left + (right-left)/2
        

        count := 1
        lastPlaced := position[0]
        for i := 0; i < len(position); i++ {
            diff := position[i]-lastPlaced
            if diff >= mid {
                lastPlaced = position[i]
                count++
            }
        }

        if count < m {
            right = mid-1
        } else {
            ans = mid
            left = mid+1
        }
    }
    return ans
}