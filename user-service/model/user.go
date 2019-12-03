package model

import (
    "github.com/jinzhu/gorm"
    pb "github.com/nonfu/laracom/user-service/proto/user"
    "strconv"
)

type User struct {
    gorm.Model
    Name string  `gorm:"type:varchar(100)"`
    Email string `gorm:"type:varchar(100);unique_index"`
    Password string
    Status uint8 `gorm:"default:1"`
    StripeId string
    CardBrand string
    CardLastFour string
    RememberToken string
}

func (model *User) ToORM(req *pb.User) (*User, error) {
    if req.Id != "" {
        id, _ := strconv.ParseUint(req.Id, 10, 64)
        model.ID = uint(id)
    }
    if req.Email != "" {
        model.Email = req.Email
    }
    if req.Name != "" {
        model.Name = req.Name
    }
    if req.Password != "" {
        model.Password = req.Password
    }
    if req.Status != "" {
        status, _ := strconv.ParseUint(req.Id, 10, 64)
        model.Status = uint8(status)
    }
    if req.StripeId != "" {
        model.StripeId = req.StripeId
    }
    if req.CardBrand != "" {
        model.CardBrand = req.CardBrand
    }
    if req.CardLastFour != "" {
        model.CardLastFour = req.CardLastFour
    }
    if req.RememberToken != "" {
        model.RememberToken = req.RememberToken
    }
    return model, nil
}

func (model *User) ToProtobuf() (*pb.User, error) {
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