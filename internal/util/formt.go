package util

import (
	"fmt"
	"time"
)

func ColonNotation(d time.Duration) string {
	minus := false

	if d < 0 {
		minus = true
		d *= -1
	}

	mmSs := d % time.Hour
	hh := d - mmSs
	ss := mmSs % time.Minute
	mm := mmSs - ss
	cn := fmt.Sprintf("%d:%02d:%02d", hh/time.Hour, mm/time.Minute, ss/time.Second)

	if minus {
		cn = "-" + cn
	}

	return cn
}
