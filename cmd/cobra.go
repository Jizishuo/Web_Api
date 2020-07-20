package cmd

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"os"
	"github.com/spf13/cobra"
	"Web_Api/cmd/api"
)

var rootCmd = &cobra.Command{
	Use: "go-admin",
	Short: "-v",
	SilenceUsage: true,
	DisableAutoGenTag: true,
	Long: `go-admin`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		return nil
	},
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		usageStr := `web-api welcome, use -h `
		log.Printf("%\n", usageStr)
	},
}

func init()  {
	rootCmd.AddCommand(api.StartCmd)
	// rootCmd.AddCommand(migrate.StartCmd)
}

func Execute()  {
	if err := rootCmd.Execute();err!=nil{
		os.Exit(-1)
	}
}