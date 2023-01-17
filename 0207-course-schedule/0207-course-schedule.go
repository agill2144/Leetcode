// approach: dependency resolution , connected components, top sort using BFS
// time: o(v+e)
// e is all the edges!
// v is all the courses (0 to numCourses-1)
// space: o(v+e) we built graph(adjList)  + o(v) for indgrees ( but o(v+e) is dominating term )
func canFinish(numCourses int, prerequisites [][]int) bool {
    // graph
    adjList := map[int][]int{}

    // count of incoming edges at a specific node
    // in this example, count of dependencies for a specific course ( where each course is the ith idx )
    indegrees := make([]int, numCourses)
    
    for i := 0; i < len(prerequisites); i++ {
        course := prerequisites[i][0]
        dependsOn := prerequisites[i][1]
        // dependsOn is the independent node
        // dependsOn is the parent node
        // dependsOn will have an outgoing edge
        // course is the dependent node
        // it will have an incoming edge ( ie: indgree +1 )
        indegrees[course]++
        // {parent: [childs]}
        // {independent: [dependent nodes]}
        adjList[dependsOn] = append(adjList[dependsOn], course)
    }
    
    q := []int{}
    for i := 0; i < numCourses; i++ {
        if indegrees[i] == 0 {
            q = append(q, i)
        }
    }
    if len(q) == 0 {return false}
    coursesTaken := 0
    for len(q) != 0 {
        // course taken successfully
        dq := q[0]
        q = q[1:]
        coursesTaken++

        // if we have taken $dq course successfully
        // it may have fulfilled other courses dependencies
        // so we need to search for courses that depend on $dq course
        // which is what our adjList has {independent: [dependent courses]}        
        childs := adjList[dq]
        for _, child := range childs {
            // reduce each course dependency by 1 since $dq course was taken
            indegrees[child]--
            // now if this dependent course is no longer dependent on anything, we can take it
            if indegrees[child] == 0 {
                q = append(q, child)
            }
        }
    }
    return numCourses == coursesTaken
}
