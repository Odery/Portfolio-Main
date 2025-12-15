package internal

import "github.com/gofiber/fiber/v2"

type Server struct {
	App  *fiber.App
	Port string
}

func NewServer(port, name string) *Server {
	app := fiber.New(fiber.Config{
		AppName: name,
	})

	//Server static files
	app.Static("/", "./static")

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
