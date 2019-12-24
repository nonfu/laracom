package repo

import (
    "github.com/jinzhu/gorm"
    "github.com/nonfu/laracom/product-service/model"
)

type AttributeRepositoryInterface interface {
    CreateAttribute(attribute *model.Attribute) error
    UpdateAttribute(attribute *model.Attribute) error
    DeleteAttribute(attribute *model.Attribute) error
    CreateValue(value *model.AttributeValue) error
    UpdateValue(value *model.AttributeValue) error
    DeleteValue(value *model.AttributeValue) error
    CreateProductAttribute(attribute *model.ProductAttribute) error
    UpdateProductAttribute(attribute *model.ProductAttribute) error
    DeleteProductAttribute(attribute *model.ProductAttribute) error
    GetAttribute(id uint) (*model.Attribute, error)
    GetAttributes() ([]*model.Attribute, error)
    GetAttributeValue(id uint) (*model.AttributeValue, error)
    GetAttributeValues(attributeId uint) ([]*model.AttributeValue, error)
    GetProductAttribute(id uint) (*model.ProductAttribute, error)
    GetProductAttributes(productId uint) ([]*model.ProductAttribute, error)
}

type AttributeRepository struct {
    Db *gorm.DB
}

func (repo *AttributeRepository) CreateAttribute(attribute *model.Attribute) error {
    if err := repo.Db.Create(attribute).Error; err != nil {
        return err
    }
    return nil
}

func (repo *AttributeRepository) UpdateAttribute(attribute *model.Attribute) error {
    if err := repo.Db.Save(attribute).Error; err != nil {
        return err
    }
    return nil
}

func (repo *AttributeRepository) DeleteAttribute(attribute *model.Attribute) error {
    if err := repo.Db.Delete(attribute).Error; err != nil {
        return err
    }
    return nil
}

func (repo *AttributeRepository) CreateValue(value *model.AttributeValue) error {
    if err := repo.Db.Create(value).Error; err != nil {
        return err
    }
    return nil
}

func (repo *AttributeRepository) UpdateValue(value *model.AttributeValue) error {
    if err := repo.Db.Save(value).Error; err != nil {
        return err
    }
    return nil
}

func (repo *AttributeRepository) DeleteValue(value *model.AttributeValue) error {
    if err := repo.Db.Delete(value).Error; err != nil {
        return err
    }
    return nil
}

func (repo *AttributeRepository) CreateProductAttribute(attribute *model.ProductAttribute) error {
    if err := repo.Db.Create(attribute).Error; err != nil {
        return err
    }
    return nil
}

func (repo *AttributeRepository) UpdateProductAttribute(attribute *model.ProductAttribute) error {
    if err := repo.Db.Save(attribute).Error; err != nil {
        return err
    }
    return nil
}

func (repo *AttributeRepository) DeleteProductAttribute(attribute *model.ProductAttribute) error {
    if err := repo.Db.Delete(attribute).Error; err != nil {
        return err
    }
    return nil
}

func (repo *AttributeRepository) GetAttribute(id uint) (*model.Attribute, error) {
    attribute := &model.Attribute{}
    if err := repo.Db.First(attribute, id).Error; err != nil {
        return nil, err
    }
    return attribute, nil
}

func (repo *AttributeRepository) GetAttributes() ([]*model.Attribute, error) {
    var attributes []*model.Attribute
    if err := repo.Db.Find(&attributes).Error; err != nil {
        return nil, err
    }
    return attributes, nil
}

func (repo *AttributeRepository) GetAttributeValue(id uint) (*model.AttributeValue, error) {
    value := &model.AttributeValue{}
    if err := repo.Db.First(value, id).Error; err != nil {
        return nil, err
    }
    return value, nil
}

func (repo *AttributeRepository) GetAttributeValues(attributeId uint) ([]*model.AttributeValue, error) {
    var values []*model.AttributeValue
    // 获取指定属性下的所有属性值
    if err := repo.Db.Find(&values, "attribute_id = ?", attributeId).Error; err != nil {
        return nil, err
    }
    return values, nil
}

func (repo *AttributeRepository) GetProductAttribute(id uint) (*model.ProductAttribute, error) {
    attribute := &model.ProductAttribute{}
    if err := repo.Db.First(attribute, id).Error; err != nil {
        return nil, err
    }
    return attribute, nil
}

func (repo *AttributeRepository) GetProductAttributes(productId uint) ([]*model.ProductAttribute, error) {
    var productAttributes []*model.ProductAttribute
    // 获取嵌套的关联关系
    if err := repo.Db.Where("product_id = ?", productId).Preload("AttributeValues.Attribute").Find(&productAttributes).Error; err != nil {
        return nil, err
    }
    return productAttributes, nil
}
