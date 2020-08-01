package utils

import (
	"time"
)

func FormateDate(date string) (string, error) {
	layoutDb := "2006-01-02T15:04:05-07:00" //"	Mon Jan 2 15:04:05 -0700 MST 2006"
	layoutBlog := "02-Jan-2006 3:04 PM"
	if t, err := time.Parse(layoutDb, date); err != nil {
		return "", err
	} else {
		return t.Format(layoutBlog), nil
	}
}
