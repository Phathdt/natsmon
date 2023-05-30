package fibernats

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"natsmon/common"
	"natsmon/modules/natsbiz"
	sctx "natsmon/service-context"
	"natsmon/service-context/component/natspub"
	"natsmon/service-context/core"
)

type AddMessageParams struct {
	Payload string `json:"payload" form:"payload" binding:"required"`
}

func AddMessage(sc sctx.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		stream := c.Param("stream")
		var data AddMessageParams

		if err := c.ShouldBind(&data); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		publisher := sc.MustGet(common.KeyNatsPubComp).(natspub.NatsPubComponent)

		biz := natsbiz.NewPublishMessageBiz(publisher)
		if err := biz.Response(c, stream, data.Payload); err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData("ok"))
	}
}
