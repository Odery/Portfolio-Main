package internal

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

type Server struct {
	App  *fiber.App
	Port string
}

func NewServer(port, name string, staticFS embed.FS) *Server {
	app := fiber.New(fiber.Config{
		AppName: name,
	})

	//Server static files
	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(staticFS),
		PathPrefix: "static",
		Browse:     false,
	}))

	//Handle routes
	app.Get("/", IndexPage)
	app.Get("/404", NotFoundPage)

	return &Server{
		App:  app,
		Port: port,
	}
}
func (s *Server) Start() error {
	return s.App.Listen(":" + s.Port)
}

func (s *Server) Stop() error {
	return s.App.Shutdown()
}
