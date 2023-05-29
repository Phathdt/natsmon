package fibernats

import (
	"github.com/gofiber/fiber/v2"
	"natsmon/common"
	"natsmon/modules/natsbiz"
	"natsmon/modules/natsrepo"
	sctx "natsmon/service-context"
	"natsmon/service-context/component/natsc"
)

func GetConsumer(sc sctx.ServiceContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		stream := c.Params("stream")

		natsComponent := sc.MustGet(common.KeyNatsComp).(natsc.NatsComponent)
		manager := natsComponent.GetManager()

		repo := natsrepo.NewRepo(manager)
		biz := natsbiz.NewGetConsumerBiz(repo)

		consumers, err := biz.Response(c.Context(), stream)
		if err != nil {
			panic(err)
		}

		return c.JSON(consumers)
	}
}
