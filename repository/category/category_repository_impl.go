package repository

import (
	"context"
	"pinjam-buku/helper"
	"pinjam-buku/model/domain"

	"gorm.io/gorm"
)

// membuat implementasi dari repository (kontrak sebelumnya) untuk data category

// membuat struct category implementation
type CategoryRepositoryImpl struct {
}

// mendefinisikan constructor
func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

// membuat method milik struct CategoryRepositoryImpl yang mana menerapkan kontrak sebelumnya
func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *gorm.DB, category domain.Category) domain.Category {

	// insert data category
	err := tx.WithContext(ctx).Create(&category).Error

	// cek error
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, category domain.Category) domain.Category {

	// update data category
	err := tx.WithContext(ctx).
		Model(&category).Updates(domain.Category{
		Name: category.Name,
	}).Error

	// cek error
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, category domain.Category) {

	// delete category
	err := tx.WithContext(ctx).Delete(&category).Error

	// cek error
	helper.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *gorm.DB, categoryId int) (domain.Category, error) {

	category := domain.Category{}

	// mencari category berdasarkan id
	err := tx.WithContext(ctx).
		First(&category, categoryId).Error

	// jika error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB) []domain.Category {

	var categories []domain.Category

	// mengambil semua data category
	err := tx.WithContext(ctx).
		Find(&categories).Error

	// cek error
	helper.PanicIfError(err)

	return categories
}
