package main

import (
	"log"

	"github.com/purini-to/zapmw"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/koneal2013/flightify/middleware"
	"github.com/koneal2013/flightify/server"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
)

type cli struct {
	cfg server.Config
}

func (c *cli) setupConfig(cmd *cobra.Command, args []string) error {
	if configFile, err := cmd.Flags().GetString("config-file"); err != nil {
		return err
	} else {
		viper.SetConfigFile(configFile)
		if err = viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
				return err
			}
		}
		c.cfg.Port = viper.GetInt("port")
		c.cfg.IsDevelopment = viper.GetBool("is-development")
		if viper.GetBool("enable-logging-middleware") {
			// log each request with the global zap logger (initialized in server.NewHTTPServer)
			c.cfg.MiddlewareFuncs = append(c.cfg.MiddlewareFuncs, zapmw.WithZap(zap.L()), zapmw.Request(zapcore.InfoLevel, "processing request"), zapmw.Recoverer(zapcore.ErrorLevel, "recover", zapmw.RecovererDefault))
		}
		c.cfg.MiddlewareFuncs = append(c.cfg.MiddlewareFuncs, middleware.Recovery)
	}
	return nil
}

func (c *cli) run(cmd *cobra.Command, args []string) error {
	srv := server.NewHTTPServer(c.cfg)
	return srv.ListenAndServe()
}

func setupFlags(cmd *cobra.Command) error {
	cmd.Flags().String("config-file", "", "Path to config file.")
	cmd.Flags().Bool("is-development", true, "Flag to set log level.")
	cmd.Flags().Int("port", 8080, "Port to serve on.")
	cmd.Flags().Bool("enable-logging-middleware", true, "Enable logging of each request")
	return viper.BindPFlags(cmd.Flags())
}

//	@title			Flightify API
//	@version		1.0
//	@description	Api documentation for Flightify service.

//	@contact.name	Kenston O'Neal
//	@contact.email	koneal2013@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host	localhost:8080
func main() {
	cli := &cli{}

	cmd := &cobra.Command{
		Use:     "flightify",
		PreRunE: cli.setupConfig,
		RunE:    cli.run,
	}
	if err := setupFlags(cmd); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
