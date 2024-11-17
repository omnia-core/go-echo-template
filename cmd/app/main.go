package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"

	_ "github.com/omnia-core/go-echo-template/docs"
	userRouter "github.com/omnia-core/go-echo-template/internal/user/router"
	userStore "github.com/omnia-core/go-echo-template/internal/user/store"
	userUsecase "github.com/omnia-core/go-echo-template/internal/user/usecase"
	"github.com/omnia-core/go-echo-template/pkg/config"
	"github.com/omnia-core/go-echo-template/pkg/db"
	"github.com/omnia-core/go-echo-template/pkg/log"
	echoRouter "github.com/omnia-core/go-echo-template/pkg/router/echo"
)

var (
	GitInfo = "default git info"
	Tag     = "v0.0.0"

	Commit = "default commit"
	Branch = "default branch"
)

func setupLogger() {
	log.Logger = logrus.New()
	log.Logger.SetFormatter(&logrus.JSONFormatter{})
}

func versionPrint() {
	fmt.Println("========================================")
	fmt.Printf("[version]:%v\n", GitInfo)
	fmt.Printf("[tag]:%v\n", Tag)
	fmt.Println("========================================")

	log.Branch = Branch
	log.Commit = Commit

	serverIp, _ := exec.Command("hostname").Output()
	echoRouter.ServerIp = string(serverIp)
	echoRouter.Branch = Branch
	echoRouter.Commit = Commit
}

// @query.collection.format multi
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	setupLogger()
	versionPrint()

	app := fx.New(
		fx.Provide(
			config.New,
			db.NewPostgresqlDB,
			echoRouter.New,

			userStore.NewUserStore,
			userUsecase.NewUserUsecase,
		),
		fx.Invoke(
			userRouter.NewUserRouter,
			serve,
		),
	)
	startCtx, cancel := context.WithTimeout(context.Background(), 600*time.Second)
	defer cancel()

	if err := app.Start(startCtx); err != nil {
		log.New().Fatalf("%v", err)
	}

	<-app.Done()

}

func serve(lifecycle fx.Lifecycle, cfg *config.Config, echo *echo.Echo) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := echo.Start(cfg.Listen); err != nil {
					log.New().Fatalf("err start echo server: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.New().Fatalf("stopping server")
			return nil
		},
	})
}
