package main

import (
	"fmt"
	"strconv"
	"strings"
)

type framerates string

const (
	ndf25 string = "NDF25"
	df30  string = "DF30"
)

func validateFramerate(framerate string) (ok bool) {
	switch strings.ToUpper(framerate) {
	case ndf25, df30:
		ok = true
	}
	return ok
}

func isDropFrame(framerate string) (df bool) {
	switch strings.ToUpper(framerate) {
	case ndf25:
		df = false
	case df30:
		df = true
	}
	return df
}

func getRate(framerate string) (rate uint64) {
	switch strings.ToUpper(framerate) {
	case ndf25:
		rate = 25
	case df30:
		rate = 30
	}
	return rate
}

func makeArrayFromTimecode(timecode string) (tcArr []uint64, err error) {
	// figure out logic
	// throw error if can't make valid array
	tcArrStr := strings.Split(timecode, ":")
	if len(tcArrStr) != 4 { // use cap()?
		return tcArr, fmt.Errorf("Invalid timecode:", timecode)
	}
	for i, num := range tcArrStr {
		iNum, err := strconv.ParseInt(num, 10, 64)
		if err != nil {
			return tcArr, err
		}
		num := uint64(iNum)
		if i > 0 && i < 3 && num > 59 {
			return tcArr, fmt.Errorf("Invalid timecode:", timecode)
		}
	}

	// additional handling for dropframe
	return tcArr, nil
}

func validateTimecode(timecode string, framerate string) (ok bool, err error) {
	// figure out logic
	tcArr, tcErr := makeArrayFromTimecode(timecode)
	if tcErr != nil {
		return false, tcErr
	}
	// validate that the numbers make sense

	// additional dropframe logic

	return ok, nil
}

func convertTcToFr(timecode string, framerate string) (frames uint64) {
	if !isDropFrame(framerate) {
		rate := getRate(framerate)
		tcArr, _ := makeArrayFromTimecode(timecode)
	}
	return frames
}

// Duration represents a timecode label Label, a framerate Rate,
// and its calculated number of frames Frames
type Duration struct {
	Label  string
	Rate   string
	Frames uint64
}

// NewDurationFromString returns a new Duration object
func NewDurationFromString(timecode string, framerate string) (Duration, error) {
	// timecode and framerate validation
	ok, err := validateTimecode(timecode, framerate)
	if !ok {
		return Duration{}, err
	}

	return Duration{
		Label:  timecode,
		Rate:   framerate,
		Frames: convertTcToFr(timecode, framerate),
	}, nil
}

func main() {
	// Create duration object of 1:02:03:04 @ NDF25
	duration, err := NewDurationFromString("1:02:03:04", "NDF25")
	if err != nil {
		panic(err)
	}
	fmt.Println("Duration:", duration)
	// Create a timecode object at 10:31:00;02 @ DF30
	// Create duration from 10000 frames @ DF30
	// Create a timecode object from 90000 frames @ NDF25

	// Convert between NDF25 and DF30
	// Convert between number of frames and timecode label
	// Convert between duration and timecode label

	// Add a duration to a timecode
	// Subtract a timecode from a duration
	// Add a duration to a duration
	// Subtract a timecode from a timecode
}
