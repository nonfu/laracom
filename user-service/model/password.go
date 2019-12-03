package model

import (
    pb "github.com/nonfu/laracom/user-service/proto/user"
    "time"
)

type PasswordReset struct {
    Email string `gorm:"index"`
    Token string `gorm:"not null"`
    CreatedAt time.Time
}

func (model *PasswordReset) ToORM(req *pb.PasswordReset) (*PasswordReset, error) {
    model.Email = req.Email
    model.Token = req.Token
    return model, nil
}

func (model *PasswordReset) ToProtobuf() (*pb.PasswordReset, error) {
    var reset = &pb.PasswordReset{}
    reset.Email = model.Email
    reset.Token = model.Token
    reset.CreatedAt = model.CreatedAt.Format("2006-01-02 15:04:05")
    return reset, nil
}