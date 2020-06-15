package cmd

import (
	"github.com/bancodobrasil/stop-analyzing-api/service"
	"github.com/bancodobrasil/stop-analyzing-api/service/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the HTTP REST APIs server",
	RunE: func(cmd *cobra.Command, args []string) error {
		builder := new(config.ServiceBuilder).Init(viper.GetViper())
		server := new(service.Server).InitFromServiceBuilder(builder)
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
