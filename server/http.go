package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/swaggo/http-swagger"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
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
	r.HandleFunc("/swagger/*", httpSwagger.Handler(httpSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", cfg.Port)))).Methods("GET")
	for _, middlewareFunc := range cfg.MiddlewareFuncs {
		r.Use(middlewareFunc)
	}
	return &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: r,
	}
}

// Status godoc
//
//	@Description	Return 200 OK if server is ready to accept requests
//	@Success		200	{object}	string
//
//	@Router			/status [get]
func handleStatus(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// Calculate godoc
//
//	@Description	Returns a flight itinerary with origin airport code and final destination airport code
//	@Accept			json
//	@Produce		json
//	@Param			input	body		string	true	"slice of flight segments (i.e. [['ATL', 'EWR'], ['SFO', 'ATL']])"
//	@Success		200		{object}	string
//	@Failure		400		{object}	string
//	@Router			/calculate [post]
func handleCalculate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var input [][]string
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	flight := &flightItinerary{}
	for _, val := range input {
		origin := val[0]
		dest := val[1]
		flight.Segments = append(flight.Segments, &flightSegment{
			Origin:      origin,
			Destination: dest,
		})
	}
	if err := flight.computeOrigin(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if err := flight.computeFinalDestination(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		output := make([][]string, 1)
		output[0] = make([]string, 2)
		output[0][0] = flight.Origin
		output[0][1] = flight.FinalDestination
		w.Header().Add("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(output)
	}
}
