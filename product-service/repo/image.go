package repo

import (
    "github.com/jinzhu/gorm"
    "github.com/nonfu/laracom/product-service/model"
)

type ImageRepositoryInterface interface {
    Create(image *model.ProductImage) error
    Update(image *model.ProductImage) error
    Delete(image *model.ProductImage) error
    GetById(id uint) (*model.ProductImage, error)
    GetByProductId(productId uint) ([]*model.ProductImage, error)
}

type ImageRepository struct {
    Db *gorm.DB
}

func (repo *ImageRepository) Create(image *model.ProductImage) error {
    if err := repo.Db.Create(image).Error; err != nil {
        return err
    }
    return nil
}

func (repo *ImageRepository) Update(image *model.ProductImage) error {
    if err := repo.Db.Save(image).Error; err != nil {
        return err
    }
    return nil
}

func (repo *ImageRepository) Delete(image *model.ProductImage) error {
    if err := repo.Db.Delete(image).Error; err != nil {
        return err
    }
    return nil
}

func (repo *ImageRepository) GetById(id uint) (*model.ProductImage, error) {
    image := &model.ProductImage{}
    if err := repo.Db.First(&image, id).Error; err != nil {
        return nil, err
    }
    return image, nil
}

func (repo *ImageRepository) GetByProductId(productId uint) ([]*model.ProductImage, error) {
    var images []*model.ProductImage
    // 获取指定商品的所有图片
    if err := repo.Db.Find(&images, "product_id = ?", productId).Error; err != nil {
        return nil, err
    }
    return images, nil
}
