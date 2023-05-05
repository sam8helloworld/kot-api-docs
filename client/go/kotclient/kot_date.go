package kotclient

import (
	"strings"
	"time"
)

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
