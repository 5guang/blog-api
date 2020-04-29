package util

import "time"

func Timestamp2str(timestamp int64) string  {
	if timestamp != 0 {
		tm := time.Unix(timestamp, 0)
		return tm.Format("2006-01-02 15:04:05")
	}
	return ""
}