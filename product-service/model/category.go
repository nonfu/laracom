package model

import (
    "github.com/jinzhu/gorm"
    pb "github.com/nonfu/laracom/product-service/proto/product"
)

type Category struct {
    gorm.Model
    Name string `gorm:"type:varchar(255);unique_index"`
    Slug string `gorm:"type:varchar(255)"`
    Description string `gorm:"type:text"`
    Cover string `gorm:"type:varchar(255)"`
    Status uint8 `gorm:"unsigned,default:0"`
    ParentId uint `gorm:"unsigned,default:0"`
    Lft uint32 `gorm:"undefined,default:0;index"`
    Rgt uint32 `gorm:"undefined,default:0;index"`
    Products []*Product `gorm:"many2many:category_product;"`
}

func (model *Category) ToORM(req *pb.Category) (*Category, error) {
    if req.Id != 0 {
        model.ID = uint(req.Id)
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
    if req.Status != 0 {
        model.Status = uint8(req.Status)
    }
    model.ParentId = uint(req.ParentId)
    model.Lft = req.Lft
    model.Rgt = req.Rgt
    return model, nil
}

func (model *Category) ToProtobuf() (*pb.Category, error) {
    var category = &pb.Category{}
    category.Id = uint32(model.ID)
    category.Name = model.Name
    category.Slug = model.Slug
    category.Description = model.Description
    category.Cover = model.Cover
    category.Status = uint32(model.Status)
    category.ParentId = uint32(model.ParentId)
    category.Lft = model.Lft
    category.Rgt = model.Rgt
    category.CreatedAt = model.CreatedAt.Format("2006-01-02 15:04:05")
    category.UpdatedAt = model.UpdatedAt.Format("2006-01-02 15:04:05")
    if model.Products != nil {
        products := make([]*pb.Product, len(model.Products))
        for index, value := range model.Products {
            product, _ := value.ToProtobuf()
            products[index] = product
        }
        category.Products = products
    }
    return category, nil
}