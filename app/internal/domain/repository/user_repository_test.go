package repository

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nasunagisa/restapi/app/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when initializing gorm", err)
	}

	repo := NewUserRepository(gormDB)

	userId := int64(1)
	expectedUsername := "testuser"
	rows := sqlmock.NewRows([]string{"id", "username"}).AddRow(userId, expectedUsername)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "user" WHERE id = $1 ORDER BY "user"."id" LIMIT 1`)).
		WithArgs(userId).
		WillReturnRows(rows)

	user := &domain.User{}
	err = repo.GetUser(userId, user)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if user.Name != expectedUsername {
		t.Errorf("expected username to be %s, but got %s", expectedUsername, user.Name)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}