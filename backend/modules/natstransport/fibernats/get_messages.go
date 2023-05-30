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

type GetMessageParams struct {
	Offset int64 `json:"offset" form:"offset" binding:"required"`
}

func GetMessages(sc sctx.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		stream := c.Param("stream")

		var data GetMessageParams

		if err := c.ShouldBind(&data); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		natsComponent := sc.MustGet(common.KeyNatsComp).(natsc.NatsComponent)
		manager := natsComponent.GetManager()
		js := natsComponent.GetJs()

		repo := natsrepo.NewRepo(manager, js)
		biz := natsbiz.NewGetMessageBiz(repo)
		messages, err := biz.Response(c, stream, data.Offset)
		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(messages))
	}
}
