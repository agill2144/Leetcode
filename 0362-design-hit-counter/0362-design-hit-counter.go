/*
    approach: binary search
    - because the hits timestamps are guranteed to be in sorted order
    - we can save the hit timestamps in an array as they come in
    - and naturally they will be in a sorted order
    - when we have to find total number of hits from a current time - 5 mins
    - we will use binary search
        - since we have to go back 5 mins, whats the current time go back from?
        - thats the timestamp passed in as arg, this is the current time
        - now go back 5 mins ( timestamp-300s )
        - make the above subtraction our target for binary search, however!
        - we have to get number of events that happened after the target ( not including the target )
        - essentially search for the next bigger number thats the smallest after target
        - if not found, that means 0 events
        - if found, than to count number of events would be right most idx - target idx + 1
    
    time: o(logn)
    space: o(1) - not counting items array as extra space.
*/
type HitCounter struct {
    items []int
}


func Constructor() HitCounter {
    return HitCounter{items: []int{}}
}


func (this *HitCounter) Hit(timestamp int)  {
    this.items = append(this.items, timestamp) 
}


func (this *HitCounter) GetHits(timestamp int) int {
    target := timestamp-300
    idx := -1
    left := 0
    right := len(this.items)-1
    for left <= right {
        mid := left + (right-left) / 2
        if this.items[mid] > target {
            idx = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    if idx == -1 {return 0}
    return (len(this.items)-1) - idx + 1
}


/**
 * Your HitCounter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Hit(timestamp);
 * param_2 := obj.GetHits(timestamp);
 */