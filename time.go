package utils

import "time"

func MaxDuration(ds ...time.Duration) (d time.Duration) {
	d = 0

	for _, d1 := range ds {
		if d1 > d {
			d = d1
		}
	}

	return
}
