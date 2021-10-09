package user

import (
	"fmt"
	"goblog/app/models"
	"goblog/pkg/logger"
	"goblog/pkg/model"
	"goblog/pkg/password"
)

type User struct {
	models.BaseModel
	Name            string `gorm:"type:varchar(255);not null;unique" valid:"name"`
	Email           string `gorm:"type:varchar(255);unique" valid:"email"`
	Password        string `gorm:"type:varchar(255)" valid:"password"`
	PasswordConfirm string `gorm:"-" valid:"password_confirm"`
}

func (user *User) Create() (err error) {
	if err = model.DB.Create(&user).Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}

// ComparePassword 对比密码是否匹配

func (u User) ComparePassword(_password string) bool {
	fmt.Println(_password, u.Password)
	return password.CheckHash(_password, u.Password)
}
