package TimeSeriesAccessCounter

import "sync"

type TimeSeriesAccessCounter struct {
}

var instance *TimeSeriesAccessCounter
var once sync.Once

func GetInstance() *TimeSeriesAccessCounter {
	once.Do(func() {
		instance = &TimeSeriesAccessCounter{}
	})

	return instance
}

func (t *TimeSeriesAccessCounter) Count(ip string) int {
	return 1
}