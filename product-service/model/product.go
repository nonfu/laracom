package model

import (
    "github.com/jinzhu/gorm"
    pb "github.com/nonfu/laracom/product-service/proto/product"
)

type Product struct {
    gorm.Model
    BrandId uint32 `gorm:"unsigned,default:0"`
    Sku string `gorm:"type:varchar(255)"`
    Name string `gorm:"type:varchar(255)"`
    Slug string `gorm:"type:varchar(255)"`
    Description string `gorm:"type:text"`
    Cover string `gorm:"type:varchar(255)"`
    Quantity uint32 `gorm:"unsigned,default:0"`
    Price float32 `gorm:"type:decimal(8,2)"`
    SalePrice float32 `gorm:"type:decimal(8,2)"`
    Status uint8 `gorm:"unsigned,default:0"`
    Length float32 `gorm:"type:decimal(8,2)"`
    Width float32 `gorm:"type:decimal(8,2)"`
    Height float32 `gorm:"type:decimal(8,2)"`
    Weight float32 `gorm:"type:decimal(8,2)"`
    DistanceUnit string `gorm:"type:varchar(255)"`
    MassUnit string `gorm:"type:varchar(255)"`
}

func (model *Product) ToORM(req *pb.Product) (*Product, error) {
    if req.Id != 0 {
        model.ID = uint(req.Id)
    }
    if req.BrandId != 0 {
        model.BrandId = req.BrandId
    }
    if req.Sku != "" {
        model.Sku = req.Sku
    }
    if req.Name != "" {
        model.Name = req.Name
    }
    if req.Slug != "" {
        model.Slug = req.Slug
    }
    if req.Description != "" {
        model.Description = req.Description
    }
    if req.Cover != "" {
        model.Cover = req.Cover
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
    if req.Status != 0 {
        model.Status = uint8(req.Status)
    }
    if req.Length != 0 {
        model.Length = req.Length
    }
    if req.Width != 0 {
        model.Width = req.Width
    }
    if req.Height != 0 {
        model.Height = req.Height
    }
    if req.Weight != 0 {
        model.Weight = req.Weight
    }
    if req.DistanceUnit != "" {
        model.DistanceUnit = req.DistanceUnit
    }
    if req.MassUnit != "" {
        model.MassUnit = req.MassUnit
    }
    return model, nil
}

func (model *Product) ToProtobuf() (*pb.Product, error) {
    var product = &pb.Product{}
    product.Id = uint32(model.ID)
    product.BrandId = model.BrandId
    product.Sku = model.Sku
    product.Name = model.Name
    product.Slug = model.Slug
    product.Description = model.Description
    product.Cover = model.Cover
    product.Quantity = model.Quantity
    product.Price = model.Price
    product.SalePrice = model.SalePrice
    product.Status = uint32(model.Status)
    product.Length = model.Length
    product.Width = model.Width
    product.Height = model.Height
    product.Weight = model.Weight
    product.DistanceUnit = model.DistanceUnit
    product.MassUnit = model.MassUnit
    product.CreatedAt = model.CreatedAt.Format("2006-01-02 15:04:05")
    product.UpdatedAt = model.UpdatedAt.Format("2006-01-02 15:04:05")
    return product, nil
}
