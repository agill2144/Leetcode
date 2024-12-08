func maxTwoEvents(events [][]int) int {
    n := len(events)
    sort.Slice(events, func(i, j int)bool {
        return events[i][0] < events[j][0]
    })
    maxValues := make([]int, n)
    maxValues[n-1] = events[n-1][2]
    for i := n-2; i >= 0; i-- {
        maxValues[i] = max(events[i][2], maxValues[i+1])
    }

    maxSum := 0
    for i := 0; i < n; i++ {
        end, val := events[i][1], events[i][2]
        maxSum = max(maxSum, val)
        idx := search(events, i+1, end+1)
        if idx != -1 {
            maxSum = max(maxSum, max(val, val+maxValues[idx]))
        }
    }
    return maxSum
}

func search(events [][]int, left int, targetStart int) int {
    right := len(events)-1
    idx := -1
    for left <= right {
        mid := left + (right-left)/2
        if events[mid][0] >= targetStart {
            idx = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return idx
}