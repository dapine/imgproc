package math

import (
	"github.com/h2non/bimg"
)

func abs(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}

func NearestAngle(angle int64) bimg.Angle {
	angles := []int64{0, 45, 90, 135, 180, 235, 270, 315}

	minDiff := abs(angles[0] - angle)
	targetAngleIndex := 0

	for i, a := range angles {
		diff := abs(a - angle)
		if diff < minDiff {
			minDiff = diff
			targetAngleIndex = i
		}
	}

	return bimg.Angle(angles[targetAngleIndex])
}
