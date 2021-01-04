package TimeSeriesAccessCounter

import (
	"sync"
	"time"
)

type TimeSeriesAccessCounter struct {
	Data []map[string]interface{}
}

var instance *TimeSeriesAccessCounter
var once sync.Once

func GetInstance() *TimeSeriesAccessCounter {
	once.Do(func() {
		instance = &TimeSeriesAccessCounter{}
	})

	return instance
}

func (t *TimeSeriesAccessCounter) Count(ip string, seconds int) int {
	count := 0

	for _, data := range t.Data {
		if data["Ip"] == ip {
			count++
		}
	}

	return count
}

func (t *TimeSeriesAccessCounter) Insert(ip string) {
	t.Data = append(t.Data, map[string]interface{}{"Ip": ip, "Timestamp":time.Now().Unix()})
}