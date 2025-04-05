package util

import (
	"fmt"
	"strings"
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
	ssMs := mmSs % time.Minute //nolint:staticcheck
	mm := mmSs - ssMs
	ms := ssMs % time.Second
	ss := ssMs - ms
	cn := fmt.Sprintf("%d:%02d:%02d", hh/time.Hour, mm/time.Minute, ss/time.Second)

	if ms > 0 {
		cn = fmt.Sprintf("%s.%09d", cn, ms)
		cn = strings.TrimRight(cn, "0")
	}

	if minus {
		cn = "-" + cn
	}

	return cn
}
