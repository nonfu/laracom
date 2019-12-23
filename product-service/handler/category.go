package handler

import (
    "context"
    "errors"
    "github.com/jinzhu/gorm"
    "github.com/nonfu/laracom/product-service/model"
    pb "github.com/nonfu/laracom/product-service/proto/product"
    "github.com/nonfu/laracom/product-service/repo"
)

type CategoryService struct {
    CategoryRepo repo.CategoryRepositoryInterface
}

func (srv *CategoryService) Get(ctx context.Context, req *pb.Category, res *pb.CategoryResponse) error {
    if req.Id == 0 {
        return errors.New("类目 ID 不能为空")
    }
    category, err := srv.CategoryRepo.GetById(uint(req.Id))
    if err != nil && err != gorm.ErrRecordNotFound {
        return err
    }
    if category != nil {
        res.Category, _ = category.ToProtobuf()
    }
    return nil
}

func (srv *CategoryService) GetAll(ctx context.Context, req *pb.Request, res *pb.CategoryResponse) error {
    categories, err := srv.CategoryRepo.GetAll()
    if err != nil && err != gorm.ErrRecordNotFound {
        return err
    }
    categoryItems := make([]*pb.Category, len(categories))
    for index, category := range categories {
        categoryItem, _ := category.ToProtobuf()
        categoryItems[index] = categoryItem
    }
    res.Categories = categoryItems
    return nil
}

func (srv *CategoryService) GetWithProducts(ctx context.Context, req *pb.Category, res *pb.CategoryResponse) error {
    if req.Id == 0 {
        return errors.New("类目 ID 不能为空")
    }
    category, err := srv.CategoryRepo.GetWithProducts(uint(req.Id))
    if err != nil && err != gorm.ErrRecordNotFound {
        return err
    }
    if category != nil {
        res.Category, _ = category.ToProtobuf()
    }
    return nil
}

func (srv *CategoryService) Create(ctx context.Context, req *pb.Category, res *pb.CategoryResponse) error {
    categoryModel := &model.Category{}
    category, _ := categoryModel.ToORM(req)
    if err := srv.CategoryRepo.Create(category); err != nil {
        return err
    }
    res.Category, _ = category.ToProtobuf()
    return nil
}

func (srv *CategoryService) Update(ctx context.Context, req *pb.Category, res *pb.CategoryResponse) error {
    if req.Id == 0 {
        return errors.New("分类 ID 不能为空")
    }
    category, err := srv.CategoryRepo.GetById(uint(req.Id))
    if err != nil {
        return err
    }
    category, _ = category.ToORM(req)
    if err := srv.CategoryRepo.Update(category); err != nil {
        return err
    }
    res.Category, _ = category.ToProtobuf()
    return nil
}

func (srv *CategoryService) Delete(ctx context.Context, req *pb.Category, res *pb.CategoryResponse) error {
    if req.Id == 0 {
        return errors.New("分类 ID 不能为空")
    }
    category, err := srv.CategoryRepo.GetById(uint(req.Id))
    if err != nil {
        return err
    }
    if err := srv.CategoryRepo.Delete(category); err != nil {
        return err
    }
    res.Category = nil
    return nil
}
