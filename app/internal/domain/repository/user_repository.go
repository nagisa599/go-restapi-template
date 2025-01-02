package repository

import (
	"github.com/nasunagisa/restapi/app/internal/domain"
	"gorm.io/gorm"
)
type DbUser struct {
	ID   int64  `gorm:"column:id"`
	Username string `gorm:"column:username"`
}

type IUserRepository interface {
	GetUser(userId int64, user *domain.User) error
}

type userRepository struct{
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUser(userId int64, user *domain.User) error {

	dbUser := &DbUser{}
	if err := ur.db.Table("user").Where("id = ?", userId).First(dbUser).Error; err != nil {
		return err
	}
	print(dbUser.Username)
	user.Name = dbUser.Username
	return nil
}
