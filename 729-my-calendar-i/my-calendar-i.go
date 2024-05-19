type MyCalendar struct {
    cal [][]int
}


func Constructor() MyCalendar {
    return MyCalendar{cal: [][]int{}}
}


// everytime this func is invoked
// o(logn) + o(logn) + o(n)
// o(n) time
// space: o(1)
func (this *MyCalendar) Book(start int, end int) bool {
    if len(this.cal) == 0 {
        this.cal = append(this.cal, []int{start, end})
        return true
    }
    // o(logn)
    prevIdx := this.searchLeft(start)
    if prevIdx == -1 {prevIdx = 0}
    prev := this.cal[prevIdx]
    if start < prev[1] && end > prev[0] {return false}
    // if end > prev[0] && start < prev[1]  {return false}
    

    // o(logn)
    nextIdx := this.searchRight(start)
    if nextIdx == -1 {nextIdx = 0}
    next := this.cal[nextIdx]
    if start < next[1] && end > next[0] {return false}
    // if end > next[0] && start < next[1]  {return false}
    
    // o(n)
    this.insert(start, end)
    return true 
}

func (this *MyCalendar) insert(start, end int) {  
    curr := []int{start, end}
    this.cal = append(this.cal, curr)
    for i := len(this.cal)-1; i > 0; i-- {
        prev := this.cal[i-1]
        if curr[0] < prev[0] {
            this.cal[i], this.cal[i-1] = this.cal[i-1], this.cal[i]
        }
    }
}

func (this *MyCalendar) searchLeft(target int) int {
    left := 0
    right := len(this.cal)-1
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        if this.cal[mid][0] <= target {
            if this.cal[mid][0] == target {
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

func (this *MyCalendar) searchRight(target int) int {
    left := 0
    right := len(this.cal)-1
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        if this.cal[mid][0] >= target {
            if this.cal[mid][0] == target {
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