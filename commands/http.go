package commands

import (
	"bifrost/internal/router"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	cmd := &cobra.Command{
		Use: "start_http",
		Short: "Start http web server",
		Run: startHttp,
	}
	rootCmd.AddCommand(cmd)
}


func startHttp(_ *cobra.Command, _ []string){
	r := router.Routers()
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}

}