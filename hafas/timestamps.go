package hafas

import "time"

func (c *HafasClient) parseTime(timestamp string, day string, startDate time.Time) time.Time {
	layout := "20060102150405"

	var ts time.Time

	if timestamp == "" {
		return ts
	}

	ts, _ = time.Parse(layout, day+timestamp)

	// if startDate.After(ts) {
	// 	// this is neccessary for overnight journeys.
	// 	ts = ts.AddDate(0, 0, 1)
	// }

	return ts
}
