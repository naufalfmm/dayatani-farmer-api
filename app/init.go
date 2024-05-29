package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/naufalfmm/dayatani-farmer-api/resources/config"
)

type App struct {
	ge *gin.Engine

	c *config.EnvConfig
}

func Init() App {
	ge := gin.New()

	c, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	return App{
		ge: ge,
		c:  c,
	}
}

func (app App) Run() {
	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", app.c.Port),
		Handler: app.ge,
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(
		sc,
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	<-sc

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctxShutDown); err != nil {
		panic(err)
	}
}
