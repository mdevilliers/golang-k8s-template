package logger

import (
	"io"
	"os"

	"github.com/{{cookiecutter.github_account}}/{{cookiecutter.project_name}}/internal/version"
	"github.com/rs/zerolog"
)

// New returns an instntiated, configured logger
// This really should be in a shared library!
func New(logLevel string, useConsole, makeVerbose bool) zerolog.Logger {

	var w io.Writer = os.Stdout
	if useConsole {
		w = zerolog.ConsoleWriter{
			Out: os.Stdout,
		}
	}
	// Parse level from config
	lvl, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		lvl = zerolog.InfoLevel
	}
	// Override level with verbose
	if makeVerbose {
		lvl = zerolog.DebugLevel
	}
	return zerolog.New(w).Level(lvl).With().Fields(map[string]interface{}{
		"version": version.Version,
		"app":     "cache-service",
	}).Timestamp().Logger()
}
