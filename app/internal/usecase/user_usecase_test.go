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
		expectedUser := domain.User{ Name: "testuser"}

		mockRepo.EXPECT().GetUser(userId, &user).DoAndReturn(
			func(id int64, user *domain.User) error {
				user.Name = "testuser"
				return nil
			},
		)
		_ , err := userUsecase.GetUser(userId)
		assert.NoError(t, err)
		assert.Equal(t, expectedUser, user)
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