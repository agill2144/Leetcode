type HitCounter struct {
    times []int
}


func Constructor() HitCounter {
    return HitCounter{times: []int{}}    
}


func (this *HitCounter) Hit(t int)  {
    this.times = append(this.times, t)
}


func (this *HitCounter) GetHits(t int) int {
    diff := t-300
    if diff < 0 {return len(this.times)}
    left := 0
    right := len(this.times)-1
    idx := len(this.times)
    for left <= right {
        mid := left + (right-left)/2
        if this.times[mid] > diff {
            idx = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return len(this.times)-idx
}


/**
 * Your HitCounter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Hit(timestamp);
 * param_2 := obj.GetHits(timestamp);
 */