package usecase

import "github.com/nasunagisa/restapi/app/internal/domain"

type IUserUsecase interface {
	// todoリストの一覧を取得
	GetUser(userId int64) (domain.User, error)
}
type userUsecase struct{}

func NewUserUsecase() IUserUsecase {
	return &userUsecase{}
}

func (tu *userUsecase) GetUser(userId int64) (domain.User, error) {
	return domain.User{
		Name: "user1",
	}, nil
}