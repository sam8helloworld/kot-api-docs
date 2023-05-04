package openapi

import (
	"strings"
	"time"
)

// NOTE: https://developer.kingtime.jp/#header-%E6%97%A5%E6%99%82%E8%A1%A8%E7%8F%BE
// 打刻時刻の書式: 2016-05-10T10:00:00+09:00
// 打刻時刻以外の書式: 2016-05-10T10:00+09:00
const (
	customTimeFormat               = "2006-01-02T15:04:05-07:00"
	customTimeFormatWithoutSeconds = "2006-01-02T15:04-07:00"
)

type KotDate struct {
	time.Time
}

func (kd *KotDate) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), "\"")
	t, err := time.Parse(customTimeFormat, str)
	if err != nil {
		t, err = time.Parse(customTimeFormatWithoutSeconds, str)
		if err != nil {
			return err
		}
	}
	kd.Time = t
	return nil
}

func (kd KotDate) MarshalJSON() ([]byte, error) {
	return []byte(kd.Time.Format(customTimeFormat)), nil
}
