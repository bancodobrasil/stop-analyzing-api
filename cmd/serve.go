package cmd

import (
	"github.com/bancodobrasil/stop-analyzing-api/internal/api"
	"github.com/bancodobrasil/stop-analyzing-api/internal/api/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the HTTP REST APIs server",
	RunE: func(cmd *cobra.Command, args []string) error {
		builder := new(config.APIBuilder).Init(viper.GetViper())
		server := new(api.Server).InitFromAPIBuilder(builder)
		server.Run()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	config.AddFlags(serveCmd.Flags())

	err := viper.GetViper().BindPFlags(serveCmd.Flags())
	if err != nil {
		panic(err)
	}
}
