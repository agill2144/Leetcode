type LUPrefix struct {
    ptr int
    uploads []bool
}


func Constructor(n int) LUPrefix {
    return LUPrefix{
        ptr: 1,
        uploads: make([]bool, n+1),
    }
}


func (this *LUPrefix) Upload(video int)  {
    this.uploads[video] = true
    for this.ptr < len(this.uploads) && this.uploads[this.ptr] {
        this.ptr++
    }
}


func (this *LUPrefix) Longest() int {
    return this.ptr-1
}


/**
 * Your LUPrefix object will be instantiated and called as such:
 * obj := Constructor(n);
 * obj.Upload(video);
 * param_2 := obj.Longest();
 */