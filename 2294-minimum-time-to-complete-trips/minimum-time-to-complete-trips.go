func minimumTime(time []int, totalTrips int) int64 {
    left := 1
    right := slices.Max(time)*totalTrips
    var ans int64 = -1
    for left <= right {
        // atMax time allowed per bus
        mid := left + (right-left)/2
        trips := 0
        for i := 0; i < len(time); i++ {
            trips += (mid/time[i])
        }
        // when does mid not work?
        if trips < totalTrips {
            left = mid+1
        } else {
            ans = int64(mid)
            right = mid-1
        }
    }
    return ans
}