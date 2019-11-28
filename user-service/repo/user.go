package repo

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    pb "github.com/nonfu/laracom/user-service/proto/user"
)

type Repository interface {
    Create(user *pb.User) error
    Update(user *pb.User) error
    Get(id string) (*pb.User, error)
    GetByEmail(email string) (*pb.User, error)
    GetAll() ([]*pb.User, error)
}

type UserRepository struct {
    Db *gorm.DB
}

func (repo *UserRepository) Create(user *pb.User) error {
    if err := repo.Db.Create(user).Error; err != nil {
        return err
    }
    return nil
}

func (repo *UserRepository) Update(user *pb.User) error {
    if err := repo.Db.Save(user).Error; err != nil {
        return err
    }
    return nil
}

func (repo *UserRepository) Get(id string) (*pb.User, error) {
    user := &pb.User{}
    user.Id = id
    if err := repo.Db.First(&user).Error; err != nil {
        return nil, err
    }
    return user, nil
}

func (repo *UserRepository) GetByEmail(email string) (*pb.User, error) {
    user := &pb.User{}
    if err := repo.Db.Where("email = ?", email).First(&user).Error; err != nil {
        return nil, err
    }
    return user, nil
}

func (repo *UserRepository) GetAll() ([]*pb.User, error) {
    var users []*pb.User
    if err := repo.Db.Find(&users).Error; err != nil {
        return nil, err
    }
    return users, nil
}