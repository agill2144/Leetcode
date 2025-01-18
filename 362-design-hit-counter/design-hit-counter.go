type HitCounter struct {
    times []int
}


func Constructor() HitCounter {
    return HitCounter{times:[]int{}}    
}


func (this *HitCounter) Hit(timestamp int)  {
    this.times = append(this.times, timestamp)    
}


func (this *HitCounter) GetHits(timestamp int) int {
    target := timestamp-300
    if target < 0 {return len(this.times)}
    left := 0
    right := len(this.times)-1
    idx := len(this.times)
    for left <= right {
        mid := left + (right-left)/2
        if this.times[mid] > target {
            idx = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    res := len(this.times)-idx
    // fmt.Println("GetHits(",timestamp,"): ")
    // fmt.Println("target: ", target)
    // fmt.Println("found idx: ", idx)
    // fmt.Println(this.times)
    // fmt.Println("-----------------------------------")
    return res
}


/**
 * Your HitCounter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Hit(timestamp);
 * param_2 := obj.GetHits(timestamp);
 */