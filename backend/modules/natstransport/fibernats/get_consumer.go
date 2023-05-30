package fibernats

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"natsmon/common"
	"natsmon/modules/natsbiz"
	"natsmon/modules/natsrepo"
	sctx "natsmon/service-context"
	"natsmon/service-context/component/natsc"
	"natsmon/service-context/core"
)

func GetConsumer(sc sctx.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		stream := c.Param("stream")

		natsComponent := sc.MustGet(common.KeyNatsComp).(natsc.NatsComponent)
		manager := natsComponent.GetManager()
		js := natsComponent.GetJs()

		repo := natsrepo.NewRepo(manager, js)
		biz := natsbiz.NewGetConsumerBiz(repo)

		consumers, err := biz.Response(c, stream)
		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(consumers))
	}
}
