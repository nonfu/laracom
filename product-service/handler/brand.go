package handler

import (
    "context"
    "errors"
    "github.com/jinzhu/gorm"
    "github.com/nonfu/laracom/product-service/model"
    pb "github.com/nonfu/laracom/product-service/proto/product"
    "github.com/nonfu/laracom/product-service/repo"
)

type BrandService struct {
    BrandRepo repo.BrandRepositoryInterface
}

func (srv *BrandService) Get(ctx context.Context, req *pb.Brand, res *pb.BrandResponse) error {
    if req.Id == 0 {
        return errors.New("品牌 ID 不能为空")
    }
    brandModel, err := srv.BrandRepo.GetById(uint(req.Id))
    if err != nil && err != gorm.ErrRecordNotFound {
        return err
    }
    if brandModel != nil {
        res.Brand, _ = brandModel.ToProtobuf()
    }
    return nil
}

func (srv *BrandService) GetAll(ctx context.Context, req *pb.Request, res *pb.BrandResponse) error {
    brands, err := srv.BrandRepo.GetAll()
    if err != nil && err != gorm.ErrRecordNotFound {
        return err
    }
    brandItems := make([]*pb.Brand, len(brands))
    for index, brand := range brands {
        brandItem, _ := brand.ToProtobuf()
        brandItems[index] = brandItem
    }
    res.Brands = brandItems
    return nil
}

func (srv *BrandService) GetWithProducts(ctx context.Context, req *pb.Brand, res *pb.BrandResponse) error {
    if req.Id == 0 {
        return errors.New("品牌 ID 不能为空")
    }
    brandModel, err := srv.BrandRepo.GetWithProducts(uint(req.Id))
    if err != nil && err != gorm.ErrRecordNotFound {
        return err
    }
    if brandModel != nil {
        res.Brand, _ = brandModel.ToProtobuf()
    }
    return nil
}

func (srv *BrandService) Create(ctx context.Context, req *pb.Brand, res *pb.BrandResponse) error {
    brandModel := &model.Brand{}
    brand, _ := brandModel.ToORM(req)
    if err := srv.BrandRepo.Create(brand); err != nil {
        return err
    }
    res.Brand, _ = brand.ToProtobuf()
    return nil
}

func (srv *BrandService) Update(ctx context.Context, req *pb.Brand, res *pb.BrandResponse) error {
    if req.Id == 0 {
        return errors.New("品牌 ID 不能为空")
    }
    brand, err := srv.BrandRepo.GetById(uint(req.Id))
    if err != nil {
        return err
    }
    brand, _ = brand.ToORM(req)
    if err := srv.BrandRepo.Update(brand); err != nil {
        return err
    }
    res.Brand, _ = brand.ToProtobuf()
    return nil
}

func (srv *BrandService) Delete(ctx context.Context, req *pb.Brand, res *pb.BrandResponse) error {
    if req.Id == 0 {
        return errors.New("品牌 ID 不能为空")
    }
    brand, err := srv.BrandRepo.GetById(uint(req.Id))
    if err != nil {
        return err
    }
    if err := srv.BrandRepo.Delete(brand); err != nil {
        return err
    }
    res.Brand = nil
    return nil
}
