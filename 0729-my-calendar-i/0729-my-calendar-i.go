type MyCalendar struct {
    cal [][]int
}


func Constructor() MyCalendar {
    return MyCalendar{cal: [][]int{}}
}


func (this *MyCalendar) Book(start int, end int) bool {
    if len(this.cal) == 0 {
        this.cal = append(this.cal, []int{start, end})
        return true
    }
    prevEvent := this.searchLeft(start)
    if prevEvent == -1 {prevEvent = 0}
    prev := this.cal[prevEvent]
    if start < prev[1] && end > prev[0] {return false}

    nextEvent := this.searchRight(start)
    if nextEvent == -1 {nextEvent = 0}
    next := this.cal[nextEvent]
    if start < next[1] && end > next[0] {return false}

    this.cal = append(this.cal, []int{start, end})
    sort.SliceStable(this.cal, func(i, j int) bool{
        return this.cal[i][0] < this.cal[j][0]
    })
    return true 
}

func (this *MyCalendar) searchLeft(start int) int {
    left := 0
    right := len(this.cal)-1
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        if this.cal[mid][0] <= start {
            if this.cal[mid][0] == start {
                return mid
            }
            ans = mid
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return ans
}

func (this *MyCalendar) searchRight(start int) int {
    left := 0
    right := len(this.cal)-1
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        if this.cal[mid][0] >= start {
            if this.cal[mid][0] == start {
                return mid
            }
            ans = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return ans
}


/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */