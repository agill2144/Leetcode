func minimumTime(time []int, totalTrips int) int64 {
    left := 1
    right := 0
    for i := 0; i < len(time); i++ {
        right = max(right, time[i])
    }
    right *= totalTrips
    var ans int64 = -1

    for left <= right {
        mid := left + (right-left)/2
        
        count := 0
        for i := 0; i < len(time); i++ {
            count += mid/time[i]
        }
        if count >= totalTrips {
            ans = int64(mid)
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return ans
}