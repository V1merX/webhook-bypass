package app

import (
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/V1merX/webhook-bypass/internal/handlers"
	"github.com/gin-gonic/gin"
)

func Run() error {
	router := gin.Default()

	router.POST("/", handlers.Route)

	srv := &http.Server{
		Addr: ":9090",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			slog.Error("failed to start server", slog.String("err", err.Error()))
			return
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	<-signalChan

	slog.Info("shutdown api...")

	return nil
}
