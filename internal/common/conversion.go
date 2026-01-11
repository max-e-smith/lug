package common

import "math"

const GB = 1000 * 1000 * 1000

func ByteToGB(bytes int64) float64 {
	gb := float64(bytes) / GB
	return math.Trunc(gb*100) / 100
}
