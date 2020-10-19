package cmd

import (
	"fmt"

	"github.com/myugen/hexagonal-go-architecture/utils/constants"
	"github.com/spf13/cobra"
)

var (
	upCmd = &cobra.Command{
		Use:   "up",
		Short: fmt.Sprintf("Up %s server", constants.AppName),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Done!")
		},
	}
)
