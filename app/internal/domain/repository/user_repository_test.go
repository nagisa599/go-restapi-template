package repository

import (
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
			
		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `user` WHERE id = ? ORDER BY `user`.`id` LIMIT ?")).
		WithArgs(userId,1).  // id と LIMIT の値として 1 を期待
		WillReturnRows(rows)
		user := &domain.User{}
		// 実行のエラーが出ないこと
		err = repo.GetUser(userId, user)
		assert.NoError(t, err, "Error was not expected while fetching user")
		assert.Equal(t, expectedUsername, user.Name, "Expected username does not match actual username")
		assert.NoError(t, mock.ExpectationsWereMet(), "There were unfulfilled expectations")
	})

	// t.Run("GetUser_NotFound", func(t *testing.T) {
	// 	userId := int64(2)

	// 	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "user" WHERE id = ? ORDER BY "user"."id" LIMIT 1`)).
	// 		WithArgs(userId).
	// 		WillReturnError(gorm.ErrRecordNotFound)

	// 	user := &domain.User{}
	// 	err = repo.GetUser(userId, user)
	// 	if err == nil {
	// 		t.Errorf("expected error but got none")
	// 	}

	// 	if err != gorm.ErrRecordNotFound {
	// 		t.Errorf("expected error to be %s, but got %s", gorm.ErrRecordNotFound, err)
	// 	}

	// 	if err := mock.ExpectationsWereMet(); err != nil {
	// 		t.Errorf("there were unfulfilled expectations: %s", err)
	// 	}
	// })

	// t.Run("GetUser_DBError", func(t *testing.T) {
	// 	userId := int64(3)

	// 	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "user" WHERE id = ? ORDER BY "user"."id" LIMIT 1`)).
	// 		WithArgs(userId).
	// 		WillReturnError(gorm.ErrInvalidDB)

	// 	user := &domain.User{}
	// 	err = repo.GetUser(userId, user)
	// 	if err == nil {
	// 		t.Errorf("expected error but got none")
	// 	}

	// 	if err != gorm.ErrInvalidDB {
	// 		t.Errorf("expected error to be %s, but got %s", gorm.ErrInvalidDB, err)
	// 	}

	// 	if err := mock.ExpectationsWereMet(); err != nil {
	// 		t.Errorf("there were unfulfilled expectations: %s", err)
	// 	}
	// })
}