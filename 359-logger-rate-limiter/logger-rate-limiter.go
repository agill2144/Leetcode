type Logger struct {
    lastPrinted map[string]int
}


func Constructor() Logger {
    return Logger{lastPrinted: map[string]int{}}
}


func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
    lastT, ok := this.lastPrinted[message]
    if ok {
        if timestamp >= lastT + 10 {
            this.lastPrinted[message] = timestamp
            return true
        } else {
            return false
        }
    }
    this.lastPrinted[message] = timestamp
    return true
}


/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */