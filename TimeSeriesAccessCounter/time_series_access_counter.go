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

func (t *TimeSeriesAccessCounter) Count(ip string, seconds int) int {
	return 1
}

func (t *TimeSeriesAccessCounter) Insert(ip string) {
	panic("implement me")
}