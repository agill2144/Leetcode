func minimumTime(time []int, totalTrips int) int64 {
    left := 1
    right := slices.Max(time)*totalTrips
    var ans int64
    // T = max(time)*totalTrips
    // tc = o(logT * n)
    for left <= right {
        mid := left + (right-left)/2
        // how many trips can be completed if mid was atMax time allowed per bus
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