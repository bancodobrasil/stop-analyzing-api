package cmd

import (
	"github.com/bancodobrasil/stop-analyzing-api/internal/migration"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate <url or path>",
	Short: "Import data from external source",
	Long: `
The migrate command instructs the application to load database information from an external source.
The external source can be a local file in json format or a json payload hosted on an external server.
	
Example of usage:

From json url:
	./stop-analyzing-api migrate https://my-server.com/migrate.json

From filesystem:
	./stop-analyzing-api migrate /home/user/migration.txt

In both cases is possible to start from a clean database using the '--recreate-database' flag:
	./stop-analyzing-api migrate https://my-server.com/migrate.json --recreate-database

	`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		recreateDB := viper.GetViper().GetBool("recreate-database")

		return migration.Do(args[0], recreateDB)
	},
}

func init() {

	migrateCmd.Flags().Bool("recreate-database", false, "[optional] Recreates the entire database before loading migration")

	err := viper.GetViper().BindPFlags(migrateCmd.Flags())
	if err != nil {
		panic(err)
	}

	rootCmd.AddCommand(migrateCmd)
}
