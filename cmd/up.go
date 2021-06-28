package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/myugen/hexagonal-go-architecture/infrastructure/echo/errors"

	"github.com/myugen/hexagonal-go-architecture/internal/articles/app"

	"github.com/myugen/hexagonal-go-architecture/infrastructure/echo/middlewares"

	"github.com/myugen/hexagonal-go-architecture/infrastructure/logger"

	"github.com/spf13/viper"

	"github.com/myugen/hexagonal-go-architecture/infrastructure/postgres"

	"github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"

	"github.com/myugen/hexagonal-go-architecture/utils/constants"
	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: fmt.Sprintf("Up %s server", constants.AppName),
	Run:   runUp,
}

func runUp(_ *cobra.Command, _ []string) {
	logger.Log().Debug("Connecting postgres database")
	if err := postgres.Initialize(); err != nil {
		logger.Log().Fatalf("Error on database connection: %s", err)
	}
	serverConfig := viper.GetStringMapString("server")
	logger.Log().Debugf("Starting API Server on port: %s", serverConfig["port"])
	server := setupServer()

	go func() {
		if err := server.Start(fmt.Sprintf(":%s", serverConfig["port"])); err != nil {
			logger.Log().Fatalf("Error on server: %s", err)
		}
	}()

	graceful(server, 2*time.Second)
}

func setupServer() *echo.Echo {
	e := echo.New()
	e.Logger = logger.Log()
	e.Use(middleware.CORS())
	e.Use(middleware.Gzip())
	if viper.GetBool("verbose") {
		e.Use(middlewares.Logger())
	}
	e.HTTPErrorHandler = errors.HTTPErrorHandler

	apiRoute := e.Group("/api")

	app.SetupArticleApp(apiRoute)
	return e
}

func graceful(server *echo.Echo, timeout time.Duration) {
	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	logger.Log().Info("Server shutdown")
	if err := server.Shutdown(ctx); err != nil {
		logger.Log().Errorf("Server shutdown: %s\n", err)
	}

	// After shutdown operations
	//---------------------------
	logrus.Info("Closing connections after shutdown")
	if err := postgres.Close(); err != nil {
		logger.Log().Error("Database shutdown error: %s\n", err)
	}
	//---------------------------
}
