func countStudents(students []int, sandwiches []int) int {
    st := []int{}
    for i := 0; i < len(sandwiches); i++ {st = append(st, sandwiches[i])}
    // count to keep track of how many we could not feed
    count := 0
    // when count == len of sandwiches, it means all students went around the circle
    // if there are 3 students and 3 sandwiches
    // all of them want sq but top is circle
    // we can safely exit once the very first student comes back from the back of the line
    // count keeps track of how many students we could not feed
    // if count becomes 3, it means we save 3 consec students which would not take the top sandwich
    // and if we only had 3 sandwiches in total and all 3 students did not take it, 
    // it means we ran a full cycle of the queue without anyone making dents in the stack
    for count != len(st) {
        pref := students[0]
        students = students[1:]
        if pref == st[0] {
            st = st[1:]
            count = 0
        } else {
            students = append(students, pref)
            count++
        }
    }
    return count
}