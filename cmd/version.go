package cmd

import (
	"fmt"

	"github.com/myugen/hexagonal-go-architecture/utils/constants"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: fmt.Sprintf("Print %s version", constants.AppName),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s version %s", constants.AppLabel, viper.GetString("version"))
		},
	}
)
