package hafas

import "time"

func (c *HafasClient) parseTime(timestamp string, day string) time.Time {
	layout := "20060102150405"

	var ts time.Time

	if timestamp == "" {
		return ts
	}

	dayOffset := 0

	if len(timestamp) == 8 {
		dayOffset = 1
		timestamp = timestamp[2:]
	}

	loc, _ := time.LoadLocation("Europe/Berlin")

	ts, _ = time.ParseInLocation(layout, day+timestamp, loc)

	ts = ts.AddDate(0, 0, dayOffset)

	return ts
}
