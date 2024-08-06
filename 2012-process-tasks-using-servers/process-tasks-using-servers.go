func assignTasks(servers []int, tasks []int) []int {
    busy := &busyServers{servers: []server{}}
    idle := &idleServers{servers: []server{}}
    for i := 0; i < len(servers); i++ {
        srv := server{weight: servers[i], idx: i}
        heap.Push(idle, srv)
    }


    currTime := 0
    freeUpBusy := func(ctime int) {
        for len(busy.servers) > 0 && busy.servers[0].endTime <= ctime {
            heap.Push(idle, heap.Pop(busy))
        }                
    }

    out := make([]int, len(tasks))
    for i := 0; i < len(tasks); i++ {
        currTime = max(currTime, i)
        // reconcile busy ones, they might be free now
        freeUpBusy(currTime)
       
        // if we have no free servers
        // wait until currTime is at earliest endtime
        // i.e ; fastForward our currTime to earliest endTime
        if len(idle.servers) == 0 {
            currTime = busy.servers[0].endTime
            freeUpBusy(currTime)
        }


        taskTime := tasks[i]
        endTime := currTime+taskTime
        // now we have for sure have a free server
        srv := heap.Pop(idle).(server)
        srv.endTime = endTime
        heap.Push(busy, srv)
        out[i] = srv.idx 
    }
    return out
}


type server struct {
    endTime int
    weight int
    idx int
}

// minHeap, sorted by weight
// if weight is same, sorted by idx
type idleServers struct {
	servers []server
}

func (m *idleServers) Less(i, j int) bool {
	if m.servers[i].weight == m.servers[j].weight {
        return m.servers[i].idx < m.servers[j].idx
    }
    return m.servers[i].weight < m.servers[j].weight
}
func (m *idleServers) Swap(i, j int) {
	m.servers[i], m.servers[j] = m.servers[j], m.servers[i]
}
func (m *idleServers) Len() int {
	return len(m.servers)
}
func (m *idleServers) Push(x interface{}) {
	m.servers = append(m.servers, x.(server))
}
func (m *idleServers) Pop() interface{} {
	out := m.servers[len(m.servers)-1]
	m.servers = m.servers[:len(m.servers)-1]
	return out
}



// minHeap, sorted by endTime
// which server is available the earliest
type busyServers struct {
	servers []server
}

func (m *busyServers) Less(i, j int) bool {
	return m.servers[i].endTime < m.servers[j].endTime
}
func (m *busyServers) Swap(i, j int) {
	m.servers[i], m.servers[j] = m.servers[j], m.servers[i]
}
func (m *busyServers) Len() int {
	return len(m.servers)
}
func (m *busyServers) Push(x interface{}) {
	m.servers = append(m.servers, x.(server))
}
func (m *busyServers) Pop() interface{} {
	out := m.servers[len(m.servers)-1]
	m.servers = m.servers[:len(m.servers)-1]
	return out
}
