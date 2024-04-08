func countStudents(students []int, sandwiches []int) int {
    // count to keep track of how many we could not feed
    count := 0
    for count != len(sandwiches) {
        pref := students[0]
        students = students[1:]
        if pref == sandwiches[0] {
            sandwiches = sandwiches[1:]
            count = 0
        } else {
            students = append(students, pref)
            count++
        }
    }
    return count
}