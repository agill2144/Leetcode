type MyCalendar struct {
    cal [][]int
}


func Constructor() MyCalendar {
    return MyCalendar{cal: [][]int{}}
}


func (this *MyCalendar) Book(start int, end int) bool {
    for i := 0; i < len(this.cal); i++ {
        s := this.cal[i][0]
        e := this.cal[i][1]
        if s < end && start < e {return false}
    }
    this.cal = append(this.cal, []int{start, end})
    return true
   
}

func (this *MyCalendar) nextSmallestStartTime(start int) int {
    ans := -1
    left := 0
    right := len(this.cal)-1
    
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