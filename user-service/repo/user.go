package repo

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/nonfu/laracom/user-service/model"
)

type Repository interface {
    Create(user *model.User) error
    Update(user *model.User) error
    Get(id uint) (*model.User, error)
    GetByEmail(email string) (*model.User, error)
    GetAll() ([]*model.User, error)
}

type UserRepository struct {
    Db *gorm.DB
}

func (repo *UserRepository) Create(user *model.User) error {
    if err := repo.Db.Create(user).Error; err != nil {
        return err
    }
    return nil
}

func (repo *UserRepository) Update(user *model.User) error {
    if err := repo.Db.Save(user).Error; err != nil {
        return err
    }
    return nil
}

func (repo *UserRepository) Get(id uint) (*model.User, error) {
    user := &model.User{}
    user.ID = id
    if err := repo.Db.First(&user).Error; err != nil {
        return nil, err
    }
    return user, nil
}

func (repo *UserRepository) GetByEmail(email string) (*model.User, error) {
    user := &model.User{}
    if err := repo.Db.Where("email = ?", email).First(&user).Error; err != nil {
        return nil, err
    }
    return user, nil
}

func (repo *UserRepository) GetAll() ([]*model.User, error) {
    var users []*model.User
    if err := repo.Db.Find(&users).Error; err != nil {
        return nil, err
    }
    return users, nil
}