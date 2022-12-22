package main

import (
	"log"

	"github.com/koneal2013/flightify/server"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	}
	return nil
}

func (c *cli) run(cmd *cobra.Command, args []string) error {
	srv := server.NewHTTPServer(c.cfg)
	return srv.ListenAndServe()
}

func setupFlags(cmd *cobra.Command) error {
	cmd.Flags().String("config-file", "", "Path to config file.")
	cmd.Flags().Bool("is-development", false, "Flag to set log level.")
	cmd.Flags().Int("port", 8080, "Port to serve on.")
	return viper.BindPFlags(cmd.Flags())
}

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
