package main

import (
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

//User model structure to export into DB
type User struct {
	ID        string `gorm:"primary_key"`
	FirstName string
	LastName  string
	Email     string `gorm:"type:varchar(255);unique;not null"`
	Password  string `gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

//BeforeCreate GOORM Lifecycle hook  to add UUID to User before Creating DB Record
func (model *User) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", uuid.String())
}

//AfterFind GOORM Lifecycle hook  to add Login Time to PropertyCredentials when logging in
func (model *PropertyCredentials) AfterFind(scope *gorm.Scope) error {
	return scope.SetColumn("LastLoginAt", time.Now())
}