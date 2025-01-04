//go:generate mockgen -source=../domain/repository/user_repository.go -destination=./mock_user_repository.go -package=usecase

package usecase

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nasunagisa/restapi/app/internal/domain"
	"github.com/nasunagisa/restapi/app/mock/repository_mock"
	"github.com/stretchr/testify/assert"
)


func TestUsecaseGetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo :=  repository_mock.NewMockIUserRepository(ctrl)
	userUsecase := NewUserUsecase(mockRepo)

	t.Run("Success", func(t *testing.T) {
		userId := int64(1)
		user := domain.User{}
		mockRepo.EXPECT().GetUser(userId, &user).DoAndReturn(
			func(id int64, u *domain.User) error {
			u.Name = "testuser"  // user を u に変更して、引数から直接参照する
			return nil
			},
		)

		resultUser , err := userUsecase.GetUser(userId)
		expectedUser := domain.User{ Name: "testuser"}
		assert.NoError(t, err,"呼び出しエラー")
		assert.Equal(t, expectedUser, resultUser,"返り値エラー")
	})

	// t.Run("UserNotFound", func(t *testing.T) {
	// 	userId := int64(2)

	// 	mockRepo.EXPECT().GetUser(userId, gomock.Any()).Return(errors.New("user not found"))

	// 	user, err := userUsecase.GetUser(userId)
	// 	assert.Error(t, err)
	// 	assert.Equal(t, domain.User{}, user)
	// })

	// t.Run("DBError", func(t *testing.T) {
	// 	userId := int64(3)

	// 	mockRepo.EXPECT().GetUser(userId, gomock.Any()).Return(errors.New("db error"))

	// 	user, err := userUsecase.GetUser(userId)
	// 	assert.Error(t, err)
	// 	assert.Equal(t, domain.User{}, user)
	// })
}