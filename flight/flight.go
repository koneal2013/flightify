package flight

import (
	"errors"
)

var (
	computeOriginErr           = errors.New("invalid flight plan provided: unable to compute origin")
	computeFinalDestinationErr = errors.New("invalid flight plan provided: unable to compute final destination")
	noConnectionsErr           = errors.New("invalid flight plan provided: flight itinerary is not non stop and one or more connections are missing")
)

type Segment struct {
	Origin      string `json:"0"`
	Destination string `json:"1"`
}

type Itinerary struct {
	Origin           string     `json:"origin"`
	FinalDestination string     `json:"finalDestination"`
	Segments         []*Segment `json:"segments"`
}

func New(segments [][]string) (*Itinerary, error) {
	if len(segments) <= 0 {
		return nil, errors.New("no segments provided")
	}
	itinerary := &Itinerary{}
	for _, segment := range segments {
		newSegment := &Segment{
			Origin:      segment[0],
			Destination: segment[1],
		}
		itinerary.Segments = append(itinerary.Segments, newSegment)
	}
	return itinerary, nil
}

func (i *Itinerary) isNonStop() bool {
	return len(i.Segments) == 1
}

// ComputeOrigin finds the flight itinerary's origin from a list of flight segments
func (i *Itinerary) ComputeOrigin() error {
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

// ComputeFinalDestination finds the flight itinerary's final destination from a list of flight segments
func (i *Itinerary) ComputeFinalDestination() error {
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
	return nil
}
