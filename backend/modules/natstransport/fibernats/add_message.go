package fibernats

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"natsmon/common"
	"natsmon/modules/natsbiz"
	sctx "natsmon/service-context"
	"natsmon/service-context/component/natspub"
	"natsmon/service-context/core"
)

type AddMessageParams struct {
	Payload string `json:"payload" form:"payload" validate:"required"`
}

func AddMessage(sc sctx.ServiceContext) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		stream := ctx.Params("stream")
		var data AddMessageParams

		if err := ctx.BodyParser(&data); err != nil {
			panic(err)
		}

		publisher := sc.MustGet(common.KeyNatsPubComp).(natspub.NatsPubComponent)

		biz := natsbiz.NewPublishMessageBiz(publisher)
		if err := biz.Response(ctx.Context(), stream, data.Payload); err != nil {
			panic(err)
		}

		return ctx.Status(http.StatusOK).JSON(core.SimpleSuccessResponse("ok"))
	}
}
