type Logger struct {
    m map[string]int
}


func Constructor() Logger {
    return Logger{m: map[string]int{}}
}


func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
    lastTime, ok := this.m[message]
    if !ok {
        this.m[message] = timestamp
        return true
    }
    if timestamp-lastTime >= 10 {
        this.m[message] = timestamp
        return true
    }
    return false
}


/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */