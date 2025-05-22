package main

import (
	"log"
	"log/slog"
	"os"
	"sync"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/kliuchnikovv/engi"
	"github.com/kliuchnikovv/engi-example/entity"
	"github.com/kliuchnikovv/engi-example/services"
	"github.com/kliuchnikovv/engi-example/store"
	"github.com/kliuchnikovv/engi/definition/response" // TODO:
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(&entity.Task{}, &entity.Note{}); err != nil {
		log.Fatal(err)
	}

	var waitGroup = &sync.WaitGroup{}

	go runServer(db, waitGroup, newEngi)
	go runServer(db, waitGroup, newGorillaMux)

	waitGroup.Wait()
}

func newEngi(db *gorm.DB) error {
	noteStore := store.NewNoteStore(db)

	var engine = engi.New(":8080",
		engi.ResponseAsJSON(response.AsIs),
		engi.WithLogger(slog.NewTextHandler(os.Stdout,
			&slog.HandlerOptions{
				Level: slog.LevelDebug,
			},
		)),
	)

	if err := engine.RegisterServices(
		services.NewNotesAPI(*noteStore),
	); err != nil {
		return err
	}

	return engine.Start()
}

func newGorillaMux(db *gorm.DB) error {
	// Initialize stores
	taskStore := store.NewTaskStore(db)

	// Initialize APIs
	tasksAPI := services.NewTasksAPI(taskStore)

	// Create router and register routes
	r := mux.NewRouter()
	tasksAPI.RegisterRoutes(r)

	// Start HTTP server
	return http.ListenAndServe(":8081", r)
}

func runServer(db *gorm.DB, wg *sync.WaitGroup, f func(db *gorm.DB) error) {
	wg.Add(1)
	defer wg.Done()

	if err := f(db); err != nil {
		log.Fatal(err)
	}
}
