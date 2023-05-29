package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/cobra"
	"natsmon/common"
	sctx "natsmon/service-context"
	"natsmon/service-context/component/fiberc"
)

const (
	serviceName = "channel_service"
	version     = "1.0.0"
)

func newServiceCtx() sctx.ServiceContext {
	return sctx.NewServiceContext(
		sctx.WithName(serviceName),
		sctx.WithComponent(fiberc.NewFiberComp(common.KeyCompFiber)),
	)
}

var rootCmd = &cobra.Command{
	Use:   serviceName,
	Short: fmt.Sprintf("start %s", serviceName),
	Run: func(cmd *cobra.Command, args []string) {
		serviceCtx := newServiceCtx()

		log := sctx.GlobalLogger().GetLogger("service")

		time.Sleep(time.Second * 5)

		if err := serviceCtx.Load(); err != nil {
			log.Fatal(err)
		}

		fiberComp := serviceCtx.MustGet(common.KeyCompFiber).(fiberc.FiberComponent)

		router := fiberComp.GetRouter()

		router.Use(logger.New(logger.Config{
			Format: `{"ip":${ip}, "timestamp":"${time}", "status":${status}, "latency":"${latency}", "method":"${method}", "path":"${path}"}` + "\n",
		}))

		router.Get("/ping", func(c *fiber.Ctx) error {
			return c.SendString("Hello, World 👋!")
		})

		router.Static("/", "./public")

		if err := router.Listen(fmt.Sprintf(":%d", fiberComp.GetPort())); err != nil {
			log.Error(err)
		}
	},
}

func Execute() {
	rootCmd.AddCommand(outEnvCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
