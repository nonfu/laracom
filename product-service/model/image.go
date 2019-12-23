package model

import (
    pb "github.com/nonfu/laracom/product-service/proto/product"
)

type ProductImage struct {
    ID uint `gorm:"primary_key"`
    ProductId uint `gorm:"unsigned,default:0;index"`
    Src string `gorm:"type:varchar(255)"`
}

func (model *ProductImage) ToORM(req *pb.ProductImage) (*ProductImage, error) {
    if req.Id != 0 {
        model.ID = uint(req.Id)
    }
    if req.ProductId != 0 {
        model.ProductId = uint(req.ProductId)
    }
    if req.Src != "" {
        model.Src = req.Src
    }
    return model, nil
}

func (model *ProductImage) ToProtobuf() (*pb.ProductImage, error) {
    var image = &pb.ProductImage{}
    image.Id = uint32(model.ID)
    image.ProductId = uint32(model.ProductId)
    image.Src = model.Src
    return image, nil
}
