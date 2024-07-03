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

	ts, _ = time.Parse(layout, day+timestamp)

	ts = ts.AddDate(0, 0, dayOffset)

	return ts
}
