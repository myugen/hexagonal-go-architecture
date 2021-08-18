package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/myugen/hexagonal-go-architecture/infrastructure/config"

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
	viper.SetDefault("mode", "development")
	viper.SetDefault("version", "0.0.1")
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("db.host", "0.0.0.0")
	viper.SetDefault("db.port", "5432")
	viper.SetDefault("db.database", "dev")
	viper.SetDefault("db.user", "dev")
	viper.SetDefault("db.password", "changeme")
	viper.SetDefault("verbose", false)
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
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			fmt.Println("No Config file found, loaded config from Environment - Default path ./conf")
		default:
			log.Fatalf("Error when Fetching Configuration - %s", err)
		}
	}

	var appConfig config.AppConfig
	if err := viper.Unmarshal(&appConfig); err != nil {
		log.Fatalf("Error when Fetching Configuration - %s", err)
	}
	config.Initialize(appConfig)
}
