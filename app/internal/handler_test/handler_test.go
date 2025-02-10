package handler_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/nasunagisa/restapi/app/internal/domain/repository"
	"github.com/nasunagisa/restapi/app/internal/handler"
	"github.com/nasunagisa/restapi/app/internal/usecase"
	"github.com/pressly/goose/v3"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var gdb *gorm.DB // GORMのデータベース接続をグローバルに保持
var userHandler handler.IUserHandler
func TestMain(m *testing.M) {
	ctx := context.Background()
	mysqlC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image: "mysql:8.0.32",
			Env: map[string]string{
				"MYSQL_DATABASE":             "testdb",
				"MYSQL_ALLOW_EMPTY_PASSWORD": "yes",
			},
			ExposedPorts: []string{"3306/tcp"},
			WaitingFor:   wait.ForListeningPort("3306/tcp"),
		},
		Started: true,
	})
	if err != nil {
		log.Fatalf("failed to create mysql container: %s", err)
	}
	defer func() {
		if err := mysqlC.Terminate(ctx); err != nil {
			log.Fatalf("failed to terminate mysql container: %s", err)
		}
	}()


	host, err := mysqlC.Host(ctx)
	if err != nil {
		log.Fatalf("failed to get host: %s", err)
	}
	port, err := mysqlC.MappedPort(ctx, "3306/tcp")
	if err != nil {
		log.Fatalf("failed to get externally mapped port: %s", err)
	}

	dsn := fmt.Sprintf("root:@tcp(%s:%d)/testdb?charset=utf8mb4&parseTime=true&loc=Local",
		host, port.Int())
	gdb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to open database connection: %s", err)
	}

	sqlDB, err := gdb.DB()


	migrateDB(sqlDB)
	insertTestData(gdb)

	fmt.Println("Migrate")
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	if err != nil {
		log.Fatalf("failed to get generic database object from GORM: %s", err)
	}
	if err = sqlDB.Ping(); err != nil {
		log.Fatalf("failed to verify a connection to the database: %s", err)
	}
	userRepository := repository.NewUserRepository(gdb)
    userUsecase := usecase.NewUserUsecase(userRepository)
    
   	userHandler = handler.NewUserHandler(userUsecase)  // グローバル変数への正しい代入
	if userHandler == nil {
		log.Fatalf("failed to create userHandler")
	}

	m.Run() // テストを実行し、終了コードを取得
	sqlDB.Close()        // データベース接続を閉じる

	
}

func migrateDB(db *sql.DB) {
	// マイグレーションファイルが存在するディレクトリへのパスを指定
	migrationsDir := "../../../databases/migrate"// マイグレーションファイルのディレクトリを指定

	err := goose.SetDialect("mysql")
	if err != nil {
		log.Fatal("Failed to set dialect:", err)
	}

	err = goose.Up(db, migrationsDir)
	fmt.Println("Migrate")
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
}
func insertTestData(db *gorm.DB) {
    // テスト用のユーザーデータを挿入
    users := []repository.DbUser{
        {ID: 1, Username: "testuser1"},
		{ID: 2, Username: "testuser2"},
    }
    for _, user := range users {
        err := db.Create(&user).Error
        if err != nil {
            log.Fatalf("failed to insert test data: %s", err)
        }
    }
}



