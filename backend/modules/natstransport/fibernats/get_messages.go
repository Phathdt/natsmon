package fibernats

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"natsmon/common"
	"natsmon/modules/natsbiz"
	"natsmon/modules/natsrepo"
	sctx "natsmon/service-context"
	"natsmon/service-context/component/natsc"
	"natsmon/service-context/core"
)

type GetMessageParams struct {
	Offset int64 `json:"offset" query:"offset" validated:"required"`
}

func GetMessages(sc sctx.ServiceContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		stream := c.Params("stream")

		var data GetMessageParams

		if err := c.QueryParser(&data); err != nil {
			panic(err)
		}

		natsComponent := sc.MustGet(common.KeyNatsComp).(natsc.NatsComponent)
		manager := natsComponent.GetManager()
		js := natsComponent.GetJs()

		repo := natsrepo.NewRepo(manager, js)
		biz := natsbiz.NewGetMessageBiz(repo)
		messages, err := biz.Response(c.Context(), stream, data.Offset)
		if err != nil {
			panic(err)
		}

		return c.Status(http.StatusOK).JSON(core.SimpleSuccessResponse(messages))
	}
}
