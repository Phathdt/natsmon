package cmd

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"natsmon/common"
	"natsmon/modules/natstransport/fibernats"
	sctx "natsmon/service-context"
	"natsmon/service-context/component/ginc"
	"natsmon/service-context/component/ginc/middleware"
	"natsmon/service-context/component/natsc"
	"natsmon/service-context/component/natspub"
)

const (
	serviceName = "natsmon"
	version     = "1.0.0"
)

func newServiceCtx() sctx.ServiceContext {
	return sctx.NewServiceContext(
		sctx.WithName(serviceName),
		sctx.WithComponent(ginc.NewGin(common.KeyCompGin)),
		sctx.WithComponent(natsc.NewNatsComp(common.KeyNatsComp)),
		sctx.WithComponent(natspub.NewNatsPubComponent(common.KeyNatsPubComp)),
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

		ginComp := serviceCtx.MustGet(common.KeyCompGin).(ginc.GinComponent)

		router := ginComp.GetRouter()

		router.Use(gin.Recovery(), gin.Logger(), middleware.Recovery(serviceCtx))
		router.Use(static.Serve("/", static.LocalFile("./public", true)))

		router.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		api := router.Group("/api")
		{
			jetstreams := api.Group("/jetstreams")
			{
				jetstreams.GET("", fibernats.ListJetstream(serviceCtx))
				jetstreams.GET("/:stream", fibernats.GetStream(serviceCtx))
				jetstreams.GET("/:stream/consumers", fibernats.GetConsumer(serviceCtx))
				jetstreams.GET("/:stream/messages", fibernats.GetMessages(serviceCtx))
				jetstreams.POST("/:stream/messages", fibernats.AddMessage(serviceCtx))
			}
		}

		if err := router.Run(fmt.Sprintf(":%d", ginComp.GetPort())); err != nil {
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
