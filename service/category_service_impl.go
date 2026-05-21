package service

import (
	"context"
	"pinjam-buku/exception"
	"pinjam-buku/helper"
	"pinjam-buku/model/domain"
	"pinjam-buku/model/web/category"
	repository "pinjam-buku/repository/category"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CategoryServiceImpl struct {
	// di category service implementation membutuhkan repository, maka perlu di tambahkan di atribut
	CategoryRepository repository.CategoryRepository

	// butuh koneksi database juga, maka tambahkan attribute sql
	// db bentuknya adalah struct bukan interface, maka di set sebagai pointer
	DB *gorm.DB

	// menambahkan attribute validate
	// validate ditambahkan di function create dan update
	// di function yang lain tidak dibutuhkan validasi, karena tidak ada payload (data asli) di parameter nya
	Validate *validator.Validate
}

// mendefinisikan construktor dari controller

// membuat untuk router dengan return CategoryService (interface)
func NewCategoryService(categoryRepository repository.CategoryRepository, DB *gorm.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

// mengimplementasi category service (membuat method yang ada di category  agar dimiliki oleh CategoryServiceImpl)
// karena database yang kita gunakan adalah transactional (mysql), maka requestnya nanti dalam bentuk transactonal

func (service CategoryServiceImpl) Create(ctx context.Context, request category.CategoryCreateRequest) (category.CategoryResponse, error) {
	// mengimplementasi validation sebelum melakukan transaction
	err := service.Validate.Struct(request)

	// mengecek error
	if err != nil {
		return category.CategoryResponse{}, err
	}

	// memulai koneksi database transactional
	tx := service.DB.Begin()

	// cek transaction error
	if tx.Error != nil {
		return category.CategoryResponse{}, tx.Error
	}

	// mencegah ketika terjadi error setelah proses selesai semuanya dengan defer (pencegahan terakhir jika terjadi error)
	defer func() {
		helper.CommitOrRollback(tx, err)
	}()

	// category entity
	categoryDomain := domain.Category{
		Name: request.Name,
	}

	// save category
	categoryDomain, err = service.CategoryRepository.Save(
		ctx,
		tx,
		categoryDomain,
	)

	// check error save
	if err != nil {
		return category.CategoryResponse{}, err
	}

	// return response
	return helper.ToCategoryResponse(categoryDomain), nil
}

func (service CategoryServiceImpl) Update(ctx context.Context, request category.CategoryUpdateRequest) (category.CategoryResponse, error) {
	// mengimplementasi validation sebelum melakukan transaction
	err := service.Validate.Struct(request)

	// mengecek transaction error
	if err != nil {
		return category.CategoryResponse{}, err
	}

	// memulai koneksi database transactional
	tx := service.DB.Begin()

	// cek transaction error
	if tx.Error != nil {
		return category.CategoryResponse{}, tx.Error
	}

	// mencegah ketika terjadi error setelah proses selesai semuanya dengan defer (pencegahan terakhir jika terjadi error)
	defer func() {
		helper.CommitOrRollback(tx, err)
	}()

	// melakukan pencarian data category terlebih dahulu sebelum dilakukan  dengan function FindById
	categoryDomain, err := service.CategoryRepository.FindById(ctx, tx, request.Id)

	// jika data tidak ditemukan
	if err != nil {
		return category.CategoryResponse{},
			exception.NewNotFoundError(err.Error())
	}

	// assign data baru
	categoryDomain.Name = request.Name

	// save update
	categoryDomain, err = service.CategoryRepository.Update(
		ctx,
		tx,
		categoryDomain,
	)

	if err != nil {
		return category.CategoryResponse{}, err
	}

	// return response
	return helper.ToCategoryResponse(categoryDomain), nil
}

func (service CategoryServiceImpl) Delete(ctx context.Context, categoryId string) (err error) {
	// memulai koneksi database transactional
	tx := service.DB.Begin()

	// cek transaction error
	if tx.Error != nil {
		return tx.Error
	}

	// mencegah ketika terjadi error setelah proses selesai semuanya dengan defer (pencegahan terakhir jika terjadi error)
	defer func() {
		helper.CommitOrRollback(tx, err)
	}()

	// melakukan pencarian data category terlebih dahulu sebelum dilakukan  dengan function FindById
	categoryDomain, err := service.CategoryRepository.FindById(ctx, tx, categoryId)

	// melakukan pengecekan ketika data category tidak ditemukan
	if err != nil {
		return exception.NewNotFoundError(err.Error())
	}

	// delete category
	err = service.CategoryRepository.Delete(
		ctx,
		tx,
		categoryDomain,
	)

	if err != nil {
		return err
	}

	return nil
}

func (service CategoryServiceImpl) FindById(ctx context.Context, categoryId string) (response category.CategoryResponse, err error) {
	// memulai koneksi database transactional
	tx := service.DB.Begin()

	// cek transaction error
	if tx.Error != nil {
		return response, tx.Error
	}

	// mencegah ketika terjadi error setelah proses selesai semuanya dengan defer (pencegahan terakhir jika terjadi error)
	defer func() {
		helper.CommitOrRollback(tx, err)
	}()

	// melakukan pencarian data category terlebih dahulu sebelum dilakukan  dengan function FindById
	categoryDomain, err := service.CategoryRepository.FindById(ctx, tx, categoryId)

	// melakukan pengecekan ketika data category tidak ditemukan
	if err != nil {
		return response,
			exception.NewNotFoundError(err.Error())
	}

	// mengkonversi data category dari object menjadi response
	response = helper.ToCategoryResponse(categoryDomain)

	return response, nil
}

func (service CategoryServiceImpl) FindAll(ctx context.Context) (response []category.CategoryResponse, err error) {
	// memulai koneksi database transactional
	tx := service.DB.Begin()

	// cek transaction error
	if tx.Error != nil {
		return response, tx.Error
	}

	// mencegah ketika terjadi error setelah proses selesai semuanya dengan defer (pencegahan terakhir jika terjadi error)
	defer func() {
		helper.CommitOrRollback(tx, err)
	}()

	// melakukan pencarian data category terlebih dahulu sebelum dilakukan  dengan function FindById
	categories, err := service.CategoryRepository.FindAll(ctx, tx)

	if err != nil {
		return response, err
	}

	// melakkan return dan konversi ke bentuk slice category responses
	return helper.ToCategoryResponses(categories), nil
}
