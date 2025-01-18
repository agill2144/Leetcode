func minimumTime(time []int, totalTrips int) int64 {
    left := slices.Min(time)
    right := left * totalTrips
    var res int64
    for left <= right {
        mid := left + (right-left)/2
        trips := 0
        for i := 0; i < len(time); i++ {
            trips += (mid/time[i])
        }
        if trips >= totalTrips {
            res = int64(mid)
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return res
}