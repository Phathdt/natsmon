package fibernats

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"natsmon/common"
	"natsmon/modules/natsrepo"
	sctx "natsmon/service-context"
	"natsmon/service-context/component/natsc"
	"natsmon/service-context/core"
)

type filter struct {
	Offset int64 `query:"offset"`
}

func GetMessages(sc sctx.ServiceContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		stream := c.Params("stream")
		var data filter
		if err := c.QueryParser(&data); err != nil {
			panic(err)
		}
		natsComponent := sc.MustGet(common.KeyNatsComp).(natsc.NatsComponent)
		manager := natsComponent.GetManager()
		js := natsComponent.GetJs()

		repo := natsrepo.NewRepo(manager, js)
		messages, err := repo.GetMessages(c.Context(), stream, data.Offset)
		if err != nil {
			panic(err)
		}

		return c.Status(http.StatusOK).JSON(core.SimpleSuccessResponse(messages))
	}
}
