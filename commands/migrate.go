package commands

import (
	"bifrost/gen/entmodels/migrate"
	"bifrost/internal/app"
	"context"
	"github.com/spf13/cobra"
	"log"
)

func init()  {
	cmd := &cobra.Command{
		Use: "migrate",
		Short: "Migrate database when change database schemas",
		Run: databaseMigrate,
	}
	rootCmd.AddCommand(cmd)
}

func databaseMigrate(_ *cobra.Command, _[]string)  {
	ctx := context.Background()
	err := app.EntClient.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}