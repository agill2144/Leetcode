/*
    approach: stack
    - we can simulate this process by using a stack
    - once a function starts
    - we push that func into a stack ( id, and its startTime )
    - when a func ends, the top of the stack
        will always be the function that started,
        therefore the end event is for the 
        function sitting at the top of the stack
    - we simply need to make sure that 
    - top of the stack represents true start time of a function

    - what happens if a function is already running
    - and we get another start log
    - at this very moment, we have the 2 data points
    - to calc the top function running time so far
    - before pushing the new event into the stack
    - therefore when a function already exists in the stack
    - and we get another start event instead of end event
    - we have the new event start time
    - and we also have the existing function start time
    - we can take the difference and add to the res of the top func id
    - i.e time - topTime
    - WE DO NOT POP THE TOP OF THE STACK
    - because this function is paused, a new one is starting
    - and only 1 invocation can run at a time because 1 cpu core

    - when we run into an end event
    - we know the top of stack is the function thats ending
    - so we have the ending time ( current ith log event )
    - and we have the start time ( top of stack )
    - to calculate its running time = time - topTime + 1
    - we can add this to the top func idx in our output array
    - id == idx in the output array
    - its possible a function is starts / stops multiple times
    - therefore summing all the execution times in output array
    - REMEMBER top of the stack represents current active function
    - active function thats running 
    - so when an end event is processed, top of stack is popped
    - the next top will resume its execution ( it was paused )
    - but now its start time if off
    - the function that had started some time in past, got paused for x time
    - and now its resuming
    - so what will be the next top's corrected start time?
    - current event's end time + 1
    - if time now right now is this end event time
    - and processing the end event resumes the prev paused func
    - its obvious that the func that is resumed, its new start time will be curr time + 1
    - hence if there is an item in the top of stack, after processing end event
    - update its start time to current time + 1
    - what if there are multiple such paused functions?
    - each end event will update the newely resumed function start time
    - so its child function will get its start time updated
         when next end event gets processed


    tc = o(n)
    sc = o(n)

*/
func exclusiveTime(n int, logs []string) []int {
    st := [][]int{} // [id, startTime]
    out := make([]int, n)
    for i := 0; i < len(logs); i++ {
        log := strings.Split(logs[i], ":")
        id, _ := strconv.Atoi(log[0])
        logType := log[1]
        time, _ := strconv.Atoi(log[2])
        if logType == "start" {
            if len(st) > 0 {
                top := st[len(st)-1]
                out[top[0]] += (time-top[1])
            }
            st = append(st, []int{id, time})
        } else {
            top := st[len(st)-1]
            st = st[:len(st)-1]
            out[top[0]] += (time-top[1]+1)
            if len(st) > 0 {
                st[len(st)-1][1] = time+1
            }
        }
    }
    return out
}