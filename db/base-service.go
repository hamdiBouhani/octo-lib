package db

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseService[T IObject] struct {
	Database *gorm.DB
}

// Create a new object
func (s *BaseService[T]) Create(ctx context.Context, obj T) error {
	return s.Database.WithContext(ctx).Create(obj).Error
}

// Get by UUID
func (s *BaseService[T]) Get(ctx context.Context, id uuid.UUID) (T, error) {
	var obj T
	err := s.Database.WithContext(ctx).First(&obj, "id = ?", id).Error
	return obj, err
}

// List all (with soft delete filtering)
func (s *BaseService[T]) List(ctx context.Context) ([]T, error) {
	var objs []T
	err := s.Database.WithContext(ctx).Find(&objs).Error
	return objs, err
}

// Update object
func (s *BaseService[T]) Update(ctx context.Context, obj T) error {
	base := obj.GetBaseObject()
	base.MarkUpdate(s.Database)
	return s.Database.WithContext(ctx).Save(obj).Error
}

// Soft delete
func (s *BaseService[T]) Delete(ctx context.Context, id uuid.UUID) error {
	var obj T
	err := s.Database.WithContext(ctx).First(&obj, "id = ?", id).Error
	if err != nil {
		return err
	}
	return s.Database.WithContext(ctx).Delete(&obj).Error
}

// Hard delete (permanent)
func (s *BaseService[T]) HardDelete(ctx context.Context, id uuid.UUID) error {
	var obj T
	err := s.Database.WithContext(ctx).Unscoped().First(&obj, "id = ?", id).Error
	if err != nil {
		return err
	}
	return s.Database.WithContext(ctx).Unscoped().Delete(&obj).Error
}

// Exists check
func (s *BaseService[T]) Exists(ctx context.Context, id uuid.UUID) (bool, error) {
	var count int64
	err := s.Database.WithContext(ctx).
		Model(new(T)).
		Where("id = ?", id).
		Count(&count).Error

	return count > 0, err
}

// Pagination
func (s *BaseService[T]) ListPaginated(ctx context.Context, page, size int) ([]T, int64, error) {
	var objs []T
	var total int64

	s.Database.WithContext(ctx).Model(new(T)).Count(&total)

	err := s.Database.WithContext(ctx).
		Limit(size).
		Offset((page - 1) * size).
		Find(&objs).Error

	return objs, total, err
}

// Find by field
func (s *BaseService[T]) FindBy(ctx context.Context, field string, value any) ([]T, error) {
	var objs []T
	err := s.Database.WithContext(ctx).
		Where(field+" = ?", value).
		Find(&objs).Error

	return objs, err
}
