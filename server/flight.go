package server

import (
	"errors"
)

var (
	ComputeOriginErr           = errors.New("invalid flight plan provided: unable to compute origin")
	ComputeFinalDestinationErr = errors.New("invalid flight plan provided: unable to compute final destination")
)

type flightSegment struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
}

type flightItinerary struct {
	Origin           string           `json:"origin"`
	FinalDestination string           `json:"finalDestination"`
	Segments         []*flightSegment `json:"segments"`
}

func (i *flightItinerary) isNonStop() bool {
	return len(i.Segments) == 1
}

func (i *flightItinerary) computeOrigin() error {
	if i.isNonStop() {
		i.Origin = i.Segments[0].Origin
		return nil
	}
	originCount := make(map[string]int, len(i.Segments))
	for _, segment := range i.Segments {
		if segment.Origin == segment.Destination {
			return ComputeOriginErr
		}
		// count the occurrence of origin and destination airport codes to determine connections
		originCount[segment.Origin]--
		originCount[segment.Destination]++
	}
	for key, count := range originCount {
		if count == -1 {
			i.Origin = key
		}
	}
	if i.Origin == "" {
		return ComputeOriginErr
	}
	return nil
}

func (i *flightItinerary) computeFinalDestination() error {
	if i.isNonStop() {
		i.FinalDestination = i.Segments[0].Destination
		return nil
	}
	destinationCount := make(map[string]int, len(i.Segments))
	for _, segment := range i.Segments {
		if segment.Origin == segment.Destination {
			return ComputeFinalDestinationErr
		}
		// count the occurrence of origin and destination airport codes to determine connections
		destinationCount[segment.Destination]--
		destinationCount[segment.Origin]++
	}
	for key, count := range destinationCount {
		if count == -1 {
			i.FinalDestination = key
		}
	}
	if i.FinalDestination == "" {
		return ComputeFinalDestinationErr
	}
	return nil
}
