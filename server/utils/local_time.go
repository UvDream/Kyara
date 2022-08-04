package utils

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type LocalTime time.Time

var (
	TimeFormat = "2006-01-02 15:04:05"
	Zone       = "Asia/Shanghai"
)

func (t *LocalTime) MarshalJSON() ([]byte, error) {
	buf := make([]byte, 0, len(TimeFormat)+2)
	buf = append(buf, '"')
	buf = time.Time(*t).AppendFormat(buf, TimeFormat)
	buf = append(buf, '"')
	return buf, nil
}

func (t *LocalTime) UnmarshalJSON(data []byte) (err error) {
	var now time.Time
	now, err = time.ParseInLocation(`"`+TimeFormat+`"`, string(data), time.Local)
	if err != nil {
		*t = LocalTime(now)
	}
	return
}

func (t LocalTime) String() string {
	return time.Time(t).Format(TimeFormat)
}

func (t LocalTime) Local() time.Time {
	loc, _ := time.LoadLocation(Zone)
	return time.Time(t).In(loc)
}

func (t *LocalTime) Scan(value interface{}) error {
	if v, ok := value.(time.Time); ok {
		*t = LocalTime(v)
		return nil
	}

	return fmt.Errorf("can not convert %v to LocalTime", value)
}

func (t LocalTime) Value() (driver.Value, error) {
	value := time.Time(t)
	return value, nil
}
