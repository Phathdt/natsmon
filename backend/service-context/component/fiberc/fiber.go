package fiberc

import (
	"flag"

	"github.com/gofiber/fiber/v2"
	sctx "natsmon/service-context"
)

type FiberComponent interface {
	GetRouter() *fiber.App
	GetPort() int
}

type fiberComp struct {
	id     string
	logger sctx.Logger
	port   int
	app    *fiber.App
}

func NewFiberComp(id string) *fiberComp {
	return &fiberComp{id: id}
}

func (f *fiberComp) ID() string {
	return f.id
}

func (f *fiberComp) InitFlags() {
	flag.IntVar(&f.port, "fiber-port", 3000, "fiber port")
}

func (f *fiberComp) Activate(sc sctx.ServiceContext) error {
	f.logger = sc.Logger(f.id)
	f.app = fiber.New()

	return nil
}

func (f *fiberComp) Stop() error {
	return nil
}

func (f *fiberComp) GetRouter() *fiber.App {
	return f.app
}

func (f *fiberComp) GetPort() int {
	return f.port
}
