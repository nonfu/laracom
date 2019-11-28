package laracom_service_user

import (
    "github.com/jinzhu/gorm"
    uuid "github.com/satori/go.uuid"
    "time"
)

func (model *User) BeforeCreate(scope *gorm.Scope) error {
    _ = scope.SetColumn("CreatedAt", time.Now().Format(time.RFC3339))
    uuid4 := uuid.NewV4()
    return scope.SetColumn("Id", uuid4.String())
}

func (model *User) BeforeSave(scope *gorm.Scope) error {
    _ = scope.SetColumn("UpdatedAt", time.Now().Format(time.RFC3339))
    return nil
}

func (model *PasswordReset) BeforeCreate(scope *gorm.Scope) error {
    _ = scope.SetColumn("CreatedAt", time.Now().Format(time.RFC3339))
    return nil
}