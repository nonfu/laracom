package handler

import (
    "context"
    "errors"
    "github.com/jinzhu/gorm"
    "github.com/nonfu/laracom/product-service/model"
    pb "github.com/nonfu/laracom/product-service/proto/product"
    "github.com/nonfu/laracom/product-service/repo"
)

type ImageService struct {
    ImageRepo repo.ImageRepositoryInterface
}

func (srv *ImageService) Get(ctx context.Context, req *pb.ProductImage, res *pb.ImageResponse) error {
    if req.Id == 0 {
        return errors.New("商品图片 ID 不能为空")
    }
    imageModel, err := srv.ImageRepo.GetById(uint(req.Id))
    if err != nil && err != gorm.ErrRecordNotFound {
        return err
    }
    if imageModel != nil {
        res.Image, _ = imageModel.ToProtobuf()
    }
    return nil
}

func (srv *ImageService) GetByProduct(ctx context.Context, req *pb.Product, res *pb.ImageResponse) error {
    images, err := srv.ImageRepo.GetByProductId(uint(req.Id))
    if err != nil && err != gorm.ErrRecordNotFound {
        return err
    }
    if images != nil {
        imageItems := make([]*pb.ProductImage, len(images))
        for index, value := range images {
            imageItem, _ := value.ToProtobuf()
            imageItems[index] = imageItem
        }
        res.Images = imageItems
    }
    return nil
}

func (srv *ImageService) Create(ctx context.Context, req *pb.ProductImage, res *pb.ImageResponse) error {
    imageModel := &model.ProductImage{}
    image, _ := imageModel.ToORM(req)
    if err := srv.ImageRepo.Create(image); err != nil {
        return err
    }
    res.Image, _ = image.ToProtobuf()
    return nil
}

func (srv *ImageService) Update(ctx context.Context, req *pb.ProductImage, res *pb.ImageResponse) error {
    if req.Id == 0 {
        return errors.New("商品图片 ID 不能为空")
    }
    image, err := srv.ImageRepo.GetById(uint(req.Id))
    if err != nil {
        return err
    }
    image, _ = image.ToORM(req)
    if err := srv.ImageRepo.Update(image); err != nil {
        return err
    }
    res.Image, _ = image.ToProtobuf()
    return nil
}

func (srv *ImageService) Delete(ctx context.Context, req *pb.ProductImage, res *pb.ImageResponse) error {
    if req.Id == 0 {
        return errors.New("商品图片 ID 不能为空")
    }
    image, err := srv.ImageRepo.GetById(uint(req.Id))
    if err != nil {
        return err
    }
    if err := srv.ImageRepo.Delete(image); err != nil {
        return err
    }
    res.Image = nil
    return nil
}
