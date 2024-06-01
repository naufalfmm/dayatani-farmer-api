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
	"github.com/gin-gonic/gin/binding"
	"github.com/naufalfmm/dayatani-farmer-api/infrastructures"
	"github.com/naufalfmm/dayatani-farmer-api/middlewares"
	"github.com/naufalfmm/dayatani-farmer-api/persistents"
	"github.com/naufalfmm/dayatani-farmer-api/resources/config"
	"github.com/naufalfmm/dayatani-farmer-api/resources/db"
	"github.com/naufalfmm/dayatani-farmer-api/resources/log"
	"github.com/naufalfmm/dayatani-farmer-api/resources/validator"
	"github.com/naufalfmm/dayatani-farmer-api/usecases"

	"github.com/naufalfmm/dayatani-farmer-api/utils/encoding/base64Encoding"
	"github.com/naufalfmm/dayatani-farmer-api/utils/hashing/bcryptHash"
	validatorUtils "github.com/naufalfmm/dayatani-farmer-api/utils/validator"
)

type App struct {
	ge *gin.Engine

	c *config.EnvConfig

	validator validatorUtils.Validator
}

func Init() App {
	ge := gin.New()

	c, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	l, err := log.NewLogger(c)
	if err != nil {
		panic(err)
	}

	d, err := db.NewPostgres(c, l)
	if err != nil {
		panic(err)
	}

	prst, err := persistents.Init(d, l)
	if err != nil {
		panic(err)
	}

	uscs, err := usecases.Init(prst)
	if err != nil {
		panic(err)
	}

	bcrHash, err := bcryptHash.NewBcrypt(bcryptHash.WithCost(c.HashedCost))
	if err != nil {
		panic(err)
	}

	base64Enc, err := base64Encoding.NewBase64Encoding(base64Encoding.WithEncType(c.Base64EncodingType))
	if err != nil {
		panic(err)
	}

	vlds, err := validator.NewValidator()
	if err != nil {
		panic(err)
	}

	middls, err := middlewares.Init(bcrHash, base64Enc, c)
	if err != nil {
		panic(err)
	}

	infrs, err := infrastructures.Init(uscs, middls)
	if err != nil {
		panic(err)
	}

	infrs.Register(ge)

	return App{
		ge:        ge,
		c:         c,
		validator: vlds,
	}
}

func (app App) Run() {
	binding.Validator = app.validator

	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", app.c.Port),
		Handler: app.ge,
	}

	go httpServer.ListenAndServe() //nolint:errcheck

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
