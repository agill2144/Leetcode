/*
    pattern identification:
    - a course depends on other course(s)
    - draw out the pre-reqs on a board and you will see something end up drawing a graph
    - is it a directed graph or undirected ?
    - prereqs = [[0,1]]; indicates that to take course 0 you have to first take course 1.
    - so course-0 depends on course-1
    - which also means you have to take course-1 first and then we can take course-0
    - 1->0
    - therefore a directed graph ( there may be cycles )

    approach: top sort using BFS

*/


func canFinish(numCourses int, prerequisites [][]int) bool {
    adjList := map[int][]int{}
    // tracks number of dependencies per course
    indegrees := make([]int, numCourses)
    for i := 0; i < len(prerequisites); i++ {
        // you must take course b first if you want to take course a.
        // a depends on b
        // b is independent
        // a -> b
        a := prerequisites[i][0]
        b := prerequisites[i][1]
        adjList[b] = append(adjList[b], a)
        indegrees[a]++
    }

    // take all the courses that have no dependencies
    q := []int{}
    takenCount := 0
    visited := make([]bool, numCourses)
    for i := 0; i < len(indegrees); i++ {
        if indegrees[i] == 0 {
            q = append(q, i)
            takenCount++
            visited[i] = true
        }
    }
    if len(q) == 0 {return false}
    for len(q) != 0 {
        // take this course
        dq := q[0]
        q = q[1:]
        // once this course is done, we can reduce its dependency on all courses that dependent on it
        for _, nei := range adjList[dq] {
            indegrees[nei]--
            if indegrees[nei] == 0 && !visited[nei] {
                // this $nei course has no more dependency, we can take it
                q = append(q, nei)
                takenCount++
                visited[nei] = true
            }
        }
    }
    return takenCount == numCourses
}