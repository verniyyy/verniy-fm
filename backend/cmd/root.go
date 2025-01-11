package cmd

import (
	"os"
	"verniy-fm-backend/internal/presenter"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "verniy-fm-backend",
	RunE: func(cmd *cobra.Command, args []string) error {
		s := presenter.NewHTTPServer(8080)
		s.Serve(cmd.Context())

		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
