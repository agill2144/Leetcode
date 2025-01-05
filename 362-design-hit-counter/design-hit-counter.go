type HitCounter struct {
    leftPtr int
    hits []int
}


func Constructor() HitCounter {
    return HitCounter{hits: []int{}}
}


func (this *HitCounter) Hit(timestamp int)  {
    this.hits = append(this.hits, timestamp)
    until := timestamp-300
    if until <= 0 {return}
    for this.leftPtr < len(this.hits) &&
        this.hits[this.leftPtr] < until {
            this.leftPtr++
    }
}


func (this *HitCounter) GetHits(timestamp int) int {
    diff := timestamp-300
    if diff <= 0 {return len(this.hits)}
    left := this.leftPtr
    right := len(this.hits)-1
    res := -1
    for left <= right {
        mid := left + (right-left)/2
        if this.hits[mid] > diff {
            res = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    if res == -1 {return 0}
    return len(this.hits)-res
}


/**
 * Your HitCounter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Hit(timestamp);
 * param_2 := obj.GetHits(timestamp);
 */