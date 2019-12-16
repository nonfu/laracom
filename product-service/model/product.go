package model

import (
    "github.com/jinzhu/gorm"
    pb "github.com/nonfu/laracom/product-service/proto/product"
    "strconv"
)

type User struct {
    gorm.Model
    BrandId int `gorm:"unsigned,default:0"`
    Sku string `gorm:"type:varchar(255)"`
    Name string `gorm:"type:varchar(255)"`
    Slug string `gorm:"type:varchar(255)"`
    Description string `gorm:"type:text"`
    Cover string `gorm:"type:varchar(255)"`
    Quantity int
    Price float32 `gorm:"type:decimal(8,2)"`
    SalePrice float32 `gorm:"type:decimal(8,2)"`
    Status uint8 `gorm:"default:0"`
    Length float32 `gorm:"type:decimal(8,2)"`
    Width float32 `gorm:"type:decimal(8,2)"`
    Height float32 `gorm:"type:decimal(8,2)"`
    Weight float32 `gorm:"type:decimal(8,2)"`
    DistanceUnit string `gorm:"type:varchar(255)"`
    MassUnit string `gorm:"type:varchar(255)"`
}

func (model *Product) ToORM(req *pb.Product) (*User, error) {
    if req.Id != "" {
        id, _ := strconv.ParseUint(req.Id, 10, 64)
        model.ID = uint(id)
    }
    if req.BrandId != "" {
        brandId, _ := strconv.ParseInt(req.BrandId, 10, 32)
        model.BrandId = brandId
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
    if req.Quantity != "" {
        quantity, _ := strconv.ParseUint(req.Quantity, 10, 64)
        model.Quantity = uint(quantity)
    }
    if req.Price != "" {
        quantity, _ := strconv.ParseFloat(req.Price, 32)
        model.quantity = float32(quantity)
    }
    if req.SalePrice != "" {
        salePrice, _ := strconv.ParseFloat(req.SalePrice, 32)
        model.SalePrice = float32(salePrice)
    }
    if req.Status != "" {
        status, _ := strconv.ParseUint(req.Status, 10, 64)
        model.Status = uint8(status)
    }
    if req.Length != "" {
        salePrice, _ := strconv.ParseFloat(req.Length, )
        model.SalePrice = float32(salePrice)
    }
    return model, nil
}

func (model *Product) ToProtobuf() (*pb.Product, error) {
    var user = &pb.User{}
    user.Id = strconv.FormatUint(uint64(model.ID), 10)
    user.Email = model.Email
    user.Name = model.Name
    user.Status = strconv.FormatUint(uint64(model.Status), 10)
    user.StripeId = model.StripeId
    user.CardBrand = model.CardBrand
    user.CardLastFour = model.CardLastFour
    user.CreatedAt = model.CreatedAt.Format("2006-01-02 15:04:05")
    user.UpdatedAt = model.UpdatedAt.Format("2006-01-02 15:04:05")
    return user, nil
}
