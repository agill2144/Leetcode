// n = len(time)
// m = (minTime*totalTrips)-minTime+1
// tc = o(logm*n)
// sc = o(1)
func minimumTime(time []int, totalTrips int) int64 {
    left := math.MaxInt64
    for i := 0; i < len(time); i++ {
        left = min(left, time[i])
    }
    right := left*totalTrips
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