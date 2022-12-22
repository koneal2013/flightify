package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func NewHTTPServer(cfg Config) *http.Server {
	if cfg.IsDevelopment {
		if logger, err := zap.NewDevelopment(); err != nil {
			log.Fatal(err)
		} else {
			zap.ReplaceGlobals(logger)
		}
	} else {
		if logger, err := zap.NewProduction(); err != nil {
			log.Fatal(err)
		} else {
			zap.ReplaceGlobals(logger)
		}
	}
	r := mux.NewRouter()
	r.HandleFunc("/calculate", handleCalculate).Methods("POST")
	r.HandleFunc("/status", handleStatus).Methods("GET")
	for _, middlewareFunc := range cfg.MiddlewareFuncs {
		r.Use(middlewareFunc)
	}
	return &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: r,
	}
}

func handleStatus(w http.ResponseWriter, r *http.Request) {

}

func handleCalculate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if data, err := io.ReadAll(r.Body); err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
	} else {
		var input [][]string
		json.Unmarshal(data, &input)
		flight := &flightItinerary{}
		for _, val := range input {
			if origin, dest, ok := strings.Cut(val[0], ","); !ok {
				http.Error(w, "error parsing the input provided", http.StatusBadRequest)
				return
			} else {
				flight.Segments = append(flight.Segments, &flightSegment{
					Origin:      origin,
					Destination: dest,
				})
			}
		}
		if err := flight.computeOrigin(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else if err := flight.computeFinalDestination(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			output := make([][]string, 1)
			output[0] = make([]string, 1)
			output[0][0] = fmt.Sprintf("%s,%s", flight.Origin, flight.FinalDestination)
			w.Header().Add("Content-Type", "application/json")
			outputBytes, _ := json.Marshal(output)
			w.Write(outputBytes)
		}
	}
}
