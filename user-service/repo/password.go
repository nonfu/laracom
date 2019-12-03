package repo

import (
    "github.com/jinzhu/gorm"
    "github.com/nonfu/laracom/user-service/model"
)

type PasswordResetInterface interface {
    Create(reset *model.PasswordReset) error
    Delete(reset *model.PasswordReset) error
    GetByEmail(email string) (*model.PasswordReset, error)
    GetByToken(token string) (*model.PasswordReset, error)
}

type PasswordResetRepository struct {
    Db *gorm.DB
}

func (repo *PasswordResetRepository) Create(reset *model.PasswordReset) error {
    if err := repo.Db.Create(reset).Error; err != nil {
        return err
    }
    return nil
}

func (repo *PasswordResetRepository) GetByEmail(email string) (*model.PasswordReset, error) {
    reset := &model.PasswordReset{}
    if err := repo.Db.Where("email = ?", email).First(&reset).Error; err != nil {
        return nil, err
    }
    return reset, nil
}


func (repo *PasswordResetRepository) GetByToken(token string) (*model.PasswordReset, error) {
    reset := &model.PasswordReset{}
    if err := repo.Db.Where("token = ?", token).First(&reset).Error; err != nil {
        return nil, err
    }
    return reset, nil
}

func (repo *PasswordResetRepository) Delete(reset *model.PasswordReset) error {
    err := repo.Db.Delete(reset).Error
    return err
}