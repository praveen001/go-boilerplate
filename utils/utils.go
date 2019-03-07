package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// TimecodeToFrames .
func TimecodeToFrames(tc string, fps float64) int {
	var msec int
	if strings.Contains(tc, ".") {
		msec = TimecodeMsecToMsec(tc, fps)
	} else {
		msec = TimecodeFramesToMsec(tc, fps)
	}

	return MsecToFrames(msec, fps)
}

// TimecodeFramesToMsec .
func TimecodeFramesToMsec(tc string, fps float64) int {
	parts := strings.Split(tc, ":")
	fmt.Println(parts)
	msec := 0

	hours, _ := strconv.Atoi(parts[0])
	msec += hours * 60 * 60 * 1000

	mins, _ := strconv.Atoi(parts[1])
	msec += mins * 60 * 1000

	secs, _ := strconv.Atoi(parts[2])
	msec += secs * 1000

	frameDuration := 1000 / fps
	frames, _ := strconv.Atoi(parts[3])
	msec += int(float64(frames) * frameDuration)

	return msec
}

// TimecodeMsecToMsec .
func TimecodeMsecToMsec(tc string, fps float64) int {
	return 0
}

// MsecToFrames .
func MsecToFrames(msec int, fps float64) int {
	return 0
}