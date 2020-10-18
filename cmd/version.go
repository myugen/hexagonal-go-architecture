package cmd

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print hexagonal architecture app version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("app version %s", viper.GetString("version"))
		},
	}
)
