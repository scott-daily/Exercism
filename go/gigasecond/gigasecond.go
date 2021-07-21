// This package allows one to add a billion seconds (a gigasecond) to a given time.
package gigasecond

import "time"

func AddGigasecond(t time.Time) time.Time {

	updated_time := t.Add(time.Second * 1000000000)

	return updated_time
	
}
