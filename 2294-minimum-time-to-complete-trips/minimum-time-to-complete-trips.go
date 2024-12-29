func minimumTime(time []int, totalTrips int) int64 {
    var ans int64 = -1
    left := 1
    right := slices.Max(time)*totalTrips
    for left <= right {
        mid := left + (right-left)/2
        trips := 0
        for i := 0; i < len(time); i++ {
            trips += (mid/time[i])
        }
        if trips >= totalTrips {
            ans = int64(mid)
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return ans
}