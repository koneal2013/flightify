package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/swaggo/http-swagger"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"

	"github.com/koneal2013/flightify/adaptor"
	"github.com/koneal2013/flightify/flight"
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
	r.HandleFunc("/calculate", adaptor.GenericHttpAdaptor(handleCalculate)).Methods("POST")
	r.HandleFunc("/status", handleStatus).Methods("GET")
	r.HandleFunc("/swagger/*", httpSwagger.Handler(httpSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", cfg.Port)))).Methods("GET")
	r.Use(cfg.MiddlewareFuncs...)
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

func handleCalculate(ctx context.Context, in [][]string) (out []string, err error) {
	f, err := flight.New(in)
	if err != nil {
		return
	}
	if err = f.ComputeOrigin(); err != nil {
		return
	} else if err = f.ComputeFinalDestination(); err != nil {
		return
	}
	out = []string{f.Origin, f.FinalDestination}
	return
}
