package cmd

import (
	"fmt"

	"github.com/myugen/hexagonal-go-architecture/logger"
	"github.com/sirupsen/logrus"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"os"
)

var (
	Log         *logrus.Entry
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
	initLogger()

	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file path (default lookups [./config.yml, $HOME/.app/config.yml, /etc/app/config.yml])")
	rootCmd.PersistentFlags().StringP("author", "a", "Miguel Cabrera", "author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "apache", "name of license for the project")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "Miguel Cabrera <me@mcabsan.dev>")
	viper.SetDefault("license", "apache")
	viper.SetDefault("version", "0.0.1")
	viper.SetDefault("verbose", false)
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
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func initLogger() {
	log := logger.New()
	if viper.GetBool("verbose") {
		log.Logger.SetLevel(logrus.DebugLevel)
	} else {
		log.Logger.SetLevel(logrus.InfoLevel)
	}

	entry := &logrus.Entry{Logger: log.Logger}
	Log = entry.WithField("module", "cmd")
}
