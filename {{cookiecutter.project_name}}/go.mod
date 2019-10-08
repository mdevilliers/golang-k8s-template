module github.com/{{cookiecutter.github_account}}/{{cookiecutter.project_name}}

go {{cookiecutter.golang_version}}

require (
	github.com/rs/zerolog v1.14.3
	github.com/spf13/cobra v0.0.3
	github.com/spf13/pflag v1.0.3
	github.com/spf13/viper v1.3.1
	github.com/stretchr/testify v1.3.0
	github.com/mdevilliers/go/cli v0.0.0-20191004091409-4c0b0240794b
	github.com/mdevilliers/go/env v0.0.0-20191004091409-4c0b0240794b
	github.com/mdevilliers/go/healthchecks v0.0.0-20191004101804-7c6994f2d6cb
	github.com/mdevilliers/go/logger v0.0.0-20191004091409-4c0b0240794b
)
