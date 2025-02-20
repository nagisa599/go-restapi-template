package usecase

import (
	"github.com/nasunagisa/restapi/app/internal/domain"
	"github.com/nasunagisa/restapi/app/internal/domain/repository"
)

type IUserUsecase interface {
	// todoリストの一覧を取得
	GetUser(userId int64) (domain.User, error)
}
type userUsecase struct{
	ur repository.IUserRepository
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}

func (tu *userUsecase) GetUser(userId int64) (domain.User, error) {
	user := domain.User{}
	err := tu.ur.GetUser(userId, &user)
	if err != nil {
		return domain.User{}, domain.NewForbiddenError()
	}
	return user, nil
}