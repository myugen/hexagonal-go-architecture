package cmd

import (
	"fmt"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"os"
)

var (
	configFile  string
	userLicense string
	rootCmd     = &cobra.Command{
		Use:   "app",
		Short: "Root command for hexagonal architecture app",
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

	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is $HOME/.app/config.yaml)")
	rootCmd.PersistentFlags().StringP("author", "a", "Miguel Cabrera", "author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "apache", "name of license for the project")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "Miguel Cabrera <me@mcabsan.dev>")
	viper.SetDefault("license", "apache")
	viper.SetDefault("version", "0.0.1")
	rootCmd.AddCommand(versionCmd)
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
		viper.AddConfigPath("/etc/app/") // path to look for the config file in
		viper.AddConfigPath(fmt.Sprintf("%s/.app", home))
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
