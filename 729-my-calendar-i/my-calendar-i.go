type MyCalendar struct {
    times [][]int
}


func Constructor() MyCalendar {
    return MyCalendar{times: [][]int{}}
}


func (this *MyCalendar) Book(start int, end int) bool {
    if len(this.times) == 0 {
        this.times = append(this.times, []int{start, end})
        return true
    }
    for i := 0; i < len(this.times); i++ {
        s, e := this.times[i][0], this.times[i][1]
        // start = 33
        // s = 49
        if start < e && s < end {return false}
    }
    this.times = append(this.times, []int{start, end})
    return true
}


/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */