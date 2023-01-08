package flight

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestComputeOriginAndFinalDestination(t *testing.T) {
	for _, tc := range []struct {
		name           string
		input          *Itinerary
		expectedOutput *Itinerary
		expectedErr    error
	}{
		{
			name: "failure invalid input for origin",
			input: &Itinerary{
				Segments: []*Segment{
					{
						Origin:      "DFW",
						Destination: "DFW",
					},
					{
						Origin:      "SFO",
						Destination: "CHA",
					},
				},
			},
			expectedErr: computeOriginErr,
		},
		{
			name: "failure not non stop / no connections found",
			input: &Itinerary{
				Segments: []*Segment{
					{
						Origin:      "IND",
						Destination: "EWR",
					},
					{
						Origin:      "DFW",
						Destination: "SFO",
					},
				},
			},
			expectedErr: noConnectionsErr,
		},
		{
			name: "failure invalid input for final Destination",
			input: &Itinerary{
				Segments: []*Segment{
					{
						Origin:      "SFO",
						Destination: "DFW",
					},
					{
						Origin:      "DFW",
						Destination: "DFW",
					},
				},
			},
			expectedErr: computeFinalDestinationErr,
		},
		{
			name: "success non stop itinerary",
			input: &Itinerary{
				Segments: []*Segment{
					{
						Origin:      "SFO",
						Destination: "EWR",
					},
				},
			},
			expectedOutput: &Itinerary{Origin: "SFO", FinalDestination: "EWR"},
		},
		{
			name: "success itinerary with one connection",
			input: &Itinerary{
				Segments: []*Segment{
					{
						Origin:      "ATL",
						Destination: "EWR",
					},
					{
						Origin:      "SFO",
						Destination: "ATL",
					},
				},
			},
			expectedOutput: &Itinerary{Origin: "SFO", FinalDestination: "EWR"},
		},
		{
			name: "success itinerary with two connections",
			input: &Itinerary{
				Segments: []*Segment{
					{
						Origin:      "DFW",
						Destination: "CHA",
					},
					{
						Origin:      "FAT",
						Destination: "DFW",
					},
					{
						Origin:      "CHA",
						Destination: "ATL",
					},
				},
			},
			expectedOutput: &Itinerary{Origin: "FAT", FinalDestination: "ATL"},
		},
		{
			name: "success itinerary with more that one connection",
			input: &Itinerary{
				Segments: []*Segment{
					{
						Origin:      "IND",
						Destination: "EWR",
					},
					{
						Origin:      "DFW",
						Destination: "SFO",
					},
					{
						Origin:      "EWR",
						Destination: "CHA",
					},
					{
						Origin:      "SFO",
						Destination: "ATL",
					},
					{
						Origin:      "GSO",
						Destination: "IND",
					},
					{
						Origin:      "DEN",
						Destination: "FAT",
					},
					{
						Origin:      "FAT",
						Destination: "DFW",
					},
					{
						Origin:      "ATL",
						Destination: "GSO",
					},
				},
			},
			expectedOutput: &Itinerary{Origin: "DEN", FinalDestination: "CHA"},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			computeOriginErr := tc.input.ComputeOrigin()
			computeFinalDestinationErr := tc.input.ComputeFinalDestination()

			if tc.expectedOutput != nil {
				require.Equal(t, tc.expectedOutput.Origin, tc.input.Origin)
				require.Equal(t, tc.expectedOutput.FinalDestination, tc.input.FinalDestination)
			} else if tc.expectedErr == computeOriginErr {
				require.Error(t, computeOriginErr)
				require.Equal(t, tc.expectedErr, computeOriginErr)
			} else if tc.expectedErr == computeFinalDestinationErr {
				require.Error(t, computeFinalDestinationErr)
				require.Equal(t, tc.expectedErr, computeFinalDestinationErr)
			}
		})

	}
}
