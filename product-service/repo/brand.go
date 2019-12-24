package repo

import (
    "github.com/jinzhu/gorm"
    "github.com/nonfu/laracom/product-service/model"
)

type BrandRepositoryInterface interface {
    Create(brand *model.Brand) error
    Update(brand *model.Brand) error
    Delete(brand *model.Brand) error
    GetById(id uint) (*model.Brand, error)
    GetAll() ([]*model.Brand, error)
    GetWithProducts(brandId uint) (*model.Brand, error)
}

type BrandRepository struct {
    Db *gorm.DB
}

func (repo *BrandRepository) Create(brand *model.Brand) error {
    if err := repo.Db.Create(brand).Error; err != nil {
        return err
    }
    return nil
}

func (repo *BrandRepository) Update(brand *model.Brand) error {
    if err := repo.Db.Save(brand).Error; err != nil {
        return err
    }
    return nil
}

func (repo *BrandRepository) Delete(brand *model.Brand) error {
    if err := repo.Db.Delete(brand).Error; err != nil {
        return err
    }
    return nil
}

func (repo *BrandRepository) GetById(id uint) (*model.Brand, error) {
    brand := &model.Brand{}
    if err := repo.Db.First(brand, id).Error; err != nil {
        return nil, err
    }
    return brand, nil
}

func (repo *BrandRepository) GetAll() ([]*model.Brand, error) {
    var brands []*model.Brand
    if err := repo.Db.Find(&brands).Error; err != nil {
        return nil, err
    }
    return brands, nil
}

func (repo *BrandRepository) GetWithProducts(brandId uint) (*model.Brand, error) {
    brand  := &model.Brand{}
    // 获取与之关联的所有商品
    if err := repo.Db.Where("id = ?", brandId).Preload("Products").First(brand).Error; err != nil {
        return nil, err
    }
    return brand, nil
}