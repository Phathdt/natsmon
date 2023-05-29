package fibernats

import (
	"github.com/gofiber/fiber/v2"
	"natsmon/common"
	"natsmon/modules/natsbiz"
	"natsmon/modules/natsrepo"
	sctx "natsmon/service-context"
	"natsmon/service-context/component/natsc"
)

func ListJetstream(sc sctx.ServiceContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		natsComponent := sc.MustGet(common.KeyNatsComp).(natsc.NatsComponent)
		manager := natsComponent.GetManager()

		repo := natsrepo.NewRepo(manager)
		biz := natsbiz.NewListJetstreamBiz(repo)
		rs, err := biz.Response(c.Context(), nil)
		if err != nil {
			panic(err)
		}

		return c.JSON(rs)
	}
}
