package services

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kliuchnikovv/engi-example/entity"
	"github.com/kliuchnikovv/engi-example/store"
	"gorm.io/gorm"
)

type TasksAPI struct {
	Store *store.TaskStore
}

func NewTasksAPI(store *store.TaskStore) *TasksAPI {
	return &TasksAPI{Store: store}
}

func (api *TasksAPI) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/tasks", api.Create).Methods("POST")
	r.HandleFunc("/tasks", api.List).Methods("GET")
	r.HandleFunc("/tasks/{id:[0-9]+}", api.Get).Methods("GET")
	r.HandleFunc("/tasks/{id:[0-9]+}", api.Update).Methods("PUT")
	r.HandleFunc("/tasks/{id:[0-9]+}", api.Delete).Methods("DELETE")
}

func (api *TasksAPI) Create(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  = r.Context()
		task entity.Task
	)

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if _, err := api.Store.GetByID(ctx, task.ID); err == nil {
		http.Error(w, "task already exists", http.StatusBadRequest)
		return
	}

	if err := api.Store.Create(ctx, &task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (api *TasksAPI) List(w http.ResponseWriter, r *http.Request) {
	var ctx = r.Context()

	tasks, err := api.Store.List(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (api *TasksAPI) Get(w http.ResponseWriter, r *http.Request) {
	var ctx = r.Context()

	id, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	task, err := api.Store.GetByID(ctx, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (api *TasksAPI) Update(w http.ResponseWriter, r *http.Request) {
	var ctx = r.Context()

	id, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var task entity.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	task.ID = id

	if _, err := api.Store.GetByID(ctx, id); errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, "task doesn't exist", http.StatusNotFound)
		return
	}

	if err := api.Store.Update(ctx, &task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (api *TasksAPI) Delete(w http.ResponseWriter, r *http.Request) {
	var ctx = r.Context()

	id, err := parseID(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if _, err := api.Store.GetByID(ctx, id); errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, "task doesn't exist", http.StatusNotFound)
		return
	}

	if err := api.Store.Delete(ctx, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func parseID(idStr string) (int64, error) {
	return strconv.ParseInt(idStr, 10, 64)
}
