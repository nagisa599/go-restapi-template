package repository

import (
	"github.com/nasunagisa/restapi/app/internal/domain"
	"gorm.io/gorm"
)
type DbUser struct {
    ID       uint   `gorm:"primaryKey"`
    Username string
}

// GORM設定でテーブル名を明示的に指定
func (DbUser) TableName() string {
    return "users"
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
	if err := ur.db.Table("users").Where("id = ?", userId).First(dbUser).Error; err != nil {
		return err
	}
	user.Name = dbUser.Username
	return nil
}
