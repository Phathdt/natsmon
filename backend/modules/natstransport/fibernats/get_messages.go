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

func GetMessages(sc sctx.ServiceContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		stream := c.Params("stream")
		offset, _ := c.ParamsInt("offset", 1)

		natsComponent := sc.MustGet(common.KeyNatsComp).(natsc.NatsComponent)
		manager := natsComponent.GetManager()
		js := natsComponent.GetJs()

		repo := natsrepo.NewRepo(manager, js)
		biz := natsbiz.NewGetMessageBiz(repo)
		messages, err := biz.Response(c.Context(), stream, int64(offset))
		if err != nil {
			panic(err)
		}

		return c.Status(http.StatusOK).JSON(core.SimpleSuccessResponse(messages))
	}
}
