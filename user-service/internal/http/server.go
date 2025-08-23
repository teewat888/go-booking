package httpServer

import (
	"net"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/teewat888/user-service/internal/dependencies"
	"github.com/teewat888/user-service/internal/http/health"
)

type HTTPServer struct {
	addr  string
	fiber *fiber.App
	hm    *health.HealthModule
}

func (s *HTTPServer) Start() error {
	return s.fiber.Listen(s.addr)
}

func (s *HTTPServer) Close() error {
	return s.fiber.Shutdown()
}

func (s *HTTPServer) Router() *fiber.App {
	return s.fiber
}

func (s *HTTPServer) Configure() *HTTPServer {
	s.fiber.Use(recover.New()).Use(logger.New())
	s.fiber.Mount("/health", s.hm.Router())
	return s
}

func New(deps *dependencies.Dependencies) *HTTPServer {
	return &HTTPServer{
		addr: net.JoinHostPort(deps.Config.HTTP.Host, deps.Config.HTTP.Port),
		fiber: fiber.New(fiber.Config{
			AppName: deps.Config.ServiceId,
		}),
		hm: health.New(deps).Configure(),
	}
}
