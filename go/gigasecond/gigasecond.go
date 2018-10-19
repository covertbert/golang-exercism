package gigasecond

import "time"

// AddGigasecond adds a gigasecond (1e9 seconds) to the current time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1e9)
}
