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
*/

/*
    approach: top sort using DFS
    - dfs with cycle detection can be used for collecting the top sort
    - if dfs detects a cycle, we cannot finish all courses
        - this means there was a course that depends on a course we are currently taking
    v = num of courses
    e = total number of edges
    time = o(v+e)
    space = o(v+e) ; the adjList/graph map
*/
func canFinish(numCourses int, prerequisites [][]int) bool {
    adjList := map[int][]int{}
    for i := 0; i < len(prerequisites); i++ {
        u, v := prerequisites[i][1], prerequisites[i][0]
        adjList[u] = append(adjList[u], v)
    }

    visited := make([]bool, numCourses)
    var dfs func(node int, path []bool) bool
    dfs = func(node int, path []bool) bool {
        // base
        if path[node] {return false}
        if visited[node] {return true} 

        // logic
        visited[node] = true
        path[node] = true
        for _, nei := range adjList[node] {
            if !dfs(nei, path) {return false}
        }
        path[node] = false
        return true
    }
    p := make([]bool, numCourses)
    for i := numCourses-1; i >= 0; i-- {
        if !visited[i] {
            if !dfs(i, p) {return false}
        }
    }
    return true
}


/*
    approach: top sort using BFS
    v = num of courses
    e = total number of edges
    time = o(v+e)
    space = o(v+e) ; the adjList/graph map
*/
// func canFinish(numCourses int, prerequisites [][]int) bool {
//     adjList := map[int][]int{}
//     // tracks number of dependencies per course
//     indegrees := make([]int, numCourses)
//     for i := 0; i < len(prerequisites); i++ {
//         // you must take course b first if you want to take course a.
//         // a depends on b
//         // b is independent
//         // a -> b
//         // parent -> child
//         // {independent : [dependents] }
//         a := prerequisites[i][0]
//         b := prerequisites[i][1]
//         adjList[b] = append(adjList[b], a)
//         indegrees[a]++
//     }

//     // take all the courses that have no dependencies
//     q := []int{}
//     takenCount := 0
//     for i := 0; i < len(indegrees); i++ {
//         if indegrees[i] == 0 {
//             q = append(q, i)
//             takenCount++
//         }
//     }
//     if len(q) == 0 {return false}
//     for len(q) != 0 {
//         // take this course
//         dq := q[0]
//         q = q[1:]
//         // once this course is done, we can reduce its dependency on all courses that dependent on it
//         for _, nei := range adjList[dq] {
//             indegrees[nei]--
//             if indegrees[nei] == 0 {
//                 // this $nei course has no more dependency, we can take it
//                 q = append(q, nei)
//                 takenCount++
//             }
//         }
//     }
//     return takenCount == numCourses
// }