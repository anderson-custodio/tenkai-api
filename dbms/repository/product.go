package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/softplan/tenkai-api/dbms/model"
)

//ProductDAOInterface ProductDAOInterface
type ProductDAOInterface interface {
	CreateProduct(e model.Product) (int, error)
	CreateProductVersion(e model.ProductVersion) (int, error)
	CreateProductVersionService(e model.ProductVersionService) (int, error)
	EditProduct(e model.Product) error
	EditProductVersion(e model.ProductVersion) error
	EditProductVersionService(e model.ProductVersionService) error
	DeleteProduct(id int) error
	DeleteProductVersion(id int) error
	DeleteProductVersionService(id int) error
	ListProducts() ([]model.Product, error)
	ListProductsVersions(id int) ([]model.ProductVersion, error)
	ListProductsVersionServices(id int) ([]model.ProductVersionService, error)
	ListProductVersionServicesLatest(productID, productVersionID int) ([]model.ProductVersionService, error)
	CreateProductVersionCopying(payload model.ProductVersion) (int, error)
}

//ProductDAOImpl ProductDAOImpl
type ProductDAOImpl struct {
	Db *gorm.DB
}

// CreateProductVersionCopying create a new version product version
func (dao ProductDAOImpl) CreateProductVersionCopying(payload model.ProductVersion) (int, error) {
	id, err := dao.CreateProductVersion(payload)
	if err != nil {
		return -1, err
	}

	if payload.CopyLatestRelease {
		list, err := dao.ListProductVersionServicesLatest(payload.ProductID, id)
		if err != nil {
			return -1, err
		}
		var pvs *model.ProductVersionService
		for _, l := range list {
			pvs = &model.ProductVersionService{}
			pvs.ProductVersionID = id
			pvs.ServiceName = l.ServiceName
			pvs.DockerImageTag = l.DockerImageTag

			if _, err := dao.CreateProductVersionService(*pvs); err != nil {
				return -1, err
			}
		}
	}
	return id, nil
}

//CreateProduct - Create a new product
func (dao ProductDAOImpl) CreateProduct(e model.Product) (int, error) {
	if err := dao.Db.Create(&e).Error; err != nil {
		return -1, err
	}
	return int(e.ID), nil
}

//CreateProductVersion - Create a new product version
func (dao ProductDAOImpl) CreateProductVersion(e model.ProductVersion) (int, error) {
	if err := dao.Db.Create(&e).Error; err != nil {
		return -1, err
	}
	return int(e.ID), nil
}

//CreateProductVersionService - Create a new product version
func (dao ProductDAOImpl) CreateProductVersionService(e model.ProductVersionService) (int, error) {
	if err := dao.Db.Create(&e).Error; err != nil {
		return -1, err
	}
	return int(e.ID), nil
}

//EditProduct - Updates an existing product
func (dao ProductDAOImpl) EditProduct(e model.Product) error {
	if err := dao.Db.Save(&e).Error; err != nil {
		return err
	}
	return nil
}

//EditProductVersion - Updates an existing product version
func (dao ProductDAOImpl) EditProductVersion(e model.ProductVersion) error {
	if err := dao.Db.Save(&e).Error; err != nil {
		return err
	}
	return nil
}

//EditProductVersionService - Updates an existing product version
func (dao ProductDAOImpl) EditProductVersionService(e model.ProductVersionService) error {
	if err := dao.Db.Save(&e).Error; err != nil {
		return err
	}
	return nil
}

//DeleteProduct - Deletes a product
func (dao ProductDAOImpl) DeleteProduct(id int) error {
	if err := dao.Db.Unscoped().Delete(model.Product{}, id).Error; err != nil {
		return err
	}
	return nil
}

//DeleteProductVersion - Deletes a productVersion
func (dao ProductDAOImpl) DeleteProductVersion(id int) error {
	if err := dao.Db.Unscoped().Delete(model.ProductVersion{}, id).Error; err != nil {
		return err
	}
	return nil
}

//DeleteProductVersionService - Deletes a productVersionService
func (dao ProductDAOImpl) DeleteProductVersionService(id int) error {
	if err := dao.Db.Unscoped().Delete(model.ProductVersionService{}, id).Error; err != nil {
		return err
	}
	return nil
}

//ListProducts - List products
func (dao ProductDAOImpl) ListProducts() ([]model.Product, error) {
	list := make([]model.Product, 0)
	if err := dao.Db.Find(&list).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return make([]model.Product, 0), nil
		}
		return nil, err
	}
	return list, nil
}

//ListProductsVersions - List products versions
func (dao ProductDAOImpl) ListProductsVersions(id int) ([]model.ProductVersion, error) {
	list := make([]model.ProductVersion, 0)
	if err := dao.Db.Where(&model.ProductVersion{ProductID: id}).Find(&list).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return make([]model.ProductVersion, 0), nil
		}
		return nil, err
	}
	return list, nil
}

//ListProductsVersionServices - List products versions
func (dao ProductDAOImpl) ListProductsVersionServices(id int) ([]model.ProductVersionService, error) {
	list := make([]model.ProductVersionService, 0)
	if err := dao.Db.Where(&model.ProductVersionService{ProductVersionID: id}).Find(&list).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return make([]model.ProductVersionService, 0), nil
		}
		return nil, err
	}
	return list, nil
}

//ListProductVersionServicesLatest - List from the latest Product Version
func (dao ProductDAOImpl) ListProductVersionServicesLatest(productID, productVersionID int) ([]model.ProductVersionService, error) {
	item := model.ProductVersion{}
	list := make([]model.ProductVersionService, 0)

	if err := dao.Db.Where(&model.ProductVersion{ProductID: productID}).Not("id", productVersionID).Order("created_at desc").Limit(1).Find(&item).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return make([]model.ProductVersionService, 0), nil
		}
		return list, err
	}

	list, err := dao.ListProductsVersionServices(int(item.ID))
	if err != nil {
		return list, err
	}

	return list, nil
}
