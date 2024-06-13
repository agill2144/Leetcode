func minMovesToSeat(seats []int, students []int) int {
    sort.Ints(seats)
    sort.Ints(students)
    total := 0
    for i := 0; i < len(students); i++ {
        total += abs(students[i]-seats[i])
    }
    return total
}

func abs(x int) int {
    if x < 0 {return x*-1}
    return x
}