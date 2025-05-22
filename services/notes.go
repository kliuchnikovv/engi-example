package services

import (
	"context"
	"errors"

	"github.com/kliuchnikovv/engi"
	"github.com/kliuchnikovv/engi-example/entity"
	"github.com/kliuchnikovv/engi-example/store"
	"github.com/kliuchnikovv/engi/definition/middlewares"
	"github.com/kliuchnikovv/engi/definition/middlewares/auth"
	"github.com/kliuchnikovv/engi/definition/middlewares/cors"
	"github.com/kliuchnikovv/engi/definition/parameter"
	"github.com/kliuchnikovv/engi/definition/parameter/path"
	"github.com/kliuchnikovv/engi/definition/parameter/placing"
	"github.com/kliuchnikovv/engi/definition/validate"
	"gorm.io/gorm"
)

// Example service.
type NotesAPI struct {
	notesStore store.NoteStore
}

func NewNotesAPI(notesStore store.NoteStore) *NotesAPI {
	return &NotesAPI{notesStore: notesStore}
}

func (api *NotesAPI) Prefix() string {
	// All requests to this service should start with "/notes".
	return "notes"
}

func (api *NotesAPI) Middlewares() []engi.Middleware {
	// Defines middlewares for all requests to this service.
	// CORS, auth, etc...
	return []engi.Middleware{
		cors.AllowedOrigins("*"),
		cors.AllowedHeaders("*"),
		cors.AllowedMethods("*"),
		auth.NoAuth(),
	}
}

func (api *NotesAPI) Routers() engi.Routes {
	return engi.Routes{
		engi.PST(""): engi.Handle( // Using POST method to create a new note with full address "/notes/".
			api.Create,                       // Handler to handle this request.
			parameter.Body(new(entity.Note)), // Body parameter to parse request body into entity.Note.
			middlewares.Description("create new note"), // Description of this route for documentation purposes.
		),
		engi.GET(""): engi.Handle(
			api.List,
			middlewares.Description("list all notes"),
		),
		engi.GET("{id}"): engi.Handle(
			api.Get,
			path.Integer("id", validate.Greater(0)), // Path parameter to be parsed into integer.
			middlewares.Description("get note by id"),
		),
		engi.PUT("{id}"): engi.Handle(
			api.Update,
			path.Integer("id", validate.Greater(0)),
			parameter.Body(new(entity.Note)),
			middlewares.Description("update note by id"),
		),
		engi.DEL("{id}"): engi.Handle(
			api.Delete,
			path.Integer("id", validate.Greater(0)),
			middlewares.Description("delete note by id"),
		),
	}
}

func (api *NotesAPI) Create(
	ctx context.Context,
	request engi.Request,
	response engi.Response,
) error {
	var note = request.Body().(*entity.Note)

	if _, err := api.notesStore.GetByID(ctx, note.ID); err == nil {
		return response.BadRequest("note already exists")
	}

	if err := api.notesStore.Create(ctx, note); err != nil {
		return response.InternalServerError(err.Error())
	}

	return response.Created()
}

func (api *NotesAPI) List(
	ctx context.Context,
	request engi.Request,
	response engi.Response,
) error {
	notes, err := api.notesStore.List(ctx)
	if err != nil {
		return response.InternalServerError(err.Error())
	}

	return response.OK(notes)
}

func (api *NotesAPI) Get(
	ctx context.Context,
	request engi.Request,
	response engi.Response,
) error {
	var id = request.Integer("id", placing.InPath)

	note, err := api.notesStore.GetByID(ctx, id)
	if err != nil {
		return response.NotFound(err.Error())
	}

	return response.OK(note)
}

func (api *NotesAPI) Update(
	ctx context.Context,
	request engi.Request,
	response engi.Response,
) error {
	var (
		id   = request.Integer("id", placing.InPath)
		note = request.Body().(*entity.Note)
	)

	if _, err := api.notesStore.GetByID(ctx, id); errors.Is(err, gorm.ErrRecordNotFound) {
		return response.NotFound("note doesn't exists")
	}

	note.ID = id

	if err := api.notesStore.Update(ctx, note); err != nil {
		return response.InternalServerError(err.Error())
	}

	return response.NoContent()
}

func (api *NotesAPI) Delete(
	ctx context.Context,
	request engi.Request,
	response engi.Response,
) error {
	var id = request.Integer("id", placing.InPath)

	if _, err := api.notesStore.GetByID(ctx, id); errors.Is(err, gorm.ErrRecordNotFound) {
		return response.NotFound("note doesn't exists")
	}

	if err := api.notesStore.Delete(ctx, id); err != nil {
		return response.NotFound(err.Error())
	}

	return response.NoContent()
}
