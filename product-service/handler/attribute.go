package handler

import (
    "context"
    "errors"
    "github.com/jinzhu/gorm"
    "github.com/nonfu/laracom/product-service/model"
    pb "github.com/nonfu/laracom/product-service/proto/product"
    "github.com/nonfu/laracom/product-service/repo"
)

type AttributeService struct {
    AttributeRepo repo.AttributeRepositoryInterface
}

func (srv *AttributeService) GetAttribute(ctx context.Context, req *pb.Attribute, res *pb.AttributeResponse) error {
    if req.Id == 0 {
        return errors.New("属性 ID 不能为空")
    }
    attribute, err := srv.AttributeRepo.GetAttribute(uint(req.Id))
    if err != nil && err != gorm.ErrRecordNotFound {
        return err
    }
    if attribute != nil {
        res.Attribute, _ = attribute.ToProtobuf()
    }
    return nil
}

func (srv *AttributeService) GetAttributes(ctx context.Context, req *pb.Request, res *pb.AttributeResponse) error {
    attributes, err := srv.AttributeRepo.GetAttributes()
    if err != nil && err != gorm.ErrRecordNotFound {
        return err
    }
    attributeItems := make([]*pb.Attribute, len(attributes))
    if attributes != nil {
        for index, value := range attributes {
            attribute, _ := value.ToProtobuf()
            attributeItems[index] = attribute
        }
        res.Attributes = attributeItems
    }
    return nil
}

func (srv *AttributeService) CreateAttribute(ctx context.Context, req *pb.Attribute, res *pb.AttributeResponse) error {
    attributeModel := &model.Attribute{}
    attribute, _ := attributeModel.ToORM(req)
    if err := srv.AttributeRepo.CreateAttribute(attribute); err != nil {
        return err
    }
    res.Attribute, _ = attribute.ToProtobuf()
    return nil
}

func (srv *AttributeService) UpdateAttribute(ctx context.Context, req *pb.Attribute, res *pb.AttributeResponse) error {
    if req.Id == 0 {
        return errors.New("属性 ID 不能为空")
    }
    attribute, err := srv.AttributeRepo.GetAttribute(uint(req.Id))
    if err != nil {
        return err
    }
    attribute, _ = attribute.ToORM(req)
    if err := srv.AttributeRepo.UpdateAttribute(attribute); err != nil {
        return err
    }
    res.Attribute, _ = attribute.ToProtobuf()
    return nil
}

func (srv *AttributeService) DeleteAttribute(ctx context.Context, req *pb.Attribute, res *pb.AttributeResponse) error {
    if req.Id == 0 {
        return errors.New("属性 ID 不能为空")
    }
    attribute, err := srv.AttributeRepo.GetAttribute(uint(req.Id))
    if err != nil {
        return err
    }
    if err := srv.AttributeRepo.DeleteAttribute(attribute); err != nil {
        return err
    }
    res.Attribute = nil
    return nil
}

func (srv *AttributeService) GetValue(ctx context.Context, req *pb.AttributeValue, res *pb.AttributeValueResponse) error {
    if req.Id == 0 {
        return errors.New("属性值 ID 不能为空")
    }
    value, err := srv.AttributeRepo.GetAttributeValue(uint(req.Id))
    if err != nil && err != gorm.ErrRecordNotFound {
        return err
    }
    if value != nil {
        res.Value, _ = value.ToProtobuf()
    }
    return nil
}

func (srv *AttributeService) GetValues(ctx context.Context, req *pb.Attribute, res *pb.AttributeValueResponse) error {
    if req.Id == 0 {
        return errors.New("属性 ID 不能为空")
    }
    values, err := srv.AttributeRepo.GetAttributeValues(uint(req.Id))
    if err != nil && err != gorm.ErrRecordNotFound {
        return err
    }
    valueItems := make([]*pb.AttributeValue, len(values))
    if values != nil {
        for index, value := range values {
            valueItem, _ := value.ToProtobuf()
            valueItems[index] = valueItem
        }
        res.Values = valueItems
    }
    return nil
}

