type item struct{
    time int
    value string
}
type TimeMap struct {
    m map[string][]*item
}


func Constructor() TimeMap {
    return TimeMap{m:map[string][]*item{}}
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
    currItems, ok := this.m[key]
    if !ok {
        currItems = []*item{}
    }
    currItems = append(currItems, &item{timestamp,value})
    this.m[key] = currItems
}


func (this *TimeMap) Get(key string, timestamp int) string {
    items, ok := this.m[key]
    if !ok {return ""}
    
    left := 0
    right := len(items)-1
    ans := ""
    for left <= right {
        mid := left + (right-left)/2
        midTime := items[mid].time
        if midTime <= timestamp {
            if midTime == timestamp {return items[mid].value}
            ans = items[mid].value
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return ans
}


/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */