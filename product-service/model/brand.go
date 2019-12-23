package model

import (
    "github.com/jinzhu/gorm"
    pb "github.com/nonfu/laracom/product-service/proto/product"
)

type Brand struct {
    gorm.Model
    Name string `gorm:"type:varchar(255)"`
    Products []*Product
}

func (model *Brand) ToORM(req *pb.Brand) (*Brand, error) {
    if req.Id != 0 {
        model.ID = uint(req.Id)
    }
    if req.Name != "" {
        model.Name = req.Name
    }
    return model, nil
}

func (model *Brand) ToProtobuf() (*pb.Brand, error) {
    var brand = &pb.Brand{}
    brand.Id = uint32(model.ID)
    brand.Name = model.Name
    brand.CreatedAt = model.CreatedAt.Format("2006-01-02 15:04:05")
    brand.UpdatedAt = model.UpdatedAt.Format("2006-01-02 15:04:05")
    if model.Products != nil {
        products := make([]*pb.Product, len(model.Products))
        for index, value := range model.Products {
            product, _ := value.ToProtobuf()
            products[index] = product
        }
        brand.Products = products
    }
    return brand, nil
}
