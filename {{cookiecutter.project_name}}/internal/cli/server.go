package cli

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/emicklei/go-restful"
	"github.com/kelseyhightower/envconfig"
	"github.com/mdevilliers/go/healthchecks"
	"github.com/{{cookiecutter.github_account}}/{{cookiecutter.project_name}}/logger"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type serverConfig struct {
	GRPCBinding            string `envconfig:"GRPC_BINDING" default:":3000"`
	HTTPBinding            string `envconfig:"HTTP_BINDING" default:":8000"`
}

func (o *serverConfig) AddFlags(fs *pflag.FlagSet) {
	fs.StringVarP(&o.GRPCBinding, "grpc-binding", "b", o.GRPCBinding, "GRPC binding address in form of {ip}:port")
	fs.StringVarP(&o.HTTPBinding, "http-binding", "t", o.HTTPBinding, "HTTP binding address in form of {ip}:port")
}

func NewServerCmd() (*cobra.Command, error) {

	runtimeConfig := &serverConfig{}
	err := envconfig.Process("", runtimeConfig)
	if err != nil {
		return nil, errors.Wrap(err, "error initilising config")
	}

	cmd := &cobra.Command{
		Use: "server",
		RunE: func(cmd *cobra.Command, args []string) error {

			ctx := context.Background()

			lgr := logger.NewFromEnvironment(map[string]interface{}{
				"app": "server",
			})

			health := healthchecks.New(hasuraClient)
			go func() {
				if err := health.Start(":8086"); err != nil {
					lgr.Fatal().Err(err).Msg("error running healthcheck HTTP server")
				}
			}()

			// TODO : start server

			stop := make(chan os.Signal, 1)
			signal.Notify(stop, os.Interrupt)

			<-stop

			// TODO : shutdown server

			return nil
		},
	}

	runtimeConfig.AddFlags(cmd.PersistentFlags())
	return cmd, nil
}
