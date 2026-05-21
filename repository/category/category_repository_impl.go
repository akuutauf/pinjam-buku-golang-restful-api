package repository

import (
	"context"
	"errors"
	"pinjam-buku/model/domain"

	"gorm.io/gorm"
)

// membuat implementasi dari repository (kontrak sebelumnya) untuk data category

// membuat struct category implementation
type CategoryRepositoryImpl struct {
	//
}

// mendefinisikan constructor
func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

// membuat method milik struct CategoryRepositoryImpl yang mana menerapkan kontrak sebelumnya
func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *gorm.DB, category domain.Category) (domain.Category, error) {

	// insert data category
	err := tx.WithContext(ctx).Create(&category).Error

	// mengecek error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, category domain.Category) (domain.Category, error) {

	// update data category
	err := tx.WithContext(ctx).
		Model(&category).Updates(domain.Category{
		Name: category.Name,
	}).Error

	// mengecek error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, category domain.Category) error {

	// delete category
	err := tx.WithContext(ctx).Delete(&category).Error

	// mengecek error
	if err != nil {
		return err
	}

	return nil
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *gorm.DB, categoryId string) (domain.Category, error) {
	// menyiapkan data category
	var category domain.Category

	result := tx.WithContext(ctx).
		Where("id = ?", categoryId).
		First(&category)

	// jika error database
	if result.Error != nil {
		return category, result.Error
	}

	// jika data tidak ditemukan
	if result.RowsAffected == 0 {
		return category, errors.New("category not found")
	}

	return category, nil
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB) ([]domain.Category, error) {

	var categories []domain.Category

	// mengambil semua data category
	err := tx.WithContext(ctx).
		Find(&categories).Error

	// mengecek error
	if err != nil {
		return categories, err
	}

	return categories, nil
}
