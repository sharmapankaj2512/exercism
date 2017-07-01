package gigasecond

// import path for the time package from the standard library
import "time"

const testVersion = 4

const gigaSeconds = time.Duration(1e9) * time.Second

func AddGigasecond(dateTime time.Time) time.Time {
	return dateTime.Add(gigaSeconds)
}
