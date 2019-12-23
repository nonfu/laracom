package model

import (
    "github.com/jinzhu/gorm"
    pb "github.com/nonfu/laracom/product-service/proto/product"
)

type Attribute struct {
    gorm.Model
    Name string `gorm:"type:varchar(255);unique_index"`
    Values []*AttributeValue
}

func (model *Attribute) ToORM(req *pb.Attribute) (*Attribute, error) {
    if req.Id != 0 {
        model.ID = uint(req.Id)
    }
    if req.Name != "" {
        model.Name = req.Name
    }
    return model, nil
}

func (model *Attribute) ToProtobuf() (*pb.Attribute, error) {
    var attribute = &pb.Attribute{}
    attribute.Id = uint32(model.ID)
    attribute.Name = model.Name
    attribute.CreatedAt = model.CreatedAt.Format("2006-01-02 15:04:05")
    attribute.UpdatedAt = model.UpdatedAt.Format("2006-01-02 15:04:05")
    if model.Values != nil {
        attributeValues := make([]*pb.AttributeValue, len(model.Values))
        for index, value := range model.Values {
            attribute, _ := value.ToProtobuf()
            attributeValues[index] = attribute
        }
        attribute.Values = attributeValues
    }
    return attribute, nil
}

type AttributeValue struct {
    gorm.Model
    Value string `gorm:"type:varchar(255)"`
    AttributeId uint `gorm:"undefined,default:0;index"`
    Attribute Attribute
    ProductAttributes []*ProductAttribute `gorm:"many2many:attribute_value_product_attribute"`
}

func (model *AttributeValue) ToORM(req *pb.AttributeValue) (*AttributeValue, error) {
    if req.Id != 0 {
        model.ID = uint(req.Id)
    }
    if req.Value != "" {
        model.Value = req.Value
    }
    if req.AttributeId != 0 {
        model.AttributeId = uint(req.AttributeId)
    }
    return model, nil
}

func (model *AttributeValue) ToProtobuf() (*pb.AttributeValue, error) {
    var attributeValue = &pb.AttributeValue{}
    attributeValue.Id = uint32(model.ID)
    attributeValue.Value = model.Value
    attributeValue.AttributeId = uint32(model.AttributeId)
    attributeValue.CreatedAt = model.CreatedAt.Format("2006-01-02 15:04:05")
    attributeValue.UpdatedAt = model.UpdatedAt.Format("2006-01-02 15:04:05")
    if model.Attribute.ID != 0 {
        attributeValue.Attribute, _ = model.Attribute.ToProtobuf()
    }
    if model.ProductAttributes != nil {
        attributes := make([]*pb.ProductAttribute, len(model.ProductAttributes))
        for index, value := range model.ProductAttributes {
            attribute, _ := value.ToProtobuf()
            attributes[index] = attribute
        }
        attributeValue.ProductAttributes = attributes
    }
    return attributeValue, nil
}

type ProductAttribute struct {
    gorm.Model
    ProductId uint `gorm:"undefined,default:0;index"`
    Quantity uint32 `gorm:"unsigned,default:0"`
    Price float32 `gorm:"type:decimal(8,2)"`
    SalePrice float32 `gorm:"type:decimal(8,2)"`
    Default uint8 `gorm:"unsigned,default:0"`
    Product Product
    AttributeValues []*AttributeValue `gorm:"many2many:attribute_value_product_attribute"`
}

func (model *ProductAttribute) ToORM(req *pb.ProductAttribute) (*ProductAttribute, error) {
    if req.Id != 0 {
        model.ID = uint(req.Id)
    }
    if req.ProductId != 0 {
        model.ProductId = uint(req.ProductId)
    }
    if req.Quantity != 0 {
        model.Quantity = req.Quantity
    }
    if req.Price != 0 {
        model.Price = req.Price
    }
    if req.SalePrice != 0 {
        model.SalePrice = req.SalePrice
    }
    if req.Default != 0 {
        model.Default = uint8(req.Default)
    }
    return model, nil
}

func (model *ProductAttribute) ToProtobuf() (*pb.ProductAttribute, error) {
    var attribute = &pb.ProductAttribute{}
    attribute.Id = uint32(model.ID)
    attribute.ProductId = uint32(model.ProductId)
    attribute.Quantity = model.Quantity
    attribute.Price = model.Price
    attribute.SalePrice = model.SalePrice
    attribute.Default = uint32(model.Default)
    attribute.CreatedAt = model.CreatedAt.Format("2006-01-02 15:04:05")
    attribute.UpdatedAt = model.UpdatedAt.Format("2006-01-02 15:04:05")
    if model.AttributeValues != nil {
        attributes := make([]*pb.AttributeValue, len(model.AttributeValues))
        for index, value := range model.AttributeValues {
            attribute, _ := value.ToProtobuf()
            attributes[index] = attribute
        }
        attribute.AttributeValues = attributes
    }
    return attribute, nil
}