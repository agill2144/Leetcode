type entry struct {
    station string
    time int
}
type UndergroundSystem struct {
    userEntries map[int]*entry
    fromToAvgTimes map[string][3]float64 // {runningTotal, numberOfEntries, avg}
}


func Constructor() UndergroundSystem {
    return UndergroundSystem{
        userEntries: map[int]*entry{},
        fromToAvgTimes: map[string][3]float64{},
    }
}


func (this *UndergroundSystem) CheckIn(id int, stationName string, t int)  {
    e := &entry{station: stationName, time: t}
    this.userEntries[id] = e
}


func (this *UndergroundSystem) CheckOut(id int, stationName string, t int)  {
    checkIn := this.userEntries[id]
    time := float64(t-checkIn.time)
    key := fmt.Sprintf("%v-%v", checkIn.station, stationName)
    
    newAvgTime := [3]float64{time, 1, time}
    existingAvgTime, ok := this.fromToAvgTimes[key]
    if !ok {
        this.fromToAvgTimes[key] = newAvgTime
        return
    }
    
    newAvgTime[0] = existingAvgTime[0] + time
    newAvgTime[1] = existingAvgTime[1] + 1
    newAvgTime[2] = newAvgTime[0] / newAvgTime[1]
    this.fromToAvgTimes[key] = newAvgTime
    delete(this.userEntries, id)
}


func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
    return this.fromToAvgTimes[fmt.Sprintf("%v-%v", startStation, endStation)][2]
}


/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */