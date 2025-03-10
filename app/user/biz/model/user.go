package model

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email          string `gorm:"uniqueIndex;type:varchar(255) not null"`
	PasswordHashed string `gorm:"type:varchar(255) not null";`
	Name           string `gorm:"type:varchar(255)`
}

func (User) TableName() string {
	return "user"
}

func Create(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

func GetByEmail(db *gorm.DB, email string) (*User, error) {
	var user User
	err := db.Where("email = ?", email).First(&user).Error
	return &user, err
}

// ctx context.Context,
func GetByUserID(db *gorm.DB, userId uint32) (*User, error) {
	var user User
	err := db.Where("id = ?", userId).First(&user).Error
	return &user, err
}

func Update(ctx context.Context, db *gorm.DB, user *User, userId uint32) (*User, error) {
	var row User
	err := db.WithContext(ctx).
		Model(&User{}).
		Where("id = ?", userId).
		First(&row).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return user,err
	}
	
	if user.PasswordHashed == "" {
		user.PasswordHashed = row.PasswordHashed
	}
	if user.Name == "" {
		user.Name = row.Name
	}
	if user.Email == "" {
		user.Email = row.Email
	}

	return user, db.WithContext(ctx).
		Model(&User{}).
		Where("id = ?", userId).
		Updates(map[string]interface{}{"name": user.Name, "email": user.Email, "password_hashed": user.PasswordHashed}).Error
}
