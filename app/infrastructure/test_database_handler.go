package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/pressly/goose/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)


func NewTestDatabase() *gorm.DB {
	dsn := "root:@tcp(localhost:3307)/mysql_template?parseTime=true&loc=Local&charset=utf8"
	fmt.Print(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get sql.DB from gorm.DB:", err)
	}

	// Gooseによるマイグレーション
	migrateDB(sqlDB)
	if err != nil {
		log.Fatal("Failed to connect to database",err)
	}
	if os.Getenv("ENV") == "development" {
		db.Logger = db.Logger.LogMode(logger.Info)
	} 
return db
}
// migrateDB Gooseを使用してマイグレーションを実行
func migrateDB(db *sql.DB) {
	// マイグレーションファイルが存在するディレクトリへのパスを指定
	migrationsDir := "../../../../databases/migrate"// マイグレーションファイルのディレクトリを指定

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


// func (h *DatabaseHandler) Close() {
// 	sqlDB, _ := h.DB.DB()
// 	sqlDB.Close()
// }

// func (h *DatabaseHandler) Conn(ctx context.Context) *gorm.DB {
// 	return h.DB.WithContext(ctx)
// }