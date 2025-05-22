package store

import (
	"github.com/kliuchnikovv/engi-example/entity"
	"gorm.io/gorm"
)

type NoteStore struct {
	*Store[entity.Note]
}

func NewNoteStore(db *gorm.DB) *NoteStore {
	return &NoteStore{
		Store: NewStore[entity.Note](db),
	}
}

type TaskStore struct {
	*Store[entity.Task]
}

func NewTaskStore(db *gorm.DB) *TaskStore {
	return &TaskStore{
		Store: NewStore[entity.Task](db),
	}
}
