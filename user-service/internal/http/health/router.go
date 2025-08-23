package health

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teewat888/user-service/internal/dependencies"
	healthready "github.com/teewat888/user-service/internal/http/health/useCases/healthReady"
)

type useCases struct {
	healthready *healthready.Controller
}

type HealthModule struct {
	deps   *dependencies.Dependencies
	router *fiber.App
	uc     useCases
}

func (hm *HealthModule) Configure() *HealthModule {
	hm.router.Get("/", hm.uc.healthready.Ready)
	return hm
}

func (hm *HealthModule) Router() *fiber.App {
	return hm.router
}

func New(deps *dependencies.Dependencies) *HealthModule {
	return &HealthModule{
		deps: deps,
		router: fiber.New(fiber.Config{
			AppName: deps.Config.ServiceId,
		}),
		uc: useCases{
			healthready: healthready.New(deps),
		},
	}
}
