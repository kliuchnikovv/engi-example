package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/KlyuchnikovV/engi"
	"github.com/KlyuchnikovV/engi-example/services"
	"github.com/KlyuchnikovV/engi/definition/response" // TODO:
)

func main() {
	var engine = engi.New(":8080",
		engi.WithPrefix("api"),
		engi.ResponseAsJSON(new(response.AsIs)),
		engi.WithLogger(slog.NewTextHandler(os.Stdout,
			&slog.HandlerOptions{
				Level: slog.LevelDebug,
			},
		)),
	)

	if err := engine.RegisterServices(
		new(services.NotesAPI),
		new(services.RequestAPI),
	); err != nil {
		log.Fatal(err)
	}

	if err := engine.Start(); err != nil {
		log.Fatal(err)
	}
}
