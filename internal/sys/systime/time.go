package systime

import "time"

var nowFunc = stdNowFunc

// InitTime initialize the nowFunc
// Please use this method only in test
func InitTime(t time.Time) {
	nowFunc = func() time.Time {
		return t
	}
}

// InitNowFunc changes the now function
func InitNowFunc(f func() time.Time) {
	nowFunc = f
}

// Reset set the nowFunc to its default value
// Please use this method only in test
func Reset() {
	nowFunc = stdNowFunc
}

// Now makes time testable
func Now() time.Time {
	return nowFunc()
}

func stdNowFunc() time.Time {
	return time.Now().UTC().Truncate(time.Millisecond)
}
