type Logger struct {
    msgs map[string]int
}


func Constructor() Logger {
    return Logger{msgs: map[string]int{}}
}


func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
    t, ok := this.msgs[message]
    if ok {
        if timestamp >= t {
            this.msgs[message] = timestamp+10 
            return true
        }
        return false
    }
    this.msgs[message] = timestamp+10
    return true
}


/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */