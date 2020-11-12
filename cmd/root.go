package cmd

import (
	"fmt"

	"github.com/myugen/hexagonal-go-architecture/utils/constants"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"os"
)

var (
	configFile  string
	userLicense string
	rootCmd     = &cobra.Command{
		Use:   constants.AppLabel,
		Short: fmt.Sprintf("Root command for %s", constants.AppName),
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", fmt.Sprintf("config file path (default lookups [./config.yml, $HOME/.%s/config.yml, /etc/%s/config.yml])", constants.AppLabel, constants.AppLabel))
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "MIT", "name of license for the project")
	viper.SetDefault("author", "Miguel Cabrera <me@mcabsan.dev>")
	viper.SetDefault("license", "MIT")
	viper.SetDefault("version", "0.0.1")
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(upCmd)
}

func initConfig() {
	if configFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(configFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		// Search config in home directory with name ".config" (without extension).
		viper.AddConfigPath(fmt.Sprintf("/etc/%s/", constants.AppLabel)) // path to look for the config file in
		viper.AddConfigPath(fmt.Sprintf("%s/.%s", home, constants.AppLabel))
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
	}

	viper.SetEnvPrefix(constants.AppLabel)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
