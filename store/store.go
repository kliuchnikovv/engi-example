package store

import (
	"context"

	"gorm.io/gorm"
)

type Store[T any] struct {
	db *gorm.DB
}

func NewStore[T any](db *gorm.DB) *Store[T] {
	return &Store[T]{db: db}
}

func (s *Store[T]) Create(
	ctx context.Context,
	model *T,
) error {
	return s.db.WithContext(ctx).Create(model).Error
}

func (s *Store[T]) GetByID(ctx context.Context, id int64) (*T, error) {
	var model T
	err := s.db.WithContext(ctx).First(&model, id).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (s *Store[T]) List(ctx context.Context) ([]*T, error) {
	var models []*T
	err := s.db.WithContext(ctx).Find(&models).Error
	return models, err
}

func (s *Store[T]) Update(ctx context.Context, model *T) error {
	return s.db.WithContext(ctx).Save(model).Error
}

func (s *Store[T]) Delete(ctx context.Context, id int64) error {
	var model T
	return s.db.WithContext(ctx).Delete(&model, id).Error
}
