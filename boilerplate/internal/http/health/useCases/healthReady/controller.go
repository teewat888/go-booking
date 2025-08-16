package healthready

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teewat888/go-booking/boilerplate/internal/dependencies"
)

type Controller struct {
	deps dependencies.Dependencies
}

func (ctrl *Controller) Ready(c *fiber.Ctx) error {
	return c.JSON("health ok")
}

func New(deps *dependencies.Dependencies) *Controller {
	return &Controller{
		deps: *deps,
	}
}
