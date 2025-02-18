package repository

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nasunagisa/restapi/app/internal/domain"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("5.7.32"))

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when initializing gorm", err)
	}

	repo := NewUserRepository(gormDB)

	t.Run("正常系", func(t *testing.T) {
		// テストデータを入れる
		userId := int64(1)
		expectedUsername := "testuser"
		rows := sqlmock.NewRows([]string{"id", "username"}).AddRow(userId, expectedUsername)
			
		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE id = ? ORDER BY `users`.`id` LIMIT ?")).
		WithArgs(userId,1).  // id と LIMIT の値として 1 を期待
		WillReturnRows(rows)
		user := &domain.User{}
		// 実行のエラーが出ないこと
		err = repo.GetUser(userId, user)
		assert.NoError(t, err, "Error was not expected while fetching user")
		assert.Equal(t, expectedUsername, user.Name, "Expected username does not match actual username")
		assert.NoError(t, mock.ExpectationsWereMet(), "There were unfulfilled expectations")
	})

	t.Run("異常系: ユーザーが見つからない2", func(t *testing.T) {
		userId := int64(1)  // 存在しないユーザーID
		expectedError := gorm.ErrRecordNotFound

		// Queryが呼ばれた際にErrRecordNotFoundを返すように設定
		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE id = ? ORDER BY `users`.`id` LIMIT ?")).
			WithArgs(userId, 1).
			WillReturnError(expectedError)

		user := &domain.User{}
		err := repo.GetUser(userId, user)

		// エラーが期待通りに発生するか検証
		assert.Error(t, err, "Expected an error when user is not found")
		assert.True(t, errors.Is(err, expectedError), "Expected a 'record not found' error")
		assert.NoError(t, mock.ExpectationsWereMet(), "There were unfulfilled expectations")
	})
}
func TestGetUserList(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("5.7.32"))

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when initializing gorm", err)
	}


	repo := NewUserRepository(gormDB)

	t.Run("正常系", func(t *testing.T) {
		// テストデータを入れる
		expectedUsers := []domain.User{
			{
				Name: "testuser1",
			},
			{
				Name: "testuser2",
			},
		}
		rows := sqlmock.NewRows([]string{"id", "username"}).
			AddRow(1, expectedUsers[0].Name).
			AddRow(2, expectedUsers[1].Name)

		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`")).
			WillReturnRows(rows)
			
		users := []domain.User{}
		// 実行のエラーが出ないこと
		err = repo.GetUserList(&users)
		assert.NoError(t, err, "Error was not expected while fetching user list")
		assert.Equal(t, expectedUsers, users, "Expected user list does not match actual user list")
		assert.NoError(t, mock.ExpectationsWereMet(), "There were unfulfilled expectations")
	})
	t.Run("異常系: ユーザーが見つからない", func(t *testing.T) {
		expectedError := gorm.ErrRecordNotFound

		// Queryが呼ばれた際にErrRecordNotFoundを返すように設定
		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`")).
			WillReturnError(expectedError)

		users := []domain.User{}
		err := repo.GetUserList(&users)

		// エラーが期待通りに発生するか検証
		assert.Error(t, err, "Expected an error when user list is not found")
		assert.True(t, errors.Is(err, expectedError), "Expected a 'record not found' error")
		assert.NoError(t, mock.ExpectationsWereMet(), "There were unfulfilled expectations")
	})

}
		// エラーが期待通りに発生するか検証