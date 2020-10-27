package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4/middleware"

	"github.com/myugen/hexagonal-go-architecture/pkg/articles/adapters/api/routes"

	"github.com/labstack/echo/v4"

	"github.com/myugen/hexagonal-go-architecture/utils/constants"
	"github.com/spf13/cobra"
)

var (
	upCmd = &cobra.Command{
		Use:   "up",
		Short: fmt.Sprintf("Up %s server", constants.AppName),
		Run:   run,
	}
)

func run(cmd *cobra.Command, args []string) {
	server := setupServer()

	go func() {
		if err := server.Start(":8181"); err != nil {
			logrus.Fatal("Error setting up the server: %s", err)
		}
	}()

	graceful(server, 2*time.Second)
}

func setupServer() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Gzip())

	routes.RegisterRoute(e)
	return e
}

func graceful(server *echo.Echo, timeout time.Duration) {
	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	logrus.Info("Server shutdown")
	if err := server.Shutdown(ctx); err != nil {
		logrus.Errorf("Server shutdown: %s\n", err)
	}

	// After shutdown operations
	//---------------------------
	logrus.Info("Closing connections after shutdown")
	//---------------------------
}
