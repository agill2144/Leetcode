func passThePillow(n int, time int) int {
    if time <= n {
        if time == n {return time-1}
        return time+1
    }

    currTime := 0
    currPerson := 1
    right := true
    for currTime < time {
        currTime++
        if right {
            currPerson++
            if currPerson == n {right = false}
        } else {
            currPerson--
            if currPerson == 1 {right = true}
        }
    }
    return currPerson
}

/*
    r = true
    1 2 3 4
    | | | |
        *

    t=5

*/