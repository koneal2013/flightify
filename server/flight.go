package server

import (
	"errors"
)

var (
	computeOriginErr           = errors.New("invalid flight plan provided: unable to compute origin")
	computeFinalDestinationErr = errors.New("invalid flight plan provided: unable to compute final destination")
	noConnectionsErr           = errors.New("invalid flight plan provided: flight itinerary is not non stop and one or more connections are missing")
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

// computeOrigin finds the flight itinerary's origin from a list of flight segments
func (i *flightItinerary) computeOrigin() error {
	if i.isNonStop() {
		i.Origin = i.Segments[0].Origin
		return nil
	}
	originCount := make(map[string]int, len(i.Segments))
	for _, segment := range i.Segments {
		if segment.Origin == segment.Destination {
			return computeOriginErr
		}
		// count the occurrence of origin and destination airport codes to determine connections
		originCount[segment.Origin] += 0
		originCount[segment.Destination]++
	}
	for key, count := range originCount {
		if count == 0 {
			i.Origin = key
		}
	}
	if i.Origin == "" {
		return computeOriginErr
	}
	return nil
}

// computeFinalDestination finds the flight itinerary's final destination from a list of flight segments
func (i *flightItinerary) computeFinalDestination() error {
	if i.isNonStop() {
		i.FinalDestination = i.Segments[0].Destination
		return nil
	}
	destinationCount := make(map[string]int, len(i.Segments))
	for _, segment := range i.Segments {
		if segment.Origin == segment.Destination {
			return computeFinalDestinationErr
		}
		// count the occurrence of origin and destination airport codes to determine connections
		destinationCount[segment.Destination] += 0
		destinationCount[segment.Origin]++
	}
	var finalDestination []string
	for key, count := range destinationCount {
		if count == 0 {
			finalDestination = append(finalDestination, key)
		}
	}
	if len(finalDestination) != 1 {
		return noConnectionsErr
	}
	i.FinalDestination = finalDestination[0]
	if i.FinalDestination == "" {
		return computeFinalDestinationErr
	}
	return nil
}
