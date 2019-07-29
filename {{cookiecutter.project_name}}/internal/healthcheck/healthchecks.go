package healthcheck

import (
	"errors"
	"math/rand"
	"net/http"

	"github.com/heptiolabs/healthcheck"
	"github.com/rs/zerolog"
)

// Start initiates the healthchecks as an HTTP server
func Start(logger zerolog.Logger) {

	health := healthcheck.NewHandler()

	// life is too easy without some random failures
	// This is only here for demonstration purposes
	health.AddLivenessCheck("random-failure", func() error {

		r := rand.Intn(1000)

		logger.Info().Fields(map[string]interface{}{
			"random": r,
		}).Msg("random failure called")
		if r == 1 {
			return errors.New("boom")
		}
		return nil

	})

	// nolint : errcheck
	go http.ListenAndServe(":8086", health)

}
