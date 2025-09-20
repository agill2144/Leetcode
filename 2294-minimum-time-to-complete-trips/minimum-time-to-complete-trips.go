func minimumTime(time []int, totalTrips int) int64 {
    left := 1
    right := slices.Max(time)*totalTrips
    var ans int64
    for left <= right {
        mid := left + (right-left)/2
        trips := 0
        for i := 0; i < len(time); i++ {
            trips += int(math.Floor(float64(mid)/float64(time[i])))
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