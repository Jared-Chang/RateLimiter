package UnixTime

import "time"

type UnixTime interface {
	GetUnixNow() int64
}

type HumbleTime struct {
}

func (h HumbleTime) GetUnixNow() int64 {
	return time.Now().Unix()
}
