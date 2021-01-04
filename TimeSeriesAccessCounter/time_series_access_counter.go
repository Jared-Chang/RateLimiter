package TimeSeriesAccessCounter

import (
	"RateLimiter/UnixTime"
	"sync"
)

type TimeSeriesAccessCounter struct {
	Data        []map[string]interface{}
	UnixTime    UnixTime.UnixTime
	BufferRange int64
}

var instance *TimeSeriesAccessCounter
var once sync.Once

func GetInstance() *TimeSeriesAccessCounter {
	once.Do(func() {
		instance = &TimeSeriesAccessCounter{UnixTime: new(UnixTime.HumbleTime)}
	})

	return instance
}

func (t *TimeSeriesAccessCounter) Count(ip string, seconds int) int {

	afterTheTime := t.UnixTime.GetUnixNow() - int64(seconds)
	count := 0

	for _, data := range t.Data {
		if data["Ip"] == ip && data["Timestamp"].(int64) >= afterTheTime {
			count++
		}
	}

	return count
}

func (t *TimeSeriesAccessCounter) Insert(ip string) {

	now := t.UnixTime.GetUnixNow()

	t.Data = append(t.Data, map[string]interface{}{"Ip": ip, "Timestamp": now})

	outOfDateTime := now - t.BufferRange

	for {
		if t.Data[0]["Timestamp"].(int64) < outOfDateTime {
			t.Data = t.Data[1:]
		} else {
			return
		}
	}
}