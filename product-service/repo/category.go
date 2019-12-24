package repo

import (
    "github.com/jinzhu/gorm"
    "github.com/nonfu/laracom/product-service/model"
)

type CategoryRepositoryInterface interface {
    Create(category *model.Category) error
    Update(category *model.Category) error
    Delete(category *model.Category) error
    GetById(id uint) (*model.Category, error)
    GetAll() ([]*model.Category, error)
    GetWithProducts(categoryId uint) (*model.Category, error)
}

type CategoryRepository struct {
    Db *gorm.DB
}

func (repo *CategoryRepository) Create(category *model.Category) error {
    if err := repo.Db.Create(category).Error; err != nil {
        return err
    }
    return nil
}

func (repo *CategoryRepository) Update(category *model.Category) error {
    if err := repo.Db.Save(category).Error; err != nil {
        return err
    }
    return nil
}

func (repo *CategoryRepository) Delete(category *model.Category) error {
    if err := repo.Db.Delete(category).Error; err != nil {
        return err
    }
    return nil
}

func (repo *CategoryRepository) GetById(id uint) (*model.Category, error) {
    category := &model.Category{}
    if err := repo.Db.First(category, id).Error; err != nil {
        return nil, err
    }
    return category, nil
}

func (repo *CategoryRepository) GetAll() ([]*model.Category, error) {
    var categories []*model.Category
    if err := repo.Db.Find(&categories).Error; err != nil {
        return nil, err
    }
    return categories, nil
}

func (repo *CategoryRepository) GetWithProducts(categoryId uint) (*model.Category, error) {
    category := &model.Category{}
    // 获取与之关联的所有商品
    if err := repo.Db.Where("id = ?", categoryId).Preload("Products").First(category).Error; err != nil {
        return nil, err
    }
    return category, nil
}
