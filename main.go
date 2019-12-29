package main

import (
	"fmt"
)

type FrameCount uint64

type FrameRate string

const (
	NDF25 FrameRate = "NDF25"
	DF30  FrameRate = "DF30"
)

func (fr FrameRate) Name() (name string) {
	switch fr {
	case NDF25:
		name = "NDF25"
	case DF30:
		name = "DF30"
	}
	return name
}

func (fr FrameRate) Rate() (rate uint64) {
	switch fr {
	case NDF25:
		rate = 25
	case DF30:
		rate = 30
	}
	return rate
}

func (fr FrameRate) DropFrame() bool {
	if fr == DF30 {
		return true
	}
	return false
}

type Frame interface {
	Text() string
	NumFrames() uint64
	Rate() string
	ParseText(string, string) (Frame, error)
	ParseFrames(string, uint64) (Frame, error)
}

type FrameLabel struct {
	FrameCount FrameCount
	FrameRate  FrameRate
}

// Text returns frame as string (e.g., 01:00:00:00)
func (fl FrameLabel) Text() string {
	return ""
}

// NumFrames returns frame as number of frames (e.g., 90000)
func (fl FrameLabel) NumFrames() uint64 {
	return 0
}

// Rate returns the Frame's frame rate
func (fl FrameLabel) Rate() string {
	return ""
}

// ParseText creates a new FrameLabel from text
func (fl FrameLabel) ParseText(frameRate string, timecode string) (FrameLabel, error) {
	// convert timecode to number of frames
	numFrames, err := AsFrames(timecode)

	var rate FrameRate
	switch frameRate {
	case "NDF25":
		rate = NDF25
	case "DF30":
		rate = DF30
	default:
		return FrameLabel{}, fmt.Errorf("Invalid frame rate: ", frameRate)
	}
	return FrameLabel{numFrames, rate}, nil
}

// ParseFrames creates a new FrameLabel from uint64 number of frames
func (fl FrameLabel) ParseFrames(frameRate string, numFrames uint64) (FrameLabel, error) {

}

type AmountOfTime struct {
	FrameCount
	FrameRate
}

// Text returns frame as string (e.g., 01:00:00:00)
func (aot AmountOfTime) Text() string {
	return ""
}

// NumFrames returns frame as number of frames (e.g., 90000)
func (aot AmountOfTime) NumFrames() uint64 {
	return 0
}

// Rate returns the Frame's frame rate
func (aot AmountOfTime) Rate() string {
	return ""
}

func main() {
	fmt.Printf("%s: %d, Drop Frame: %t\n", NDF25.Name(), NDF25.Rate(), NDF25.DropFrame())
	fmt.Printf("%s: %d, Drop Frame: %t\n", DF30.Name(), DF30.Rate(), DF30.DropFrame())
	var newLabel FrameLabel
	newLabel, err := newLabel.ParseText("NDF25", "01:30:11:12")
}
