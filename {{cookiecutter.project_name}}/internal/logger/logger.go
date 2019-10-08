package logger

import (
	"github.com/mdevilliers/go/env"
	helper "github.com/mdevilliers/go/logger"
	"github.com/rs/zerolog"
)

func NewFromEnvironment(fields map[string]interface{}) zerolog.Logger {
	logLevel := env.FromEnvWithDefaultStr("{{ cookiecutter.project_name | upper }}_LEVEL", "info")
	useConsole := env.FromEnvWithDefaultBool("{{ cookiecutter.project_name | upper }}_LOG_USE_CONSOLE", false)
	return helper.New(logLevel, useConsole, fields)
}