func (srv *AttributeService) CreateValue(ctx context.Context, req *pb.AttributeValue, res *pb.AttributeValueResponse) error {
    valueModel := &model.AttributeValue{}
    value, _ := valueModel.ToORM(req)
    if err := srv.AttributeRepo.CreateValue(value); err != nil {
        return err
    }
    res.Value, _ = value.ToProtobuf()
    return nil
}

func (srv *AttributeService) UpdateValue(ctx context.Context, req *pb.AttributeValue, res *pb.AttributeValueResponse) error {
    if req.Id == 0 {
        return errors.New("属性值 ID 不能为空")
    }
    value, err := srv.AttributeRepo.GetAttributeValue(uint(req.Id))
    if err != nil {
        return err
    }
    value, _ = value.ToORM(req)
    if err := srv.AttributeRepo.UpdateValue(value); err != nil {
        return err
    }
    res.Value, _ = value.ToProtobuf()
    return nil
}

func (srv *AttributeService) DeleteValue(ctx context.Context, req *pb.AttributeValue, res *pb.AttributeValueResponse) error {
    if req.Id == 0 {
        return errors.New("属性值 ID 不能为空")
    }
    value, err := srv.AttributeRepo.GetAttributeValue(uint(req.Id))
    if err != nil {
        return err
    }
    if err := srv.AttributeRepo.DeleteValue(value); err != nil {
        return err
    }
    res.Value = nil
    return nil
}


func (srv *AttributeService) GetProductAttribute(ctx context.Context, req *pb.ProductAttribute, res *pb.ProductAttributeResponse) error {
    if req.Id == 0 {
        return errors.New("商品属性 ID 不能为空")
    }
    attribute, err := srv.AttributeRepo.GetProductAttribute(uint(req.Id))
    if err != nil && err != gorm.ErrRecordNotFound {
        return err
    }
    if attribute != nil {
        res.ProductAttribute, _ = attribute.ToProtobuf()
    }
    return nil
}

func (srv *AttributeService) GetProductAttributes(ctx context.Context, req *pb.Product, res *pb.ProductAttributeResponse) error {
    if req.Id == 0 {
        return errors.New("商品 ID 不能为空")
    }
    attributes, err := srv.AttributeRepo.GetProductAttributes(uint(req.Id))
    if err != nil && err != gorm.ErrRecordNotFound {
        return err
    }
    attributeItems := make([]*pb.ProductAttribute, len(attributes))
    if attributes != nil {
        for index, value := range attributes {
            attributeItem, _ := value.ToProtobuf()
            attributeItems[index] = attributeItem
        }
        res.ProductAttributes = attributeItems
    }
    return nil
}

func (srv *AttributeService) CreateProductAttribute(ctx context.Context, req *pb.ProductAttribute, res *pb.ProductAttributeResponse) error {
    attributeModel := &model.ProductAttribute{}
    attribute, _ := attributeModel.ToORM(req)
    if err := srv.AttributeRepo.CreateProductAttribute(attribute); err != nil {
        return err
    }
    res.ProductAttribute, _ = attribute.ToProtobuf()
    return nil
}

func (srv *AttributeService) UpdateProductAttribute(ctx context.Context, req *pb.ProductAttribute, res *pb.ProductAttributeResponse) error {
    if req.Id == 0 {
        return errors.New("商品属性 ID 不能为空")
    }
    attribute, err := srv.AttributeRepo.GetProductAttribute(uint(req.Id))
    if err != nil {
        return err
    }
    attribute, _ = attribute.ToORM(req)
    if err := srv.AttributeRepo.UpdateProductAttribute(attribute); err != nil {
        return err
    }
    res.ProductAttribute, _ = attribute.ToProtobuf()
    return nil
}

func (srv *AttributeService) DeleteProductAttribute(ctx context.Context, req *pb.ProductAttribute, res *pb.ProductAttributeResponse) error {
    if req.Id == 0 {
        return errors.New("商品属性 ID 不能为空")
    }
    attribute, err := srv.AttributeRepo.GetProductAttribute(uint(req.Id))
    if err != nil {
        return err
    }
    if err := srv.AttributeRepo.DeleteProductAttribute(attribute); err != nil {
        return err
    }
    res.ProductAttribute = nil
    return nil
}
